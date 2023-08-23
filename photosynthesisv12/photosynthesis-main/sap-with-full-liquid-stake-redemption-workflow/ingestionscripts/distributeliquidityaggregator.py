from datetime import datetime
from elasticsearch import Elasticsearch
from elasticsearch.helpers import bulk, scan
import json

# Establish a connection to the Elasticsearch instance
es = Elasticsearch('http://localhost:9200')

def format_value(value):
    """
    Attempt to format the provided value into int, float, or boolean.
    If none of these fit, return the original value.
    """
    try:
        return int(value)
    except ValueError:
        pass

    try:
        return float(value)
    except ValueError:
        pass

    if value.lower() in ['true', 'false']:
        return value.lower() == 'true'

    return value

def process_log(document):
    """
    Extract the '_source' field from the provided Elasticsearch document.
    If '_source' is not present or an error occurs, return None.
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
    Fetch the latest timestamp ('ts') from the provided Elasticsearch index.
    If no documents are present or the 'ts' field is not found, return None.
    """
    # Elasticsearch query to sort by timestamp in descending order and fetch the top document
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
        return None

def fetch_documents_after_timestamp(index_name, timestamp):
    """
    Retrieve all Elasticsearch documents from the provided index that have a timestamp greater than the given timestamp.
    This function uses the 'scroll' feature to fetch large amounts of data in chunks.
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
    documents = res['hits']['hits']

    # Continue scrolling until all documents have been retrieved
    while len(res['hits']['hits']) > 0:
        res = es.scroll(scroll_id=scroll_id, scroll='1m')
        documents.extend(res['hits']['hits'])

    return documents

def transform_document(log):
    """
    Apply transformations to the document fetched from the source index.
    In this case, it's just adding a 'processed' field.
    """
    log['processed'] = True
    return log

def process_log(document):
    """
    A placeholder function for further processing of the log.
    Currently, it just returns the '_source' field of the provided document.
    """
    return document['_source']

def round_to_nearest_three_minutes(ts_str):
    """
    Round the provided timestamp string to the nearest three-minute mark.
    """
    dt = datetime.strptime(ts_str, "%Y-%m-%dT%H:%M")
    rounded_minute = (dt.minute // 3) * 3
    return dt.replace(minute=rounded_minute, second=0)

def process_and_aggregate_documents_from_index(source_index, target_index, bulk_size=20000):
    """
    Fetch documents from the source index, aggregate them based on their rounded timestamps,
    and then store the aggregated data into the target index using bulk operations.
    """
    bulk_ops = []
    grouped_docs = {}

    # Retrieve all documents from the source index
    for document in scan(es, index=source_index):
        processed_log = process_log(document)
        if processed_log and 'message' in processed_log:
            rounded_time = round_to_nearest_three_minutes(processed_log['timestamp'][:16])
            ts_str = rounded_time.strftime("%Y-%m-%dT%H:%M")

            # Group messages by their rounded timestamps
            if ts_str not in grouped_docs:
                grouped_docs[ts_str] = {'timestamp': ts_str + ':00Z', 'message_1': processed_log['message']}
            else:
                count = 2
                while f"message_{count}" in grouped_docs[ts_str]:
                    count += 1
                grouped_docs[ts_str][f"message_{count}"] = processed_log['message']

    # Prepare bulk operations for the aggregated data
    for ts, doc in grouped_docs.items():
        bulk_ops.append({"_index": target_index, "_source": doc})

        # Execute bulk operations in chunks to avoid overloading the system
        if len(bulk_ops) == bulk_size:
            bulk(es, bulk_ops)
            bulk_ops = []

    # Commit any remaining operations
    if bulk_ops:
        bulk(es, bulk_ops)

def create_or_update_index_with_mappings(index_name):
    """
    Create or update an Elasticsearch index with the specified mappings.
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

    # Check if the index exists
    if not es.indices.exists(index=index_name):
        es.indices.create(index=index_name, body=mapping)
    else:
        # If index exists, delete and recreate it with the new mapping
        es.indices.delete(index=index_name)
        es.indices.create(index=index_name, body=mapping)

def main():
    """
    Main execution function:
    - Sets source and target index names.
    - Creates or updates the target index with necessary mappings.
    - Fetches new documents from the source index based on the latest timestamp in the target index.
    - Processes and aggregates these documents, then stores them in the target index.
    """
    source_index_name = "distributeliquiditydata"

    target_index_name = "distributeliquiditydataaggregated"

    create_or_update_index_with_mappings(target_index_name)

    last_timestamp = get_latest_timestamp(target_index_name)
    if last_timestamp:
        documents = fetch_documents_after_timestamp(source_index_name, last_timestamp)
        process_and_aggregate_documents(documents=documents, target_index=target_index_name)
    else:
        process_and_aggregate_documents_from_index(source_index=source_index_name, target_index=target_index_name)


if __name__ == "__main__":
    main()
