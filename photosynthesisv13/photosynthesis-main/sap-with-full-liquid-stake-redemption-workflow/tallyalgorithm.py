import argparse
import re
import logging
from datetime import datetime, timedelta
import pytz
from elasticsearch import Elasticsearch
from collections import deque

# Setup logging
logging.basicConfig(level=logging.DEBUG, format='%(asctime)s - %(levelname)s - %(message)s')
logger = logging.getLogger(__name__)

# Argument parser setup
parser = argparse.ArgumentParser(description='Process logs from Elasticsearch')
parser.add_argument('--host', default='localhost', help='Elasticsearch host')
parser.add_argument('--port', type=int, default=9200, help='Elasticsearch port')
parser.add_argument('--scheme', default='http', help='Elasticsearch connection scheme (http or https)')
parser.add_argument('--size', type=int, default=1000, help='Size limit for log retrieval')
parser.add_argument('--reward_index', default='distributeliquiditydataaggregated2', help='Elasticsearch index for reward logs')
parser.add_argument('--stake_index', default='liquidstakedataaggregated2', help='Elasticsearch index for liquid stake logs')
args = parser.parse_args()

# Connect to Elasticsearch
es = Elasticsearch([{'host': args.host, 'port': args.port, 'scheme': args.scheme}])
logger.info(f"Connected to Elasticsearch at {args.host}:{args.port} using {args.scheme}")


def extract_amount(message):
    match = re.search(r'(\d+)ibc/', message)
    if match:
        return int(match.group(1))
    return 0


def extract_timestamp(document):
    return document.get('timestamp')


def get_logs(index_name):
    query = {
        "size": args.size,
        "query": {
            "match_all": {}
        }
    }
    response = es.search(index=index_name, body=query)
    messages = [hit["_source"][key] for hit in response["hits"]["hits"] for key in hit["_source"]]
    return messages


def get_logs1(index_name):
    query = {
        "size": args.size,
        "query": {
            "match_all": {}
        }
    }
    response = es.search(index=index_name, body=query)
    documents = [hit["_source"] for hit in response["hits"]["hits"]]
    return documents

class PairwiseTallyHandler:
    """Handles the buffering of reward and stake data for pairing."""
    
    def __init__(self):
        self.buffered_reward = 0
        self.buffered_stake_amount = 0

    def pairwise_tally_all(self, reward, liquid_stake_data_queue):
        """Tries to find pairings for the reward with buffered data and the queue, buffering excess values."""
        paired_data_all = []
        stake_amount = 0
      
        logging.debug(f"Using buffered reward amount. Total buffered reward: {reward}")

        added_to_buffer = False  # Initialization for the buffer flag

        while liquid_stake_data_queue:
         
              if liquid_stake_data_queue:
                 data = liquid_stake_data_queue.popleft()
                 stake_amount += data[0]
                 logging.debug(f"Using buffered stake amount. New stake amount: {stake_amount}")
    

              paired_data = self._pairwise_tally(reward, stake_amount)

              if paired_data:
                 paired_data_all.extend(paired_data)
                 logging.debug(f"Successful pairing found: {paired_data}")
                 return paired_data_all  # since we have found a pairing, return immediately
              else:
                # If no pairing was found, we just move on to the next iteration of the loop
                  break
        
        logging.debug(f"Buffered reward: {self.buffered_reward}, Buffered stake amount: {self.buffered_stake_amount}")

        return paired_data_all

    def _pairwise_tally(self, reward, liquid_stake_value):
        """Pair reward and liquid stake data if reward is greater than or equal to liquid stake value."""
        paired_data = []
        if reward <= liquid_stake_value:
            paired_data.append((reward, liquid_stake_value))
        return paired_data



def timestamp_tally(logs1, logs2):
    ist = pytz.timezone('Asia/Kolkata')
    liquidity_data = []
    liquid_stake_data = []
    handler = PairwiseTallyHandler()
    for log_entry in logs1:
        # Extract timestamp
        if "timestamp" in log_entry:
            timestamp_str = log_entry["timestamp"]
            dt_obj = datetime.strptime(timestamp_str, '%Y-%m-%dT%H:%M:%SZ')
            # Convert this UTC datetime to IST
            utc = pytz.utc
            ist = pytz.timezone('Asia/Kolkata')
            dt_obj = dt_obj.replace(tzinfo=utc).astimezone(ist)
            logging.debug(f"Parsed Timestamp (IST): {dt_obj}")
           
            # Process liquidity
            liquidity = sum(extract_amount(v) for k, v in log_entry.items() if k.startswith("message"))

            if liquidity:
                liquidity_data.append((liquidity, dt_obj))

    for log in logs2:
        logging.debug(f"Parsed log: {log}")

        try:
            timestamp = datetime.strptime(log, "%Y-%m-%dT%H:%M:%SZ")
            timestamp = timestamp.replace(tzinfo=pytz.UTC)
            timestamp = timestamp.astimezone(ist)
            logging.debug(f"Parsed Timestamp (IST): {timestamp}")
        except ValueError:
            pass

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
    # Process the extracted data
    liquid_stake_data_queue = deque(liquid_stake_data)
    all_liquid_stake_timestamps = [x[1] for x in liquid_stake_data if x[0] > 0]
    all_paired_data = []
    last_processed_timestamp = None

    for stake_timestamp in all_liquid_stake_timestamps:
        adjusted_liquid_stake_timestamp = stake_timestamp + timedelta(minutes=1)
        accumulated_rewards = sum(x[0] for x in liquidity_data if (not last_processed_timestamp or last_processed_timestamp < x[1]) and x[1] > adjusted_liquid_stake_timestamp)
        paired_data = handler.pairwise_tally_all(accumulated_rewards, liquid_stake_data_queue)
        all_paired_data.extend(paired_data)
        last_processed_timestamp = adjusted_liquid_stake_timestamp

    return all_paired_data


reward_logs = get_logs1(args.reward_index)
liquid_stake_logs = get_logs(args.stake_index)

result = timestamp_tally(reward_logs, liquid_stake_logs)
if result:
    logger.info(f"Pair of (reward, first non-zero liquid stake): {result}")
    print(f"Pair of (reward, first non-zero liquid stake): {result}")
else:
    logger.warning("No valid pair found.")
    print("No valid pair found.")

