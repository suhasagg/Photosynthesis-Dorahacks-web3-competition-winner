from elasticsearch import Elasticsearch
import re
import logging
import argparse

# Set up logging
logging.basicConfig(level=logging.DEBUG)
logger = logging.getLogger(__name__)

# Set up Elasticsearch client
es = Elasticsearch()

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
        "size": args.size,
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
    parser = argparse.ArgumentParser(description="Fetch logs from Elasticsearch and extract amounts.")
    parser.add_argument('--index-name', required=True, help="Name of the Elasticsearch index.")
    parser.add_argument('--size', type=int, default=100, help="Number of logs to fetch.")
    args = parser.parse_args()

    logger.info("Script started.")
    
    logs = get_logs1(args.index_name)

    # Extracting amounts
    extracted_amounts = [extract_amount(str(log)) for log in logs]

    for log, amount in zip(logs, extracted_amounts):
        logger.info(f"Extracted {amount} from log: {log}")

    logger.info("Script execution completed.")

if __name__ == "__main__":
    main()

