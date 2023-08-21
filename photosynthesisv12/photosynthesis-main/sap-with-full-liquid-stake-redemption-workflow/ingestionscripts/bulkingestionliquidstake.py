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

def process_log(log_line):
    try:
        if not log_line:
            return False, "log_line is None"
        
        raw_data = json.loads(log_line)
        processed_data = {}
        if 'message' in raw_data:
            if isinstance(raw_data['message'], str) and is_date_format(raw_data['message']):
                date_str = raw_data['message']
                processed_data = {
                        "timestamp": datetime.strptime(date_str, '%Y-%m-%d %H:%M:%S').isoformat()
                    }             
               
            elif isinstance(raw_data['message'], dict):
                  processed_data = {
                          "message": json.dumps(raw_data['message'])  # Convert the dictionary back to a string
                      }
            else:                                
                  processed_data = {
                          "message": raw_data['message']
                      }
            return True, processed_data
             
         # Process 'timestamp' key in the log data       
        else:
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

    except Exception as e:
        print(f"Error processing log: {str(e)}")
        print(f"Offending log line: {log_line}")
        return False, f"Error: {str(e)}"


def read_and_process_log_file(file_name, index_name, bulk_size=1000):
    total_logs = 0
    successful_logs = 0
    failed_logs = 0

    with open(file_name, "r") as file:
        bulk_ops = []

        for line in file:
            total_logs += 1
            accumulated_line = line.strip()
            success, processed_log = process_log(accumulated_line)

            if success:
                successful_logs += 1
                action = {
                    "_index": index_name,
                    "_source": processed_log
                }
                bulk_ops.append(action)

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

        if bulk_ops:
            try:
                bulk(es, bulk_ops)
            except BulkIndexError as e:
                for error in e.errors:
                    print("Error:", error)
                    print(traceback.format_exc())  # Printing the stack trace

    print(f"\n--- Summary ---")
    print(f"Total logs processed: {total_logs}")
    print(f"Successfully ingested logs: {successful_logs}")
    print(f"Failed logs: {failed_logs}")

def main():
    log_file_path = "/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/liquidstakelogs"
    index_name = "liquidstakedata"

    # Define the mapping
    mapping = {
       "mappings": {
          "properties": {
            "message": {
                "type": "text",
                "fields": {
                    "keyword": {
                        "type": "keyword",
                        "ignore_above": 2147483647
                    }
                }
            }
        }
       }
    }

    # Check if the index exists
    if not es.indices.exists(index=index_name):
        es.indices.create(index=index_name, body=mapping)

    read_and_process_log_file(log_file_path, index_name)

if __name__ == "__main__":
    main()
