from datetime import datetime
from elasticsearch import Elasticsearch
from elasticsearch.helpers import bulk
import json
import mmap

es = Elasticsearch(
    ['https://localhost:9200'],
    http_auth=('elastic', '_CduV5URIkiV_ZTy*lQc'),
    verify_certs=True,
    ca_certs="/media/usbHDD1/elasticsearch-8.10.0/config/certs/http_ca.crt"
)

def format_value(value):
    try:
        # Try to convert to int
        return int(value)
    except ValueError:
        pass

    try:
        # Try to convert to float
        return float(value)
    except ValueError:
        pass

    if value.lower() in ['true', 'false']:
        # Convert to boolean
        return value.lower() == 'true'

    # If value cannot be converted to int, float, or bool, return as is
    return value



def process_log(log_line):
    try:
        # Split the line by '|', take the second part, and strip leading/trailing whitespace
        json_part = log_line.split('|', 1)[1].strip()

        # Parse json_part into a dictionary
        log = json.loads(json_part)

        index_name = "photosynthesis_archway_logs"
        # Transform 'ts' into a datetime object for Elasticsearch
        # Truncate the last 3 digits of the nanoseconds to convert it to microseconds
        log['ts'] = datetime.strptime(log['ts'][:-4]+"Z", "%Y-%m-%dT%H:%M:%S.%fZ")

        # If '_msg' is in the log, process it
        if '_msg' in log:
            # Extract the message content, and split into parts by ','
            parts = log['_msg'].split(',')
            for part in parts:
                # Split the part into key and value by ':'
                key, value = part.split(':', 1)
                # Strip leading/trailing whitespace and remove special characters
                key = key.strip().replace('"', '')
                value = value.strip().replace('"', '')
                value = format_value(value)

                # Add the parsed key-value pair to the log
                log[key] = value

        return index_name, log
    except (json.JSONDecodeError, IndexError, ValueError):
        print(f"Failed to process line: {log_line}")
        return None, None


def ingest_to_es(index_name, log):
    # Check if index exists. If not, create it.
    if not es.indices.exists(index=index_name):
        es.indices.create(index=index_name)
        
    # Index the document
    es.index(index=index_name, body=log)

def read_and_process_log_file(file_name, index_name, bulk_size=20000):
    with open(file_name, "r") as file:
        # List to store bulk API operations
        bulk_ops = []
        
        for line in file.readlines():
            _, processed_log = process_log(line)
            if processed_log:
                # Action for the Bulk API
                action = {
                    "_index": index_name,
                    "_source": processed_log
                }
                
                # Add action to the operations list
                bulk_ops.append(action)
                
                # If the bulk size is reached, execute the bulk request
                if len(bulk_ops) == bulk_size:
                    bulk(es, bulk_ops)
                    bulk_ops = []
        
        # Execute any remaining bulk operations
        if bulk_ops:
            bulk(es, bulk_ops)


def read_and_process_large_log_file(file_name, index_name, bulk_size=20000):
    bulk_ops = []  # This list will hold bulk operations.

    with open(file_name, 'r') as file:
        for line in file:  # This reads the file line by line.
            _, processed_log = process_log(line)
            if processed_log:
                # Action for the Bulk API
                action = {
                    "_index": index_name,
                    "_source": processed_log
                }
                
                # Add action to the operations list
                bulk_ops.append(action)

                # If the bulk size is reached, execute the bulk request
                if len(bulk_ops) == bulk_size:
                    bulk(es, bulk_ops)
                    bulk_ops = []  # Reset the list for the next set of operations.

    # Execute any remaining bulk operations
    if bulk_ops:
        bulk(es, bulk_ops)


def main():
    file_name = "/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/photo.log" 
    index_name = "photosynthesis_archway_logs" 
    read_and_process_large_log_file(file_name, index_name)

if __name__ == "__main__":
    main()

