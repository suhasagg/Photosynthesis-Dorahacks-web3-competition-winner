import argparse
from elasticsearch import Elasticsearch
from datetime import datetime, timedelta
import logging
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
   
def pairwise_tally_all(rewards, liquid_stake_data_list):
    """
    Pairs all rewards with liquid stake data (excluding amounts equal to zero) 
    by checking if each reward is greater than or equal to the first non-zero liquid stake.
    """
    paired_data_all = []

    # Filter out liquid stake data with amounts equal to zero
    filtered_liquid_stake_data = [x for x in liquid_stake_data_list if x[0] > 0]

    # Iterate over rewards and filtered_liquid_stake_data simultaneously
    for reward in rewards:
        if not filtered_liquid_stake_data:  # If we've exhausted our filtered_liquid_stake_data, break out
            break

        # Pair reward with the first liquid stake data
        data = filtered_liquid_stake_data.pop(0)  # directly pop the first element for processing
        paired_data = pairwise_tally(reward, data)

        if paired_data:
            paired_data_all.extend(paired_data)

    return paired_data_all

def pairwise_tally(reward, liquid_stake_tuple):
    """
    Pairs the reward with liquid stake tuple 
    by checking if the reward is greater than or equal to the liquid stake.
    """
    paired_data = []

    if not liquid_stake_tuple:
        logging.debug(f"No liquid stake data available. Reward {reward} cannot be paired.")
        return paired_data

    # Check if the reward is greater than or equal to the liquid stake data
    if reward >= liquid_stake_tuple[0]:
        logging.debug(f"Pairing reward {reward} with liquid stake {liquid_stake_tuple[0]}")
        paired_data.append((reward, liquid_stake_tuple[0]))  # Pair them
    else:
        logging.debug(f"Reward {reward} is not greater than or equal to the liquid stake data {liquid_stake_tuple[0]}. Skipping pairing for this reward.")
    
    logging.debug(f"Final paired data: {paired_data}")
    return paired_data


def timestamp_tally(logs):
    ist = pytz.timezone('Asia/Kolkata')  # Define the IST timezone

    reward_data = []
    liquid_stake_data = []
    valid_rewards = []  
    timestamp = None 
    for log in logs:
        logging.debug(f"Parsed log: {log}")

        try:
            timestamp = datetime.strptime(log, "%Y-%m-%dT%H:%M:%SZ")
            timestamp = timestamp.replace(tzinfo=pytz.UTC)  # Assign UTC timezone, because the strptime method doesn't attach any
            timestamp = timestamp.astimezone(ist)  # Convert to IST
            logging.debug(f"Parsed Timestamp (IST): {timestamp}")
        except ValueError:
            # If the log is not a timestamp, it should be a data log. 
            pass
        if 'Cumulative Reward amount:' in log:
            try:
                reward = int(log.split('Cumulative Reward amount:')[-1].strip())
                logging.debug(f"Extracted Cumulative Reward: {reward}")
            except ValueError:
                reward = None

            if reward is not None and timestamp is not None:  # Boundary check
                reward_data.append((reward, timestamp))
                logging.debug(f"Added Cumulative Reward: {reward} at {timestamp}")
                timestamp = None
                reward = None
        if 'Liquid stake total amount:' in log:
            try:
                stake = int(log.split('Liquid stake total amount:')[-1].strip())
                logging.debug(f"Extracted Liquid Stake: {stake}")
            except ValueError:
                stake = None

            if stake is not None and timestamp is not None:  # Similar boundary check for stakes
                liquid_stake_data.append((stake, timestamp))
                logging.debug(f"Added Liquid Stake: {stake} at {timestamp}")
                timestamp = None
                stake = None
    try:
       first_liquid_stake_timestamp = next(x[1] for x in liquid_stake_data if x[0] > 0)
    except StopIteration:
       first_liquid_stake_timestamp = None

  
    logging.info(f"Valid rewards count: {len(valid_rewards)}")
    # This log will display the timestamp in IST format
    logging.info(f"Liquid stake timestamp (for amount > 0): {first_liquid_stake_timestamp}")
 
    # Fetch all the timestamps where liquid stake amount is greater than zero
    all_liquid_stake_timestamps = [x[1] for x in liquid_stake_data if x[0] > 0]
    
    all_paired_data = []

    last_processed_timestamp = None  # Keep track of the last timestamp processed

    for stake_timestamp in all_liquid_stake_timestamps:
        adjusted_liquid_stake_timestamp = stake_timestamp + timedelta(minutes=1)
        logger.info(f"Adjusted Liquid Stake Timestamp: {adjusted_liquid_stake_timestamp.strftime('%Y-%m-%d %H:%M:%S')} IST")

        # If a last processed timestamp is there, start from that. Otherwise, start from the beginning
        if last_processed_timestamp:
            accumulated_rewards = [x[0] for x in reward_data if last_processed_timestamp < x[1] <= adjusted_liquid_stake_timestamp]
        else:
            accumulated_rewards = [x[0] for x in reward_data if x[1] <= adjusted_liquid_stake_timestamp]

        logger.debug(f"Accumulated rewards up until {adjusted_liquid_stake_timestamp}: {accumulated_rewards}")

        total_accumulated_reward = sum(accumulated_rewards)
        logger.debug(f"Total accumulated reward up until {adjusted_liquid_stake_timestamp}: {total_accumulated_reward}")

        paired_data = pairwise_tally_all(total_accumulated_reward, list(liquid_stake_data))  # Use a copy to prevent mutating original list
        all_paired_data.extend(paired_data)
        
        last_processed_timestamp = adjusted_liquid_stake_timestamp  # Update the last processed timestamp

    if all_paired_data:
        logger.info("Pairs of (reward, corresponding non-zero liquid stake):")
        for pair in all_paired_data:
            logger.info(pair)
            
    else:
        logger.warning("No valid pairs found.")
       

    return all_paired_data

  
# Fetch logs from Elasticsearch
reward_logs = get_logs(args.reward_index)
liquid_stake_logs = get_logs(args.stake_index)
all_logs = reward_logs + liquid_stake_logs
logger.info(f"Total logs count: {len(all_logs)}")

# Compute results
result = timestamp_tally(all_logs)
if result:
   logger.info(f"Pair of (reward, first non-zero liquid stake): {result}")
   print(f"Pair of (reward, first non-zero liquid stake): {result}")
else:
    logger.warning("No valid pair found.")
    print("No valid pair found.")

