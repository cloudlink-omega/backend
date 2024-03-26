![go report card](https://goreportcard.com/badge/github.com/cloudlink-omega/backend)

![Backend Thumbnail](https://github.com/cloudlink-omega/backend/assets/12957745/9a09a982-4c29-493e-83f7-036f05b02b9a)

# Backend
Implementation of the CL5 signaling protocol and CLΩ API in Go. Powered by Chi and Gorilla Websockets.

## Basics
All API endpoints use `application/json` or `text/plain` types. See documentation for endpoints.

Use this directory to build the server binary (use `go build .`).

Configuration is done via an environment variables file. See `.env.example`
for a template.

## Database
This backend code was designed with a MariaDB server in mind, but should be compatible with any
standard SQL server. Tables will be auto-generated on first launch.
