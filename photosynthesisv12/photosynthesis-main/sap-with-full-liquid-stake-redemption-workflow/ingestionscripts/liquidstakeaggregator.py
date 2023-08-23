from datetime import datetime
from elasticsearch import Elasticsearch
from elasticsearch.helpers import bulk, scan
import json

# Connect to Elasticsearch instance running on localhost at port 9200
es = Elasticsearch('http://localhost:9200')

def format_value(value):
    """
    Converts the value into its appropriate data type (integer, float, boolean, or string).
    """
    # Try converting to integer
    try:
        return int(value)
    except ValueError:
        pass

    # Try converting to float
    try:
        return float(value)
    except ValueError:
        pass

    # Convert strings 'true' and 'false' to their respective boolean values
    if value.lower() in ['true', 'false']:
        return value.lower() == 'true'

    # If not convertible to any of the above, return as is.
    return value

def process_log(document):
    """
    Extract the log from the '_source' field of the Elasticsearch document.
    If not present, it returns None.
    """
    try:
        if '_source' in document:
            return document['_source']
        else:
            return None
    except Exception as e:
        return None

def get_latest_timestamp(index_name):
    """
    Fetches the latest timestamp 'ts' from the given Elasticsearch index.
    """
    # Define the Elasticsearch query
    body = {
        "size": 1,
        "sort": [{"ts": {"order": "desc"}}],
        "_source": ["ts"]
    }

    # Execute the search query
    res = es.search(index=index_name, body=body)

    # Extract and return the latest timestamp if available
    if res['hits']['hits'] and '_source' in res['hits']['hits'][0] and 'ts' in res['hits']['hits'][0]['_source']:
        return res['hits']['hits'][0]['_source']['ts']
    else:
        print("No 'ts' field found in the latest document or no documents present in the index.")
        return None

def fetch_documents_after_timestamp(index_name, timestamp):
    """
    Fetches all documents from a given Elasticsearch index after the specified timestamp.
    """
    # Define the Elasticsearch query
    body = {
        "query": {
            "range": {
                "timestamp": {
                    "gt": timestamp
                }
            }
        }
    }

    # Execute the search query with scrolling enabled for large datasets
    res = es.search(index=index_name, body=body, scroll='1m')
    scroll_id = res['_scroll_id']
    scroll_size = len(res['hits']['hits'])

    documents = res['hits']['hits']
    
    # Continue fetching documents using the scroll API until all documents are fetched
    while scroll_size > 0:
        res = es.scroll(scroll_id=scroll_id, scroll='1m')
        scroll_id = res['_scroll_id']
        scroll_size = len(res['hits']['hits'])
        documents.extend(res['hits']['hits'])

    return documents

def transform_document(log):
    """
    Applies transformation on the fetched document.
    Here, it adds a 'processed' field to the document.
    """
    log['processed'] = True
    return log

def round_to_nearest_three_minutes(ts_str):
    """
    Rounds off a given timestamp string to its nearest 3-minute mark.
    """
    dt = datetime.strptime(ts_str, "%Y-%m-%dT%H:%M")
    rounded_minute = (dt.minute // 3) * 3
    return dt.replace(minute=rounded_minute, second=0)

def process_and_aggregate_documents_from_index(source_index, target_index, bulk_size=20000):
    """
    Fetches, processes, and aggregates documents from the source index,
    then bulk indexes them into the target index.
    """
    bulk_ops = []
    grouped_docs = {}

    # Using scan to get all the documents from the source index
    for document in scan(es, index=source_index):
        processed_log = process_log(document)
        if processed_log and 'message' in processed_log:
            # Convert and round the timestamp to the nearest 3-minute mark
            rounded_time = round_to_nearest_three_minutes(processed_log['timestamp'][:16])
            ts_str = rounded_time.strftime("%Y-%m-%dT%H:%M")

            # Grouping documents by their rounded timestamps
            if ts_str not in grouped_docs:
                grouped_docs[ts_str] = {
                    'timestamp': ts_str + ':00Z',
                    'message_1': processed_log['message']
                }
            else:
                # If there's already an entry for this timestamp, append new messages
                count = 2  # starting from the second message
                while f"message_{count}" in grouped_docs[ts_str]:
                    count += 1
                grouped_docs[ts_str][f"message_{count}"] = processed_log['message']

    # Prepare bulk operations
    for ts, doc in grouped_docs.items():
        bulk_ops.append({
            "_index": target_index,
            "_source": doc
        })

        # Perform bulk operation when we hit the specified bulk size
        if len(bulk_ops) == bulk_size:
            bulk(es, bulk_ops)
            bulk_ops = []

    # Perform any remaining bulk operations
    if bulk_ops:
        bulk(es, bulk_ops)

def create_or_update_index_with_mappings(index_name):
    """
    Creates or updates an Elasticsearch index with specified mappings.
    If the index exists, it deletes and recreates it.
    """
    # Define the mapping for the index
    mapping = {
        "mappings": {
            "properties": {
                "ts": {
                    "type": "date",
                    "format": "strict_date_optional_time||epoch_millis"
                },
                "message": {
                    "type": "text",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 2147483647
                        }
                    }
                },
                # Add other fields as needed
            }
        }
    }

    # Check if index exists
    if not es.indices.exists(index=index_name):
        es.indices.create(index=index_name, body=mapping)
    else:
        # If index exists, delete and then recreate with the new mapping
        es.indices.delete(index=index_name)
        es.indices.create(index=index_name, body=mapping)

def main():
    """
    Main driver function.
    """
    source_index_name = "liquidstakedata"
    target_index_name = "liquidstakedataaggregated"

    # Create or update the target index with mappings
    create_or_update_index_with_mappings(target_index_name)

    # Get the latest timestamp from the target index
    last_timestamp = get_latest_timestamp(target_index_name)

    # If there's a last timestamp, fetch documents after this timestamp and process them
    if last_timestamp:
        documents = fetch_documents_after_timestamp(source_index_name, last_timestamp)
        process_and_aggregate_documents(documents=documents, target_index=target_index_name)
    # If no last timestamp is found, process and aggregate all documents from the source index
    else:
        process_and_aggregate_documents_from_index(source_index=source_index_name, target_index=target_index_name)

if __name__ == "__main__":
    main()

