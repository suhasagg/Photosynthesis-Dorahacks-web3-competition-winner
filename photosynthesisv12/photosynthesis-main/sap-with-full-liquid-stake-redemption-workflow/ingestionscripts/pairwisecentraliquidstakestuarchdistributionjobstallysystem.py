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
                    filename='/media/usbHDD1/liquidtokensdistributiontally.log', 
                    filemode='a') 
logger = logging.getLogger(__name__)

# Argument parser setup
parser = argparse.ArgumentParser(description='Process logs from Elasticsearch')
parser.add_argument('--host', default='localhost', help='Elasticsearch host')
parser.add_argument('--port', type=int, default=9200, help='Elasticsearch port')
parser.add_argument('--scheme', default='http', help='Elasticsearch connection scheme (http or https)')
parser.add_argument('--size', type=int, default=1000, help='Size limit for log retrieval')
parser.add_argument('--reward_index', default='distributeliquiditydataaggregated2', help='Elasticsearch index for liquid tokens reward logs')
parser.add_argument('--stake_index', default='liquidstakedataaggregated2', help='Elasticsearch index for liquid stake logs')
parser.add_argument('--reward_frequency', type=int, default=6, help='Frequency (in minutes) for liquid tokens (stuarch distribution) reward cron job')
parser.add_argument('--stake_frequency', type=int, default=4, help='Frequency (in minutes) for central liquid stake cron job')
args = parser.parse_args()

# Connect to Elasticsearch
es = Elasticsearch([{'host': args.host, 'port': args.port, 'scheme': args.scheme}])
logger.info(f"Connected to Elasticsearch at {args.host}:{args.port} using {args.scheme}")


