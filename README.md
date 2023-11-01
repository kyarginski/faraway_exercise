# FaraWay exercise

Test task for Server engineer (Go)

Design and implement "Word of Wisdom" tcp server.

 - TCP server should be protected from DDOS attacks with the Proof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.
 - The choice of the POW algorithm should be explained.
 - After Proof Of Work verification, server should send one of the quotes from "word of wisdom" book or any other collection of the quotes.
 - Docker file should be provided both for the server and for the client that solves the POW challenge

## Explanation of the solution (choice of the POW algorithm)

The simple [Hashcash](https://en.wikipedia.org/wiki/Hashcash) option was chosen: it is one of the first PoW algorithms that was used to combat email spam.
The client manages to find a value (nonce), which, when combined with the data, creates a hash (SHA-256) starting with a definition of the numbers of zeros.
You can complicate the calculations by increasing the number of zeros.
Currently, 4 zeros are used for clarity, but you can use more zeros when searching for problems.

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

The file with quotes from "word of wisdom" book as a source of quotes can be setting in config file
```yaml
data_file: "config/server/wisdom.txt"`
```

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

## How to test

We use text file with quotes from "word of wisdom" book as a source of quotes.

### 1) Start server with docker-compose:

```
docker-compose -p "faraway_exercise" up -d
```

### 2) Run client from command line:

```shell
go run ./cmd/client -env=local -addr=127.0.0.1:8088 connect
```

### 3) See the result from several client runs:

We should get different results:

```
[11:12:12.049] INFO: Received PoW task: {
  "nonce": "65936"
}
Steps: 7638
[11:12:12.069] INFO: Found PoW solution: {
  "proof": "00003ce3da59631089ff666852ef79a02c4d67252f60ff793b77e22fb1be239e"
}
[11:12:12.070] INFO: Server response: {
  "response": "«Не помнящие прошлого обречены на его повторение» Джордж Сантаяна"
}
```
```
[11:44:44.268] INFO: Received PoW task: {
  "nonce": "21036"
}
Steps: 26481
[11:44:44.319] INFO: Found PoW solution: {
  "proof": "00009cd94f0b7c4de91f63c4e852425c22bf42d7355e9f69ba4df5a5a88a796d"
}
[11:44:44.319] INFO: Server response: {
  "response": "«Истинная ценность человека измеряется в тех вещах, к которым он стремится!»"
}
```
```
[11:12:12.773] INFO: Received PoW task: {
  "nonce": "20755"
}
Steps: 36718
[11:12:12.832] INFO: Found PoW solution: {
  "proof": "0000d3681725e7479ee6e1065b96257916e218b0dcdfb5df02ebd204d9e45b42"
}
[11:12:12.832] INFO: Server response: {
  "response": "«В основном свободу человек проявляет только в выборе зависимости». Герман Гессе"
}
```

Here we are generating random number (nonce) for each request, so we get different results.

```
nonce: 20755
```
And we need different steps to find solution
```
Steps: 36718
```

```
"proof": "0000d3681725e7479ee6e1065b96257916e218b0dcdfb5df02ebd204d9e45b42"
```

## Stress Testing

For stress testing of server we can use [Grafana k6](https://k6.io/) tool with TCP plugin [xk6-tcp](https://github.com/NAlexandrov/xk6-tcp).

For stress testing of client we use `expect` with script "test_client.sh":

For example
```shell
chmod +x test_client.sh

```

```shell
./test_client.sh 
```

Stress test
```shell
./stress_test_client.sh 
```