from datetime import datetime
from elasticsearch import Elasticsearch
from elasticsearch.helpers import bulk, scan
import json

es = Elasticsearch('http://localhost:9200')

def format_value(value):
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
    try:
        if '_source' in document:
            log = document['_source']
            #print(log)
            return log
        else:
            #print(f"Document does not contain '_source' field: {document}")
            return None
    except Exception as e:
        #print(f"Failed to process document: {document}. Error: {e}")
        return None

def get_latest_timestamp(index_name):
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
    Modify this function to apply transformations on the document fetched from source index.
    For now, this will return the log as is.
    """
    # Example transformation: add a new field
    log['processed'] = True
    return log
    

from datetime import datetime, timedelta
from elasticsearch.helpers import scan, bulk

def process_log(document):
    # Placeholder function for processing the log.
    # Modify it to suit your actual log processing needs.
    return document['_source']

def round_to_nearest_three_minutes(ts_str):
    dt = datetime.strptime(ts_str, "%Y-%m-%dT%H:%M")
    rounded_minute = (dt.minute // 3) * 3
    return dt.replace(minute=rounded_minute, second=0)

def process_and_aggregate_documents_from_index(source_index, target_index, bulk_size=20000):
    bulk_ops = []
    grouped_docs = {}

    # Using scan to get all the documents from the source index
    for document in scan(es, index=source_index):
        processed_log = process_log(document)
        if processed_log and 'message' in processed_log:
            # Convert and round the timestamp to the nearest 3-minute mark
            rounded_time = round_to_nearest_three_minutes(processed_log['timestamp'][:16])
            ts_str = rounded_time.strftime("%Y-%m-%dT%H:%M")

            if ts_str not in grouped_docs:
                grouped_docs[ts_str] = {
                    'timestamp': ts_str + ':00Z',
                    'message_1': processed_log['message']
                }
            else:
                count = 2  # starting from the second message
                while f"message_{count}" in grouped_docs[ts_str]:
                    count += 1
                grouped_docs[ts_str][f"message_{count}"] = processed_log['message']

    for ts, doc in grouped_docs.items():
        print(f"Preparing bulk operation for timestamp {ts}")  # Debugging line
        print(doc)
        bulk_ops.append({
            "_index": target_index,
            "_source": doc
        })

        if len(bulk_ops) == bulk_size:
            print(f"Bulk operation size reached: {bulk_size}. Committing bulk operation.")  # Debugging line
            bulk(es, bulk_ops)
            bulk_ops = []

    if bulk_ops:
        print(f"Committing remaining operations. Count: {len(bulk_ops)}")  # Debugging line
        bulk(es, bulk_ops)


def create_or_update_index_with_mappings(index_name):
    # Define the mapping
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

    if not es.indices.exists(index=index_name):
        es.indices.create(index=index_name, body=mapping)
    else:
        # Be cautious with this, as reindexing will be needed if you're updating an existing index's mapping
        # For this example, I'm assuming the index is either new or can be safely overwritten
        es.indices.delete(index=index_name)
        es.indices.create(index=index_name, body=mapping)

def main():
    source_index_name = "distributeredeemtokensdata"
    target_index_name = "distributeredeemtokensdataaggregated"

    create_or_update_index_with_mappings(target_index_name)

    last_timestamp = get_latest_timestamp(target_index_name)
    if last_timestamp:
        documents = fetch_documents_after_timestamp(source_index_name, last_timestamp)
        process_and_aggregate_documents(documents=documents, target_index=target_index_name)
    else:
        process_and_aggregate_documents_from_index(source_index=source_index_name, target_index=target_index_name)


if __name__ == "__main__":
    main()

