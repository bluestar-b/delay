# Stupid File delay server

## How it Works

The server operates as a delay mechanism from the origin. When a request is made for a media file that doesn't exist on the delay server, it downloads the file from the origin. Subsequent requests for the same file, when it already exists on the server, are faster, making it more efficient especially when deployed across multiple locations.


Note: Optimization for the initial slow request is an area under consideration for improvement.

> No more doc