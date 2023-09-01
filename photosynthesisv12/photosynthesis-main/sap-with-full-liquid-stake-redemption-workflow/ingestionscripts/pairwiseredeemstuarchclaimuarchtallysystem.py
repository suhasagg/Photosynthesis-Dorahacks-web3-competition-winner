import argparse
import re
import logging
import datetime
import pytz
from elasticsearch import Elasticsearch
from collections import deque
from datetime import timedelta


# Setup logging
logging.basicConfig(level=logging.DEBUG, 
                    format='%(asctime)s - %(levelname)s - %(message)s', 
                    filename='/media/usbHDD1/redeemtokenclaimedamountstally.log', 
                    filemode='a') 
logger = logging.getLogger(__name__)

# Argument parser setup
parser = argparse.ArgumentParser(description='Process logs from Elasticsearch')
parser.add_argument('--host', default='localhost', help='Elasticsearch host')
parser.add_argument('--port', type=int, default=9200, help='Elasticsearch port')
parser.add_argument('--scheme', default='http', help='Elasticsearch connection scheme (http or https)')
parser.add_argument('--size', type=int, default=1000, help='Size limit for log retrieval')
parser.add_argument('--redeem_stake_index', default='redeemstakedataaggregated2', help='Elasticsearch index for redeem stake logs')
parser.add_argument('--redeem_frequency', type=int, default=4, help='Frequency (in minutes) for central redeem stake (at maximum redemption rate) cron job')
args = parser.parse_args()

# Connect to Elasticsearch
es = Elasticsearch([{'host': args.host, 'port': args.port, 'scheme': args.scheme}])
logger.info(f"Connected to Elasticsearch at {args.host}:{args.port} using {args.scheme}")

def ensure_timezone_aware(dt):
    logger.debug(f"Ensuring timezone-aware for datetime: {dt}")
    if dt.tzinfo is None or dt.tzinfo.utcoffset(dt) is None:
        return dt.replace(tzinfo=pytz.UTC)
    return dt

def extract_redeem_amount(message):
    logger.debug(f"Extracting amount from message: {message}")
    
    # The regex looks for a space followed by a sequence of digits and then another space
    match = re.search(r'\s(\d+)\s', message)
    
    if match:
        logger.debug(f"Match found: {match.group(1)}")
        return int(match.group(1))
    
    logger.debug(f"No match found in message.")
    return 0


def extract_claimed_amount(message):
    """
    Extracts the difference between the "Amount before claim" and "Amount after claim"
    :param message: string message to be parsed
    :return: difference in amount or None if one of the amounts is zero
    """
    
    amount_before_claim = 0
    amount_after_claim = 0
    
    before_claim_match = re.search(r'Amount before claim: (\d+)', message)
    after_claim_match = re.search(r'Amount after claim: (\d+)', message)

    if before_claim_match:
        amount_before_claim = int(before_claim_match.group(1))
        logger.debug(f"Amount before claim extracted: {amount_before_claim}")

    if after_claim_match:
        amount_after_claim = int(after_claim_match.group(1))
        logger.debug(f"Amount after claim extracted: {amount_after_claim}")

    # Ensure both amounts are non-zero before taking the difference
    if amount_before_claim > 0 and amount_after_claim > 0:
        difference = amount_after_claim - amount_before_claim
        logger.debug(f"Difference in amount: {difference}")
        return difference
    else:
        logger.warning("One or both of the amounts are zero. Difference not computed.")
        return None

   
def get_logs(index_name):
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
    
def extract_earliest_datetime_from_logs(logs1):
    """
    Extract the earliest datetime from log

    Args:
    - logs1 (list): List of central redeem stuarch (at maximum redemption rate) logs.

    Returns:
    - datetime: Earliest datetime extracted from the logs.
    """
    # Create an empty list to store all timestamps
    timestamps = []

    # Extract timestamps from logs1
    for log_entry in logs1:
        if "timestamp" in log_entry:
            timestamp_str = log_entry["timestamp"]
            dt_obj = datetime.datetime.strptime(timestamp_str, '%Y-%m-%dT%H:%M:%SZ')
            dt_obj = ensure_timezone_aware(dt_obj)
            timestamps.append(dt_obj)

    # Return the earliest timestamp
    return min(timestamps)

