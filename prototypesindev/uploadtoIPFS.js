const ipfsClient = require('ipfs-http-client');
const fs = require('fs');

// Connect to an IPFS node
const ipfs = ipfsClient.create({ host: 'localhost', port: '5001', protocol: 'http' });

// Function to upload data to IPFS
async function uploadToIPFS(filename) {
    const file = fs.readFileSync(filename);
    const { cid } = await ipfs.add({ path: filename, content: file });
    console.log("Uploaded to IPFS with CID:", cid);
}

// 'path/to/file' path to the file to upload
uploadToIPFS('path/to/file');
