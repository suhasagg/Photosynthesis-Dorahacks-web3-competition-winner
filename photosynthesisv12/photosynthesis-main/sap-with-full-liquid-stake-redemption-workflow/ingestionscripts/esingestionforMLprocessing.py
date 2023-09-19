from datetime import datetime
from elasticsearch import Elasticsearch
from elasticsearch.helpers import bulk
import json

es = Elasticsearch(
    ['https://localhost:9200'],
    http_auth=('elastic', '_CduV5URIkiV_ZTy*lQc'),
    verify_certs=True,
    ca_certs="/media/usbHDD1/elasticsearch-8.10.0/config/certs/http_ca.crt"
)

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

def process_log(log_line):
    try:
        json_part = log_line.split('|', 1)[1].strip()
        log = json.loads(json_part)

        log['ts'] = datetime.strptime(log['ts'][:-4]+"Z", "%Y-%m-%dT%H:%M:%S.%fZ")

        if '_msg' in log:
            parts = log['_msg'].split(',')
            for part in parts:
                key, value = part.split(':', 1)
                key = key.strip().replace('"', '')
                value = value.strip().replace('"', '')
                value = format_value(value)
                log[key] = value

        return log
    except (json.JSONDecodeError, IndexError, ValueError) as e:
        print(f"Failed to process line: {log_line}. Error: {e}")
        return None

def read_and_process_log_file(file_name, index_name, bulk_size=20000):
    with open(file_name, "r") as file:
        bulk_ops = []
        grouped_docs = {}

        for line in file.readlines():
            processed_log = process_log(line)
            if processed_log:
                ts_str = processed_log['ts'].strftime("%Y-%m-%dT%H:%M:%S")  # Truncate to the second

                if ts_str not in grouped_docs:
                    processed_log['ts'] = ts_str  # Convert datetime to string
                    grouped_docs[ts_str] = processed_log
                else:
                    for key, value in processed_log.items():
                        if key not in grouped_docs[ts_str] or key == 'ts':
                            grouped_docs[ts_str][key] = value

        for ts, doc in grouped_docs.items():
            bulk_ops.append({
                "_index": index_name,
                "_source": doc
            })
            if len(bulk_ops) == bulk_size:
                bulk(es, bulk_ops)
                bulk_ops = []

        if bulk_ops:
            bulk(es, bulk_ops)


def read_and_process_large_log_file(file_name, index_name, bulk_size=20000):
    with open(file_name, "r") as file:
        bulk_ops = []
        current_logs = []
        pending_logs = []
        
        for line in file:
            processed_log = process_log(line)
            if processed_log:
                current_logs.append(processed_log)
                
                if len(current_logs) + len(pending_logs) >= bulk_size:
                    # Combine current_logs with previously pending logs
                    current_logs.extend(pending_logs)
                    pending_logs = []

                    # Process these logs for aggregation, but exclude the last `ts_str` for grace
                    last_ts_str = current_logs[-1]['ts'].strftime("%Y-%m-%dT%H:%M:%S")
                    pending_logs = [log for log in current_logs if log['ts'].strftime("%Y-%m-%dT%H:%M:%S") == last_ts_str]
                    current_logs = [log for log in current_logs if log['ts'].strftime("%Y-%m-%dT%H:%M:%S") != last_ts_str]

                    grouped_docs = aggregate_logs_by_ts(current_logs)
                    for ts, doc in grouped_docs.items():
                        bulk_ops.append({
                            "_index": index_name,
                            "_source": doc
                        })
                    
                    # Send the bulk request
                    if bulk_ops:
                        bulk(es, bulk_ops)
                        bulk_ops = []
                        current_logs = []

        # Handle any remaining logs
        current_logs.extend(pending_logs)
        grouped_docs = aggregate_logs_by_ts(current_logs)
        for ts, doc in grouped_docs.items():
            bulk_ops.append({
                "_index": index_name,
                "_source": doc
            })
            
        if bulk_ops:
            bulk(es, bulk_ops)

def aggregate_logs_by_ts(logs):
    grouped_docs = {}
    for log in logs:
        ts_str = log['ts'].strftime("%Y-%m-%dT%H:%M:%S")
        if ts_str not in grouped_docs:
            log['ts'] = ts_str
            grouped_docs[ts_str] = log
        else:
            for key, value in log.items():
                if key not in grouped_docs[ts_str] or key == 'ts':
                    grouped_docs[ts_str][key] = value
    return grouped_docs




def main():
    file_name = "/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/photo.log"  
    index_name = "photosynthesis_archway_log_aggregated" 

    if not es.indices.exists(index=index_name):
        mappings = {
            "mappings": {
                "properties": {
                    "ts": {
                        "type": "date",
                        "format": "strict_date_optional_time||epoch_millis"
                    }
                }
            }
        }
        es.indices.create(index=index_name, body=mappings)

    read_and_process_large_log_file(file_name, index_name)

if __name__ == "__main__":
    main()

