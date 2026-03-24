const http = require('http');
const fs = require('fs');
const path = require('path');
const url = require('url');

// Root directory to serve files from
const PUBLIC_DIR = path.join(__dirname, 'public');

// Basic MIME type map (extend as needed)
const MIME_TYPES = {
  '.html': 'text/html',
  '.js': 'application/javascript',
  '.css': 'text/css',
  '.json': 'application/json',
  '.png': 'image/png',
  '.jpg': 'image/jpeg',
  '.jpeg': 'image/jpeg',
  '.gif': 'image/gif',
  '.svg': 'image/svg+xml',
  '.ico': 'image/x-icon',
  '.txt': 'text/plain'
};

const server = http.createServer((req, res) => {
  // Parse URL safely
  const parsedUrl = url.parse(req.url);
  let pathname = decodeURIComponent(parsedUrl.pathname);

  // Prevent directory traversal attacks
  pathname = pathname.replace(/^(\.\.[\/\\])+/, '');

  // Default file resolution
  let filePath = path.join(PUBLIC_DIR, pathname);
  if (filePath.endsWith(path.sep)) {
    filePath = path.join(filePath, 'index.html');
  }

  // Resolve file extension
  const ext = path.extname(filePath).toLowerCase();
  const contentType = MIME_TYPES[ext] || 'application/octet-stream';

  // Check file existence and stream
  fs.stat(filePath, (err, stats) => {
    if (err || !stats.isFile()) {
      res.writeHead(404, { 'Content-Type': 'text/plain' });
      res.end('404 Not Found');
      return;
    }

    // Stream file (efficient for large files)
    res.writeHead(200, { 'Content-Type': contentType });

    const readStream = fs.createReadStream(filePath);

    readStream.on('error', (streamErr) => {
      res.writeHead(500, { 'Content-Type': 'text/plain' });
      res.end('500 Internal Server Error');
    });

    readStream.pipe(res);
  });
});

const PORT = 3000;

server.listen(PORT, () => {
  console.log(`Static server running at http://localhost:${PORT}`);
});const http = require('http');
const fs = require('fs');
const path = require('path');
const url = require('url');

// Root directory to serve files from
const PUBLIC_DIR = path.join(__dirname, 'public');

// Basic MIME type map (extend as needed)
const MIME_TYPES = {
  '.html': 'text/html',
  '.js': 'application/javascript',
  '.css': 'text/css',
  '.json': 'application/json',
  '.png': 'image/png',
  '.jpg': 'image/jpeg',
  '.jpeg': 'image/jpeg',
  '.gif': 'image/gif',
  '.svg': 'image/svg+xml',
  '.ico': 'image/x-icon',
  '.txt': 'text/plain'
};

const server = http.createServer((req, res) => {
  // Parse URL safely
  const parsedUrl = url.parse(req.url);
  let pathname = decodeURIComponent(parsedUrl.pathname);

  // Prevent directory traversal attacks
  pathname = pathname.replace(/^(\.\.[\/\\])+/, '');

  // Default file resolution
  let filePath = path.join(PUBLIC_DIR, pathname);
  if (filePath.endsWith(path.sep)) {
    filePath = path.join(filePath, 'index.html');
  }

  // Resolve file extension
  const ext = path.extname(filePath).toLowerCase();
  const contentType = MIME_TYPES[ext] || 'application/octet-stream';

  // Check file existence and stream
  fs.stat(filePath, (err, stats) => {
    if (err || !stats.isFile()) {
      res.writeHead(404, { 'Content-Type': 'text/plain' });
      res.end('404 Not Found');
      return;
    }

    // Stream file (efficient for large files)
    res.writeHead(200, { 'Content-Type': contentType });

    const readStream = fs.createReadStream(filePath);

    readStream.on('error', (streamErr) => {
      res.writeHead(500, { 'Content-Type': 'text/plain' });
      res.end('500 Internal Server Error');
    });

    readStream.pipe(res);
  });
});

const PORT = 3000;

server.listen(PORT, () => {
  console.log(`Static server running at http://localhost:${PORT}`);
});