from elasticsearch import Elasticsearch
import re
import logging
import argparse


# Setup logging
logging.basicConfig(level=logging.DEBUG, 
                    format='%(asctime)s - %(levelname)s - %(message)s', 
                    filename='/media/usbHDD1/log_file_name1.log', 
                    filemode='a')  # 'a' means append mode, use 'w' for write mode
logger = logging.getLogger(__name__)

# Argument parser setup
parser = argparse.ArgumentParser(description='Process logs from Elasticsearch')
parser.add_argument('--host', default='localhost', help='Elasticsearch host')
parser.add_argument('--port', type=int, default=9200, help='Elasticsearch port')
parser.add_argument('--scheme', default='http', help='Elasticsearch connection scheme (http or https)')
parser.add_argument('--size', type=int, default=1000, help='Size limit for log retrieval')
parser.add_argument('--reward_index', default='distributeliquiditydataaggregated2', help='Elasticsearch index for reward logs')
parser.add_argument('--stake_index', default='liquidstakedataaggregated2', help='Elasticsearch index for liquid stake logs')
parser.add_argument('--reward_frequency', type=int, default=6, help='Frequency (in minutes) for reward cron job')
parser.add_argument('--stake_frequency', type=int, default=4, help='Frequency (in minutes) for stake cron job')
args = parser.parse_args()

# Set up logging
logging.basicConfig(level=logging.DEBUG)
logger = logging.getLogger(__name__)

# Set up Elasticsearch client
# Connect to Elasticsearch
es = Elasticsearch([{'host': "localhost", 'port': 9200, 'scheme': "http"}])


def extract_amount(message):
    logger.debug(f"Extracting amount from message: {message}")
    match = re.search(r'(\d+)ibc/', message)
    if match:
        logger.debug(f"Match found: {match.group(1)}")
        return int(match.group(1))
    logger.debug(f"No match found in message.")
    return 0

def get_logs1(index_name):
    """
    Fetches full log data from the specified Elasticsearch index.

    Args:
    - index_name (str): Name of the Elasticsearch index.

    Returns:
    - list: List containing the entire `_source` data for each log entry.
    """
    logger.debug(f"Fetching full log data from Elasticsearch index: {index_name}")
    query = {
        "size": 1000,
        "query": {
            "match_all": {}
        }
    }
    response = es.search(index=index_name, body=query)
    
    # Extracting the entire _source data for each log entry
    logs = [hit["_source"] for hit in response["hits"]["hits"]]
    logger.debug(f"Retrieved {len(logs)} log entries from {index_name}")
    return logs 

def main():
 
    logger.info("Script started.")
    
    logs = get_logs1("distributeliquiditydataaggregated2")

    # Extracting amounts
    extracted_amounts = [extract_amount(str(log)) for log in logs]

    for log, amount in zip(logs, extracted_amounts):
        logger.info(f"Extracted {amount} from log: {log}")

    logger.info("Script execution completed.")

if __name__ == "__main__":
    main()

