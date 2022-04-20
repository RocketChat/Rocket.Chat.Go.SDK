# Rocket.Chat Go SDK

This is very much a work in process.  Things could break as we clean things up.  Make sure to use something like dep to lock to a commit to prevent breakage.

## Development

A local instance of Rocket.Chat is required for unit tests to confirm connection and subscription methods are functional. 

### Installing Rocket.Chat

Please see the [Development Docs](https://docs.rocket.chat/quick-start/installing-and-updating) for information on locally deploying a Rocket.Chat instance. Deploying with [Docker & Docker Compose](https://docs.rocket.chat/quick-start/installing-and-updating/rapid-deployment-methods/docker-and-docker-compose) is recommended. 

### Testing

Tests depend on an instance of Rocket.Chat running at http://localhost:3000. This is the default configuration for Rocket.Chat instances deployed with Docker Compose. 

To test the `rest` and `realtime` packages, navigate to the respective directories and run `go test`. 