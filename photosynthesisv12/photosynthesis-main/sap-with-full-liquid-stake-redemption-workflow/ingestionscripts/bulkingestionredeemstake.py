from datetime import datetime
from elasticsearch import Elasticsearch
from elasticsearch.helpers import bulk, BulkIndexError
import json
import re
import traceback

es = Elasticsearch('http://localhost:9200')


def is_date_format(s):
    """Check if a string is of the format YYYY-MM-DD HH:MM:SS"""
    pattern = r'^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$'
    return re.match(pattern, s)

# Global variable to remember the last known timestamp
last_known_timestamp = None

def process_log(log_line):
    global last_known_timestamp
    
    try:
        if not log_line:
            return False, "log_line is None"

        raw_data = json.loads(log_line)
        processed_data = {}

        if 'message' in raw_data:
            if isinstance(raw_data['message'], str) and is_date_format(raw_data['message']):
                date_str = raw_data['message']
                last_known_timestamp = datetime.strptime(date_str, '%Y-%m-%d %H:%M:%S').isoformat()
                processed_data = {
                    "timestamp": last_known_timestamp
                }
            elif isinstance(raw_data['message'], dict):
                processed_data = {
                    "message": json.dumps(raw_data['message'])  # Convert the dictionary back to a string
                }
            else:
                processed_data = {
                    "message": raw_data['message']
                }

            if "timestamp" not in processed_data:
                if last_known_timestamp:
                   processed_data["timestamp"] = last_known_timestamp
                else:
                   processed_data["timestamp"] = datetime.utcnow().isoformat()
                return True, processed_data

            if 'timestamp' in raw_data:
                log = json.loads(log_line)
                date_str = ""
                if 'timestamp' in log:
                    date_str = log['timestamp']
            # If timestamp is empty, set it to the current datetime
                if not date_str:
                   log["timestamp"] = datetime.utcnow().isoformat()
            # Otherwise, check its format and process it accordingly
                elif is_date_format(date_str):
                   log["timestamp"] = datetime.strptime(date_str, '%Y-%m-%d %H:%M:%S').isoformat()
                return True, log
        
        processed_records = []

        # Handling epoch_unbonding_record
        if 'epoch_unbonding_record' in raw_data:
            for record in raw_data['epoch_unbonding_record']:
                epoch_number = record.get('epoch_number', None)
                host_zone_unbondings = record.get('host_zone_unbondings', [])
                
                for host_zone in host_zone_unbondings:
                    data = {
                        "epoch_number": epoch_number,
                        "st_token_amount": host_zone.get('st_token_amount', None),
                        "native_token_amount": host_zone.get('native_token_amount', None),
                        "denom": host_zone.get('denom', None),
                        "host_zone_id": host_zone.get('host_zone_id', None),
                        "unbonding_time": host_zone.get('unbonding_time', None),
                        "status": host_zone.get('status', None),
                        "user_redemption_records": host_zone.get('user_redemption_records', []),
                        "timestamp": last_known_timestamp if last_known_timestamp else datetime.utcnow().isoformat()
                    }
                processed_records.append(data)

        # Handling user_redemption_record
        elif 'user_redemption_record' in raw_data:
            for record in raw_data['user_redemption_record']:
                data = {
                    "id": record.get('id', None),
                    "sender": record.get('sender', None),
                    "receiver": record.get('receiver', None),
                    "amount": record.get('amount', None),
                    "denom": record.get('denom', None),
                    "host_zone_id": record.get('host_zone_id', None),
                    "epoch_number": record.get('epoch_number', None),
                    "claim_is_pending": record.get('claim_is_pending', None),
                    "timestamp": last_known_timestamp if last_known_timestamp else datetime.utcnow().isoformat()
                }
                processed_records.append(data)

        # Universal handling for any other JSON structure
        else:
            # Just treating the entire JSON as one record
            raw_data["timestamp"] = last_known_timestamp if last_known_timestamp else datetime.utcnow().isoformat()
            processed_records.append(raw_data)

        # Return results
        if processed_records:
            return True, processed_records
        else:
            return False, "Unrecognized log format"

    except Exception as e:
        print(f"Error processing log: {str(e)}")
        print(f"Offending log line: {log_line}")
        return False, f"Error: {str(e)}"

 
           

def read_and_process_log_file(file_name, index_name, bulk_size=1000):
    total_logs = 0
    successful_logs = 0
    failed_logs = 0
    bulk_ops = []

    with open(file_name, "r") as file:
        for line in file:
            total_logs += 1
            accumulated_line = line.strip()
            success, processed_log = process_log(accumulated_line)

            if success:
                # Ensure that the processed_log is always a list
                if isinstance(processed_log, dict):  # If it's a single dict, convert to a list
                    processed_log = [processed_log]

                successful_logs += len(processed_log)

                for log in processed_log:
                    action = {
                        "_index": index_name,
                        "_source": log
                    }
                    bulk_ops.append(action)

                    # To print normal dictionaries
                    print(log)
            else:
                failed_logs += 1

            if len(bulk_ops) == bulk_size:
                try:
                    bulk(es, bulk_ops)
                    bulk_ops = []
                except BulkIndexError as e:
                    for error in e.errors:
                        print("Error:", error)
                        print(traceback.format_exc())  # Printing the stack trace

    print(f"\n--- Summary ---")
    print(f"Total logs processed: {total_logs}")
    print(f"Successfully ingested logs: {successful_logs}")
    print(f"Failed logs: {failed_logs}")

def main():
    log_file_path = "/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/redeemstake.log"
    index_name = "redeemstakedata"

    # Check if the index exists
    if not es.indices.exists(index=index_name):
        es.indices.create(index=index_name)

    # Assuming the read_and_process_log_file function is defined elsewhere in your code
    read_and_process_log_file(log_file_path, index_name)


if __name__ == "__main__":
    main()

