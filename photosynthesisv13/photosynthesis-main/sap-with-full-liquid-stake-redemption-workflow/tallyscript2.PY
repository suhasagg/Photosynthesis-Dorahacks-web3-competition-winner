import argparse
from elasticsearch import Elasticsearch
from datetime import datetime, timedelta
import logging
import pytz
from collections import deque

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

def pairwise_tally_all(reward, liquid_stake_data_queue):
    paired_data_all = []

    # If the queue is empty, return the empty list
    if not liquid_stake_data_queue:
        return paired_data_all

    # Pair reward with the first liquid stake data
    data = liquid_stake_data_queue.popleft()
    paired_data = pairwise_tally(reward, data[0])

    if paired_data:
        paired_data_all.extend(paired_data)

    return paired_data_all


def pairwise_tally(reward, liquid_stake_value):
    paired_data = []

    if reward >= liquid_stake_value:
        logging.debug(f"Pairing reward {reward} with liquid stake {liquid_stake_value}")
        paired_data.append((reward, liquid_stake_value))

    else:
        logging.debug(f"Reward {reward} is not greater than or equal to the liquid stake data {liquid_stake_value}. Skipping pairing for this reward.")
    
    logging.debug(f"Final paired data: {paired_data}")
    return paired_data

def timestamp_tally(logs):
    ist = pytz.timezone('Asia/Kolkata')

    reward_data = []
    liquid_stake_data = []
    for log in logs:
        logging.debug(f"Parsed log: {log}")

        try:
            timestamp = datetime.strptime(log, "%Y-%m-%dT%H:%M:%SZ")
            timestamp = timestamp.replace(tzinfo=pytz.UTC)
            timestamp = timestamp.astimezone(ist)
            logging.debug(f"Parsed Timestamp (IST): {timestamp}")
        except ValueError:
            pass

        if 'Cumulative Reward amount:' in log:
            try:
                reward = int(log.split('Cumulative Reward amount:')[-1].strip())
                logging.debug(f"Extracted Cumulative Reward: {reward}")
            except ValueError:
                reward = None

            if reward is not None and timestamp is not None:
                reward_data.append((reward, timestamp))
                timestamp = None
                reward = None

        if 'Liquid stake total amount:' in log:
            try:
                stake = int(log.split('Liquid stake total amount:')[-1].strip())
                logging.debug(f"Extracted Liquid Stake: {stake}")
            except ValueError:
                stake = None

            if stake is not None and stake > 0 and timestamp is not None:
               liquid_stake_data.append((stake, timestamp))
               timestamp = None
               stake = None

    liquid_stake_data_queue = deque(liquid_stake_data)

    all_liquid_stake_timestamps = [x[1] for x in liquid_stake_data if x[0] > 0]
    all_paired_data = []

    last_processed_timestamp = None

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

        paired_data = pairwise_tally_all(total_accumulated_reward, liquid_stake_data_queue)
        all_paired_data.extend(paired_data)

        last_processed_timestamp = adjusted_liquid_stake_timestamp

    if all_paired_data:
        logger.info("Pairs of (reward, corresponding non-zero liquid stake):")
        for pair in all_paired_data:
            logger.info(pair)
    else:
        logger.warning("No valid pairs found.")

    return all_paired_data

reward_logs = get_logs(args.reward_index)
liquid_stake_logs = get_logs(args.stake_index)
all_logs = reward_logs + liquid_stake_logs
logger.info(f"Total logs count: {len(all_logs)}")

result = timestamp_tally(all_logs)
if result:
    logger.info(f"Pair of (reward, first non-zero liquid stake): {result}")
    print(f"Pair of (reward, first non-zero liquid stake): {result}")
else:
    logger.warning("No valid pair found.")
    print("No valid pair found.")