def get_cron_times(frequency, start_date, start_hour, start_minute=0):
    """ 
    Generate cron times for a given frequency in minutes starting from a specific hour and minute of a day.
    
    Args:
    - frequency: int, Frequency in minutes.
    - start_date: date, The starting date.
    - start_hour: int, The starting hour.
    - start_minute: int, The starting minute.

    Returns:
    - list: List of datetime objects representing the cron times.
    """
    logger.debug(f"Generating cron times for frequency: {frequency} minutes on date {start_date} starting from {start_hour}:{start_minute}")
    ist = pytz.timezone('Asia/Kolkata')
    
    times = []
    total_minutes_in_day = 1440
    start_total_minutes = start_hour * 60 + start_minute
    for i in range((total_minutes_in_day - start_total_minutes) // frequency):
        minutes_from_start = start_total_minutes + i * frequency
        hour, minute = divmod(minutes_from_start, 60)
        times.append(datetime.datetime.combine(start_date, datetime.time(hour=hour, minute=minute, tzinfo=ist)))
    
    logger.debug(f"Cron times generated: {times}")
    return times

def extract_earliest_date_and_hour_from_logs(log1, log2):
    all_dates = []
    
    for log_entry in log1:
        if "timestamp" in log_entry:
            timestamp_str = log_entry["timestamp"]
            dt_obj = datetime.datetime.strptime(timestamp_str, '%Y-%m-%dT%H:%M:%SZ')
            all_dates.append(dt_obj)
    
    for log_entry in log2:
        try:
            timestamp = datetime.datetime.strptime(log_entry, "%Y-%m-%dT%H:%M:%SZ")
            all_dates.append(timestamp)
        except Exception as e:
            logger.error(f"Failed to parse log entry for timestamp. Error: {e}")

    # Ensure all dates are timezone aware and then get the earliest date and its hour
    earliest_date = min([ensure_timezone_aware(date) for date in all_dates])
    return earliest_date.date(), earliest_date.hour

def ensure_timezone_aware(dt):
    logger.debug(f"Ensuring timezone-aware for datetime: {dt}")
    if dt.tzinfo is None or dt.tzinfo.utcoffset(dt) is None:
        return dt.replace(tzinfo=pytz.UTC)
    return dt

def extract_amount(message):
    logger.debug(f"Extracting amount from message: {message}")
    match = re.search(r'(\d+)ibc/', message)
    if match:
        logger.debug(f"Match found: {match.group(1)}")
        return int(match.group(1))
    logger.debug(f"No match found in message.")
    return 0

def accumulate_amounts_for_cron(cron_times, log_data):
    logger.info(f"Starting accumulation for {len(cron_times)} cron times and {len(log_data)} log data entries...")

    # Debugging: Print cron_times and log_data
    logger.debug(f"Cron times: {cron_times}")
    logger.debug(f"Log data: {log_data}")

    accumulated_data = [] 
    prev_time = None
    used_indices = []

    for idx, time in enumerate(cron_times):
        # add two minutes to the current time for comparison
        time_with_delta = time + timedelta(minutes=2)
        time_with_delta_str = time_with_delta.strftime('%Y-%m-%d %H:%M:%S')
        print(f"Time with delta: {time_with_delta_str}")

        if prev_time:
            current_amounts = [x[0] for x in log_data if prev_time < x[1] <= time_with_delta]
            amount = sum(current_amounts)
            logger.debug(f"Calculated amount {amount} for interval {prev_time} to {time_with_delta}")
        else:
            current_amounts = [x[0] for x in log_data if x[1] <= time_with_delta]
            amount = sum(current_amounts)
            logger.debug(f"Calculated amount {amount} for time <= {time_with_delta}")

        # marking indices of aggregated data for removal later
        used_indices.extend([i for i, x in enumerate(log_data) if x[0] in current_amounts])

        # appending a tuple (amount, time_with_delta) to the accumulated_data list
        accumulated_data.append((amount, time_with_delta_str))
        prev_time = time  # update prev_time for the next iteration

    logger.info(f"Completed accumulation. Final accumulated data: {accumulated_data}")
    return accumulated_data


def get_logs(index_name):
    logger.debug(f"Fetching logs from Elasticsearch index: {index_name}")
    query = {
        "size": args.size,
        "query": {
            "match_all": {}
        }
    }
    response = es.search(index=index_name, body=query)
    messages = [hit["_source"][key] for hit in response["hits"]["hits"] for key in hit["_source"]]
    logger.debug(f"Retrieved {len(messages)} logs from {index_name}")
    return messages
    
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
    
def extract_earliest_datetime_from_logs(logs1, logs2):
    """
    Extract the earliest datetime from two sets of logs.

    Args:
    - logs1 (list): List of liquid tokens reward logs (stuarch distribution)
    - logs2 (list): List of central liquid stake logs.

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

    # Extract timestamps from logs2
    for log in logs2:
        try:
            timestamp = datetime.datetime.strptime(log, "%Y-%m-%dT%H:%M:%SZ")
            timestamp = ensure_timezone_aware(timestamp)
            timestamps.append(timestamp)
        except Exception as e:
            # Handle the exception if any log's timestamp can't be parsed
            logger.error(f"Failed to parse log entry for timestamp. Error: {e}")

    # Return the earliest timestamp
    return min(timestamps)
    
def pair_timestamps(left, right):
    i, j = 0, 0  # Pointers for left and right lists
    result = []  # Resultant list to store paired data

    # Continue until one of the lists is exhausted
    while i < len(left) and j < len(right):
        
        # Ensure the timestamp from the 'left' list is earlier than the timestamp from the 'right' list
        if left[i][1] < right[j][1]:
            sum_amount = 0  # Initialize aggregated amount
            
            # Accumulate amounts from the 'left' list as long as its timestamp is earlier
            # than the current timestamp from the 'right' list
            while i < len(left) and left[i][1] < right[j][1]:
                sum_amount += left[i][0]  # Aggregate the amounts
                i += 1  # Move to the next item in 'left'
            
            # Pair the aggregated amount with the last seen timestamp from 'left' 
            # and the current item from 'right'
            result.append(((sum_amount, left[i-1][1]), right[j]))
            
            j += 1  # Move to the next item in 'right'
        
        # If the current timestamp from 'right' is earlier or the same, skip it
        else:
            j += 1

    return result


def timestamp_tally(logs1, logs2, reward_frequency, stake_frequency):
    """
    Function to process logs and create timestamp-based tally.

    Args:
    - logs1 (list): List of liquid token reward logs.(stuarch distribution)
    - logs2 (list): List of liquid stake logs.
    - reward_frequency (int): Frequency (in hours) for the liquid token distribution reward cron job.
    - stake_frequency (int): Frequency (in hours) for the central liquid stake cron job.

    Returns:
    - List[Tuple[int, int]]: Pairs of accumulated amounts for liquid token reward (stuarch distribution) and central liquid stake.
    """
    logger.debug(f"Starting timestamp tally with {len(logs1)} liquid token reward logs and {len(logs2)} central liquid stake logs...")
    
    # Define the IST timezone
    ist = pytz.timezone('Asia/Kolkata')
    liquidity_data = []
    liquid_stake_data = []
    extracted_amounts = []
    # Process the liquid token reward logs (stuarch distribution) (logs1)
    for log_entry in logs1:
        if "timestamp" in log_entry:
            timestamp_str = log_entry["timestamp"]
            dt_obj = datetime.datetime.strptime(timestamp_str, '%Y-%m-%dT%H:%M:%SZ')

            # Convert the datetime object to IST timezone
            dt_obj = ensure_timezone_aware(dt_obj).astimezone(ist)
            
     # Extracting amounts
        liquidity = extract_amount(str(log_entry))
        
        # Debugging: Printing the extracted liquidity from the log entry
        logger.debug(f"Extracted liquidity: {liquidity}")

        if liquidity:
            liquidity_data.append((liquidity, dt_obj))
            
            # Debugging: Printing the liquidity data after addition
            logger.debug(f"Liquidity data after addition: {liquidity_data}")

    logger.debug(f"Processed {len(liquidity_data)} reward logs with non-zero amounts.")
    amount = 0  
    # Process the liquid stake logs (logs2)
    for log in logs2:
        try:
            timestamp = datetime.datetime.strptime(log, "%Y-%m-%dT%H:%M:%SZ")
            timestamp = ensure_timezone_aware(timestamp).astimezone(ist)            
        except Exception as e:
            logger.error(f"Failed to parse log entry for timestamp. Error: {e}")
            logger.debug("Extracting amount from log...")
    
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


    # Define the IST timezone
    ist = pytz.timezone('Asia/Kolkata')

    # Get the earliest date and hour from both sets of logs
    start_datetime = extract_earliest_datetime_from_logs(logs1, logs2)

    # Convert to IST timezone
    start_datetime_ist = start_datetime.astimezone(ist)

    # Extract date and hour
    start_date = start_datetime_ist.date()
    start_hour = start_datetime_ist.hour
    logger.debug(f"Earliest date extracted: {start_date} with hour {start_hour}.")

    # Generate cron times based on the frequencies provided
    central_liquid_stake_times = get_cron_times(stake_frequency, start_date, start_hour)
    distribute_liquidity_times = get_cron_times(reward_frequency, start_date, start_hour)
    logger.debug(f"Generated {len(central_liquid_stake_times)} central liquid stake times.")
    logger.debug(f"Generated {len(distribute_liquidity_times)} distribute liquidity times.")

    # Accumulate amounts based on the cron times generated
    central_liquid_stake_amounts = accumulate_amounts_for_cron(central_liquid_stake_times, liquid_stake_data)
    distribute_liquidity_amounts = accumulate_amounts_for_cron(distribute_liquidity_times, liquidity_data)
    logger.debug(f"Processed {len(liquid_stake_data)} central liquid stake amounts")
    logger.debug(f"Processed {len(liquidity_data)} distribute liquidity amounts")

    # Pair up the accumulated amounts from the two sets using the pair_timestamps function
    paired_data = pair_timestamps(central_liquid_stake_amounts, distribute_liquidity_amounts)

    # Logging each tuple in paired_data on a separate row
    for idx, pair in enumerate(paired_data):
        logger.debug(f"Row {idx + 1}: {pair}")

    return paired_data

if __name__ == "__main__":
    logger.info("Script started.")
    
    reward_logs = get_logs1(args.reward_index)
    liquid_stake_logs = get_logs(args.stake_index)

    result = timestamp_tally(reward_logs, liquid_stake_logs, args.reward_frequency, args.stake_frequency)
    if result:
        logger.info(f"Pairs of (liquid tokens distribution rewards, central liquid stake amount): {result}")
    logger.info("Script execution completed.")