def pair_redemptions_to_claims(redemptions, claims):
    i, j = 0, 0  # Pointers for redemptions and claims
    result = []  # Resultant list to store paired data
    aggregated_redemption = 0  # Current aggregated redemption amount

    # Debugging log
    logger.debug("Starting to pair redemptions to claims...")

    # Continue until one of the lists is exhausted
    while i < len(redemptions) and j < len(claims):
        
        # Debugging log
        logger.debug(f"Redemption pointer (i): {i}, Claim pointer (j): {j}, Aggregated Redemption: {aggregated_redemption}")
        
        # While the aggregated redemption is less than the claim, keep accumulating
        while i < len(redemptions) and aggregated_redemption < claims[j][0]:
            aggregated_redemption += redemptions[i][0]
            
            # Debugging logs inside the loop
            logger.debug(f"Adding redemption of value {redemptions[i][0]} to the aggregate.")
            logger.debug(f"New Aggregated Redemption: {aggregated_redemption}, Current Claim: {claims[j][0]}")
            
            # Debugging logs inside the loop
            logger.debug(f"Adding redemption of value {redemptions[i][0]} to the aggregate.")
            i += 1
        
        # Debugging log
        logger.debug(f"Accumulated Redemption: {aggregated_redemption}, Last Redemption Timestamp: {redemptions[i-1][1]}")

    # If the aggregated redemption is now greater than or equal to the claim, pair them
        if aggregated_redemption >= claims[j][0]:
       # Pair the aggregated redemption amount with the timestamp of the last redemption
           aggregated_pair = (aggregated_redemption, redemptions[i-1][1])
           result.append((aggregated_pair, claims[j]))
            
           # Debugging log
           logger.debug(f"Paired {aggregated_pair} with Claim {claims[j]}")
            
           j += 1  # Move to the next claim
           aggregated_redemption = 0  # Reset the aggregated redemption amount
      
    # Debugging log
    logger.debug(f"Finished. Resultant pairs: {result}")
    return result
    
       
def timestamp_tally(logs1):
    """
    Function to process logs and create timestamp-based tally.

    Args:
    - logs1 (list): List of central redeem stuarch (at maximum redemption rate) logs.
    
    Returns:
    - List[Tuple[int, int]]: Pairs of accumulated amounts for central redeem stuarch (at maximum redemption rate) and claimed amount (after processing user redemption records and epoch unbonding records) in liquid tokens provider chain
    """
    logger.debug(f"Starting timestamp tally with {len(logs1)} central redeem stuarch (at maximum redemption rate) logs ...")
    
    # Define the IST timezone
    ist = pytz.timezone('Asia/Kolkata')
    claimed_data = []
    redemption_data = []
    extracted_amounts = []
    # Process the claimed amount logs (claiming redeemed tokens) (logs1)
    for log_entry in logs1:
        logger.debug(f"Print logs: {log_entry}")
        if "timestamp" in log_entry:
            timestamp_str = log_entry["timestamp"]
            dt_obj = datetime.datetime.strptime(timestamp_str, '%Y-%m-%dT%H:%M:%SZ')

            # Convert the datetime object to IST timezone
            dt_obj = ensure_timezone_aware(dt_obj).astimezone(ist)
            
        # Extracting amounts
        claimed_amount = extract_claimed_amount(str(log_entry))
        
        # Debugging: Printing the extracted claimed amount from the log entry
        logger.debug(f"Extracted claimed amount: {claimed_amount}")

        if claimed_amount and claimed_amount > 0:
           claimed_data.append((claimed_amount, dt_obj))
    
           # Debugging: Printing the claimed data after addition
           logger.debug(f"Claimed data after redemption: {claimed_data}")

    logger.debug(f"Processed {len(claimed_data)} logs with non-zero amounts.")
    amount = 0  
    # Process the redeemed uarch logs (at maximum redemption rate) (logs1)
    for log in logs1:
        logger.debug(f"Print logs: {log}")
        if "timestamp" in log:
            timestamp_str = log_entry["timestamp"]
            dt_obj = datetime.datetime.strptime(timestamp_str, '%Y-%m-%dT%H:%M:%SZ')

            # Convert the datetime object to IST timezone
            dt_obj = ensure_timezone_aware(dt_obj).astimezone(ist)
        try:
           stake = extract_redeem_amount(str(log))
           logging.debug(f"Redeem stuarch (at maximum redemption rate): {stake}")
        except ValueError:
               stake = None

        if stake is not None and stake > 0 and dt_obj is not None:
           redemption_data.append((stake, dt_obj))
           dt_obj = None
           stake = None

    # Define the IST timezone
    ist = pytz.timezone('Asia/Kolkata')

    # Get the earliest date and hour from both sets of logs
    start_datetime = extract_earliest_datetime_from_logs(logs1)

    # Convert to IST timezone
    start_datetime_ist = start_datetime.astimezone(ist)

    # Extract date and hour
    start_date = start_datetime_ist.date()
    start_hour = start_datetime_ist.hour
    logger.debug(f"Earliest date extracted: {start_date} with hour {start_hour}.")
    paired_data = pair_redemptions_to_claims(redemption_data,claimed_data)

    # Logging each tuple in paired_data on a separate row
    for idx, pair in enumerate(paired_data):
        logger.debug(f"Row {idx + 1}: {pair}")

    return paired_data

if __name__ == "__main__":
    logger.info("Script started.")
   
    redeem_stuarch_logs = get_logs(args.redeem_stake_index)

    result = timestamp_tally(redeem_stuarch_logs)
    if result:
        logger.info(f"Pairs of (liquid tokens distribution rewards (stuarch), claim uarch (after processing user redemption record and epoch unbonding record on liquid tokens provider chain)) : {result}")
    logger.info("Script execution completed.")


