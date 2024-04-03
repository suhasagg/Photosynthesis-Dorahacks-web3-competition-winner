use reqwest;
use serde_json::Value;
use std::error::Error;

async fn upload_to_ipfs(encrypted_data: Vec<u8>) -> Result<String, Box<dyn Error>> {
    let client = reqwest::Client::new();
    let url = "http://localhost:5001/api/v0/add"; // IPFS API endpoint

    // Create a multipart form request
    let form = reqwest::multipart::Form::new()
        .part("file", reqwest::multipart::Part::bytes(encrypted_data).file_name("data"));

    // Send the POST request
    let response = client.post(url)
        .multipart(form)
        .send()
        .await?;

    // Parse the response
    if response.status().is_success() {
        let response_text = response.text().await?;
        let json: Value = serde_json::from_str(&response_text)?;
        if let Some(hash) = json["Hash"].as_str() {
            Ok(hash.to_string())
        } else {
            Err("IPFS response did not contain a hash".into())
        }
    } else {
        Err(format!("IPFS upload failed with status: {}", response.status()).into())
    }
}
