# Server Information API

This API provides server information for a specific server code without exposing the server IP or requiring direct access to the server's JSON data. It pulls data from the FiveM server frontend API (https://servers-frontend.fivem.net/api/servers/single/) and avoid CORS issues.

## Usage

To retrieve server information, make a GET request to the API endpoint with the server code as a path parameter. For example:

## GET /{server_code}

Replace `{server_code}` with the actual server FiveM CFX code you want to retrieve information for.

### Response

The API will respond with the server information in JSON format. If the server code is a vanity code (e.g., "test" or "example"), it will be mapped to the corresponding server code before fetching the data.

Example response:

```json
{
    "connect": "https://cfx.re/join/server_code",
    "endpoint": "server_code",
    "online": true,
    "hostname": "Test Server",
    "players": {
        "count": 75,
        "self-reported": 75,
        "list": [
            ...
        ]
    },
    "slots": 256,
    "private": false,
    "last-seen": "..."
}
```

## GET /original/{server_code}

Example response:

```json
{
    "connect": "https://cfx.re/join/server_code",
    "endpoint": "server_code",
    "online": true,
    "hostname": "Test Server",
    "players": {
        "count": 75,
        "self-reported": 75,
        "list": [
            ...
        ]
    },
    "original": {
        // Original API response
    }
    "slots": 256,
    "private": false,
    "last-seen": "..."
}
```

## Build From Source

1. Clone the repository: `git clone https://github.com/akatiggerx04/fivem-servers-stats-api`

2. Install modules: `go mod tidy`

3. Build: `go build main.go handlers.go`

## Deployment

The server runs at port **7747**.

## Credits

by @akatiggerx04 :)