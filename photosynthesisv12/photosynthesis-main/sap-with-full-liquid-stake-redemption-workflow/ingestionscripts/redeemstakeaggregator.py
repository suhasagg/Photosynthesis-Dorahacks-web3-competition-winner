from datetime import datetime
from elasticsearch import Elasticsearch
from elasticsearch.helpers import bulk, scan
import json

# Initialize the Elasticsearch client
es = Elasticsearch('http://localhost:9200')

def format_value(value):
    """
    Convert string values to their respective native data types:
    int, float, bool, or retain as string.
    """
    # Try converting the value to int
    try:
        return int(value)
    except ValueError:
        pass

    # Try converting the value to float
    try:
        return float(value)
    except ValueError:
        pass

    # Check if the value is a boolean string representation
    if value.lower() in ['true', 'false']:
        return value.lower() == 'true'

    return value  # If none of the above, return the string value

def process_log(document):
    """
    Extract the '_source' field from a document.
    Returns the content of '_source' or None if not present.
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
    Fetch the latest timestamp ('ts' field) from the specified index.
    Returns the timestamp or None if not present.
    """
    # Elasticsearch query to fetch the latest timestamp
    body = {
        "size": 1,
        "sort": [{
            "ts": {
                "order": "desc"
            }
        }],
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
    Fetch documents from the specified index that have a timestamp greater than the provided timestamp.
    Returns the list of documents.
    """
    # Elasticsearch query to fetch documents after a given timestamp
    body = {
        "query": {
            "range": {
                "timestamp": {
                    "gt": timestamp
                }
            }
        }
    }

    # Use Elasticsearch scroll API for pagination
    res = es.search(index=index_name, body=body, scroll='1m')
    scroll_id = res['_scroll_id']
    documents = res['hits']['hits']
    while len(documents) > 0:
        res = es.scroll(scroll_id=scroll_id, scroll='1m')
        scroll_id = res['_scroll_id']
        documents.extend(res['hits']['hits'])

    return documents

def transform_document(log):
    """
    Transform the document fetched from the source index.
    For now, adds a field indicating the document has been processed.
    """
    log['processed'] = True
    return log

def round_to_nearest_three_minutes(ts_str):
    """
    Round a given timestamp string to the nearest 3-minute mark.
    """
    dt = datetime.strptime(ts_str, "%Y-%m-%dT%H:%M")
    rounded_minute = (dt.minute // 3) * 3
    return dt.replace(minute=rounded_minute, second=0)

def process_and_aggregate_documents_from_index(source_index, target_index, bulk_size=20000):
    """
    Process and aggregate documents from the source index,
    and bulk index them into the target index.
    """
    bulk_ops = []  # List to store bulk operations
    grouped_docs = {}

    # Scan all documents in the source index
    for document in scan(es, index=source_index):
        processed_log = process_log(document)
        if processed_log and 'message' in processed_log:
            rounded_time = round_to_nearest_three_minutes(processed_log['timestamp'][:16])
            ts_str = rounded_time.strftime("%Y-%m-%dT%H:%M")

            # Group documents by rounded timestamp and store multiple messages under the same timestamp
            if ts_str not in grouped_docs:
                grouped_docs[ts_str] = {
                    'timestamp': ts_str + ':00Z',
                    'message_1': processed_log['message']
                }
            else:
                count = 2
                while f"message_{count}" in grouped_docs[ts_str]:
                    count += 1
                grouped_docs[ts_str][f"message_{count}"] = processed_log['message']

    # Prepare the bulk indexing operations
    for ts, doc in grouped_docs.items():
        bulk_ops.append({
            "_index": target_index,
            "_source": doc
        })

        # Perform the bulk indexing once reaching the specified size
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
    Main function to execute the aggregation process.
    """
    source_index_name = "redeemstakedata"
    target_index_name = "redeemstakedataaggregated"

    create_index_with_mappings(target_index_name)

    last_timestamp = get_latest_timestamp(target_index_name)
    if last_timestamp:
        documents = fetch_documents_after_timestamp(source_index_name, last_timestamp)
        process_and_aggregate_documents(documents=documents, target_index=target_index_name)
    else:
        process_and_aggregate_documents_from_index(source_index=source_index_name, target_index=target_index_name)


if __name__ == "__main__":
    main()
