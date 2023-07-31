from datetime import datetime
from elasticsearch import Elasticsearch
from elasticsearch.helpers import bulk
import json

es = Elasticsearch('http://localhost:9200')

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

        index_name = "photosynthesis_archway_logsv10"
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
        # Dictionary to store documents grouped by timestamp
        grouped_docs = {}

        for line in file.readlines():
            _, processed_log = process_log(line)
            if processed_log:
                # Convert 'ts' to a string with 1s precision
                ts_str = processed_log['ts'].strftime("%Y-%m-%dT%H:%M:%S")

                if ts_str not in grouped_docs:
                    grouped_docs[ts_str] = processed_log
                else:
                    # Merge key-value pairs
                    for key, value in processed_log.items():
                        if key not in grouped_docs[ts_str]:
                            grouped_docs[ts_str][key] = value
                        else:
                            # Modify this part to handle conflicts between key-value pairs
                            pass

        # Convert the grouped documents into bulk API operations
        bulk_ops = [
            {
                "_index": index_name,
                "_source": doc
            }
            for doc in grouped_docs.values()
        ]

        # Split bulk_ops into chunks and execute each chunk
        for i in range(0, len(bulk_ops), bulk_size):
            bulk(es, bulk_ops[i:i+bulk_size])


def main():
    file_name = "/media/usbHDD1/photo.log"  # replace with your log file name
    index_name = "photosynthesis_archway_logsv52_aggregatedsecond"  # replace with your index name
    read_and_process_log_file(file_name, index_name)

if __name__ == "__main__":
    main()

