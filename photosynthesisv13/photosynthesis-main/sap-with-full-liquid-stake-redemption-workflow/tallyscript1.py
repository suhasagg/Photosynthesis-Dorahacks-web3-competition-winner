import argparse
from elasticsearch import Elasticsearch
from datetime import datetime, timedelta
import logging
from collections import deque
from concurrent.futures import ThreadPoolExecutor
import pytz  # Added for timezone handling

# Setup logging
logging.basicConfig(level=logging.DEBUG, format='%(asctime)s - %(levelname)s - %(message)s')
logger = logging.getLogger(__name__)

# Argument parser
parser = argparse.ArgumentParser(description='Process logs from Elasticsearch')
parser.add_argument('--host', default='localhost', help='Elasticsearch host')
parser.add_argument('--port', type=int, default=9200, help='Elasticsearch port')
parser.add_argument('--scheme', default='http', help='Elasticsearch connection scheme (http or https)')
parser.add_argument('--size', type=int, default=1000, help='Size limit for log retrieval')
parser.add_argument('--reward_index', default='rewardsmoduleaggregated2', help='Elasticsearch index for reward logs')
parser.add_argument('--stake_index', default='liquidstakedataaggregated2', help='Elasticsearch index for liquid stake logs')

args = parser.parse_args()

# Connect to Elasticsearch
es = Elasticsearch([{'host': args.host, 'port': args.port, 'scheme': args.scheme}])
logger.info(f"Connected to Elasticsearch at {args.host}:{args.port} using {args.scheme}")

def get_logs(index_name):
    query = {
        "size": args.size,
        "query": {
            "match_all": {}
        }
    }
    logger.info(f"Fetching logs from index: {index_name}")
    response = es.search(index=index_name, body=query)
    logger.info(f"Retrieved {len(response['hits']['hits'])} logs from index: {index_name}")

    messages = []
    for hit in response["hits"]["hits"]:
        for key in hit["_source"]:
            messages.append(hit["_source"][key])

    return messages

# Shared deque to keep track of liquid stakes
filtered_liquid_stake_data = deque()

def process_rewards_and_stakes(all_liquid_stake_timestamps, reward_data, liquid_stake_data_list):
    # Populate the shared deque
    for stake in liquid_stake_data_list:
        try:
           if int(stake[0]) > 0:
              filtered_liquid_stake_data.append(stake)
        except ValueError:
               print(f"Unable to convert {stake[0]} to an integer.")

           
    last_processed_timestamp = None
    all_paired_data = []

    for stake_timestamp in all_liquid_stake_timestamps:
        adjusted_liquid_stake_timestamp = stake_timestamp + timedelta(minutes=1)
        logger.info(f"Adjusted Liquid Stake Timestamp: {adjusted_liquid_stake_timestamp.strftime('%Y-%m-%d %H:%M:%S')} IST")

        if last_processed_timestamp:
            accumulated_rewards = [x[0] for x in reward_data if last_processed_timestamp < x[1] <= adjusted_liquid_stake_timestamp]
        else:
            accumulated_rewards = [x[0] for x in reward_data if x[1] <= adjusted_liquid_stake_timestamp]

        logger.debug(f"Accumulated rewards up until {adjusted_liquid_stake_timestamp}: {accumulated_rewards}")

        total_accumulated_reward = sum(accumulated_rewards)
        logger.debug(f"Total accumulated reward up until {adjusted_liquid_stake_timestamp}: {total_accumulated_reward}")

        with ThreadPoolExecutor() as executor:
            results = list(executor.map(pairwise_tally, accumulated_rewards))
            all_paired_data.extend(filter(None, results))

        last_processed_timestamp = adjusted_liquid_stake_timestamp  # Update the last processed timestamp

    if all_paired_data:
        logger.info("Pairs of (reward, corresponding non-zero liquid stake):")
        for pair in all_paired_data:
            logger.info(pair)
    else:
        logger.warning("No valid pairs found.")

def pairwise_tally(reward):
    """
    Pairs the reward with liquid stake tuple 
    by checking if the reward is greater than or equal to the liquid stake.
    """
    paired_data = []

    if not filtered_liquid_stake_data:
        logger.debug(f"No liquid stake data available. Reward {reward} cannot be paired.")
        return paired_data

    # Pair reward with the first liquid stake data
    data = filtered_liquid_stake_data.popleft()  # directly pop the first element for processing

    # Check if the reward is greater than or equal to the liquid stake data
    if reward >= data[0]:
        logger.debug(f"Pairing reward {reward} with liquid stake {data[0]}")
        paired_data.append((reward, data[0]))  # Pair them
    else:
        logger.debug(f"Reward {reward} is not greater than or equal to the liquid stake data {data[0]}. Skipping pairing for this reward.")
        filtered_liquid_stake_data.appendleft(data)  # Return the unprocessed stake data back to the deque

    logger.debug(f"Final paired data: {paired_data}")
    return paired_data

# Integrate the above with the main function and other methods accordingly

# Fetch logs from Elasticsearch
reward_logs = get_logs(args.reward_index)
liquid_stake_logs = get_logs(args.stake_index)
all_logs = reward_logs + liquid_stake_logs
logger.info(f"Total logs count: {len(all_logs)}")

# Compute results using the process_rewards_and_stakes function
result = process_rewards_and_stakes(all_logs,reward_logs,liquid_stake_logs)
if result:
   logger.info(f"Pair of (reward, first non-zero liquid stake): {result}")
   print(f"Pair of (reward, first non-zero liquid stake): {result}")
else:
    logger.warning("No valid pair found.")
    print("No valid pair found.")

