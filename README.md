# Rocket.Chat Go SDK

This is very much a work in process.  Things could break as we clean things up.  Make sure to use something like dep to lock to a commit to prevent breakage.

### Development

A local instance of Rocket.Chat is required for unit tests to confirm connection and subscription methods are functional. An instance can be deployed locally with docker and docker compose. 

Start the server with `docker compose up -d`

For more infomation on local deployment, see the [Development Docs](https://docs.rocket.chat/quick-start/installing-and-updating).

### Testing

Tests depend on an instance of Rocket.Chat running at http://localhost:3000. This is the default configuration for Rocket.Chat instances deployed with docker compose. 

Run all tests with `go test ./...`