from datetime import datetime
from elasticsearch import Elasticsearch
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

        index_name = "photosynthesis_archway_logs_v1"
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

def read_and_process_log_file(file_name):
    with open(file_name, "r") as file:
        for line in file.readlines():
            print(line)  # Print each line
            index_name, processed_log = process_log(line)  # Process the log line
            if index_name and processed_log:  # If the line was successfully parsed
                ingest_to_es(index_name, processed_log)  # Ingest processed log to Elasticsearch

def main():
    file_name = "/media/usbHDD1/photo.log" # replace with your log file name
    read_and_process_log_file(file_name)

if __name__ == "__main__":
    main()
