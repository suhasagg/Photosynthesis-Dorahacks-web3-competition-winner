from datetime import datetime
from elasticsearch import Elasticsearch
from elasticsearch.helpers import bulk, scan
import json

# Establish a connection to the Elasticsearch instance
es = Elasticsearch('http://localhost:9200')

def format_value(value):
    """
    Convert a string value to int, float, boolean, or retain as string.
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

    # Convert 'true' and 'false' strings to boolean
    if value.lower() in ['true', 'false']:
        return value.lower() == 'true'

    # Return original string if no conversions apply
    return value

def process_log(document):
    """
    Extract the '_source' content from the Elasticsearch document.
    """
    try:
        if '_source' in document:
            log = document['_source']
            return log
        else:
            return None
    except Exception as e:
        return None

def get_latest_timestamp(index_name):
    """
    Fetch the latest timestamp ('ts') from a given Elasticsearch index.
    """
    # Query the latest document based on timestamp
    body = {
        "size": 1,
        "sort": [
            {
                "ts": {
                    "order": "desc"
                }
            }
        ],
        "_source": ["ts"]
    }

    res = es.search(index=index_name, body=body)
    if res['hits']['hits'] and '_source' in res['hits']['hits'][0] and 'ts' in res['hits']['hits'][0]['_source']:
        return res['hits']['hits'][0]['_source']['ts']
    else:
        print("No 'ts' field found in the latest document or no documents present in the index.")
        return None
      
def fetch_documents_after_timestamp(index_name, timestamp):
    """
    Retrieve all documents from a specified index after a given timestamp.
    Uses scrolling for efficient retrieval of large amounts of data.
    """
    body = {
        "query": {
            "range": {
                "timestamp": {
                    "gt": timestamp
                }
            }
        }
    }
    res = es.search(index=index_name, body=body, scroll='1m')
    scroll_id = res['_scroll_id']
    scroll_size = len(res['hits']['hits'])

    documents = res['hits']['hits']
    while scroll_size > 0:
        res = es.scroll(scroll_id=scroll_id, scroll='1m')
        scroll_id = res['_scroll_id']
        scroll_size = len(res['hits']['hits'])
        documents.extend(res['hits']['hits'])

    return documents

def transform_document(log):
    """
    Apply transformations to the document.
    In this placeholder function, a 'processed' field is added.
    """
    log['processed'] = True
    return log

def process_log(document):
    """
    Placeholder function for processing the log.
    """
    return document['_source']

def round_to_nearest_three_minutes(ts_str):
    """
    Round a timestamp to its nearest 3-minute mark.
    """
    dt = datetime.strptime(ts_str, "%Y-%m-%dT%H:%M")
    rounded_minute = (dt.minute // 3) * 3
    return dt.replace(minute=rounded_minute, second=0)

def process_and_aggregate_documents_from_index(source_index, target_index, bulk_size=20000):
    """
    Fetch logs, group them in 3-minute intervals, and insert into a target index.
    """
    bulk_ops = []
    grouped_docs = {}

    for document in scan(es, index=source_index):
        processed_log = process_log(document)
        if processed_log and 'message' in processed_log:
            rounded_time = round_to_nearest_three_minutes(processed_log['timestamp'][:16])
            ts_str = rounded_time.strftime("%Y-%m-%dT%H:%M")

            if ts_str not in grouped_docs:
                grouped_docs[ts_str] = {
                    'timestamp': ts_str + ':00Z',
                    'message_1': processed_log['message']
                }
            else:
                count = 2  # start from the second message
                while f"message_{count}" in grouped_docs[ts_str]:
                    count += 1
                grouped_docs[ts_str][f"message_{count}"] = processed_log['message']

    for ts, doc in grouped_docs.items():
        bulk_ops.append({
            "_index": target_index,
            "_source": doc
        })

        if len(bulk_ops) == bulk_size:
            bulk(es, bulk_ops)
            bulk_ops = []

    if bulk_ops:
        bulk(es, bulk_ops)

def create_index_with_mappings(index_name):
    """
    Create an Elasticsearch index with the specified mappings only if it doesn't exist.
    """
    # Define the index mappings
    mapping = {
        "mappings": {
            "properties": {
                "ts": {"type": "date", "format": "strict_date_optional_time||epoch_millis"},
                "message": {"type": "text", "fields": {"keyword": {"type": "keyword", "ignore_above": 2147483647}}}
            }
        }
    }

    # Create the index only if it doesn't exist
    if not es.indices.exists(index=index_name):
        es.indices.create(index=index_name, body=mapping)
        
def main():
    """
    Main execution function.
    """
    source_index_name = "distributeredeemtokensdata"
    target_index_name = "distributeredeemtokensdataaggregated"

    create_index_with_mappings(target_index_name)

    last_timestamp = get_latest_timestamp(target_index_name)
    if last_timestamp:
        documents = fetch_documents_after_timestamp(source_index_name, last_timestamp)
        process_and_aggregate_documents(documents=documents, target_index=target_index_name)
    else:
        process_and_aggregate_documents_from_index(source_index=source_index_name, target_index=target_index_name)

if __name__ == "__main__":
    main()
