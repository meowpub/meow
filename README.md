[![godoc reference](https://godoc.org/github.com/meowpub/meow?status.svg)](https://godoc.org/github.com/meowpub/meow)
[![Build Status](https://travis-ci.org/meowpub/meow.svg?branch=master)](https://travis-ci.org/meowpub/meow)

**:3c**

Setting up a development environment
------------------------------------

You will need `docker` and `docker-compose` installed.

1. Run `docker-compose up` to bring up Postgres and Redis containers.
1. (Optional) Load sample data for development - note that `meow` requires every node in a path.
  1. `go run . ingest < fixtures/localhost.json` - create `https://localhost`.
  1. `go run . ingest < fixtures/localhost-jane.json` - create `https://localhost/~jane`.
1. Run the server, accept requests to `localhost`: `go run . serve --domain localhost`.
1. Make some test requests:
  1. `curl -v localhost:8000`
  1. `curl -v localhost:8000/~jane`

Make some changes, then just restart `go run . serve --domain localhost` to have them reflected.
