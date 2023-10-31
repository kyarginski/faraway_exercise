# FaraWay exercise

Test task for Server engineer (Go)

Design and implement "Word of Wisdom" tcp server.

 - TCP server should be protected from DDOS attacks with the Proof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.
 - The choice of the POW algorithm should be explained.
 - After Proof Of Work verification, server should send one of the quotes from "word of wisdom" book or any other collection of the quotes.
 - Docker file should be provided both for the server and for the client that solves the POW challenge

### Settings

Set environment variables:

```shell
SERVER_CONFIG_PATH=config/server/prod.yaml
```

for run server in production mode.

Set environment variables:

```shell
SERVER_CONFIG_PATH=config/server/local.yaml
```

for run server in local debug mode.


### Run server

```shell
go run ./cmd/server
```


### Run client

```shell
go run ./cmd/client -addr=127.0.0.1:8088 connect 
```

## Working process within Docker containers

### Start:
```
docker-compose -p "faraway_exercise" up -d
```

### Stop:
```
docker-compose -p "faraway_exercise" down
```
