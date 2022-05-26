# poc_rabbitmq



## _Let's test the RabbitMQ message broker_

[![N|Solid](https://cldup.com/dTxpPi9lDf.thumb.png)](https://nodesource.com/products/nsolid)

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://github.com/axtoneIO)

This is a golang project that is interacting with the rabbitmq message broker running on a docker container, the whole purpose of the project was to experiment with this technology..

The project covers the following tools:
## Tech

Dillinger uses a number of open source projects to work properly:

- [RabbitMQ](https://www.rabbitmq.com/) - Most widely deployed open source message broker
- [Docker](https://www.docker.com/) - Developers Love Docker
- [Golang](https://go.dev/) - An open source language supported by Google

## Installation

Dillinger requires [Golang](https://go.dev/dl/) v1.16+ to run.

Install the dependencies:

```sh
go mod tidy
go mod vendor
```
## Docker

Running the RabbitMQ services is pretty easy using docker, before trying to run
the project make sure to run the following command:

```sh
docker run -d --hostname rabbitmq --name test-rabbit -p 15672:15672 -p 5672:5672 rabbit
mq:3-management
```
This will pull the official rabbitmq image from Docker hub
the name assigned to our service in this case will be test-rabbit but you can change it..

For testing purposes you can access the container running using the interactive mode
```sh
docker exec -it {container id} bash
```
and then publish a message into your queue and see the result on your project running
```sh
rabbitmqadmin publish exchange=amq.default routing_key="TestQueue" payload="Hello World from bash"
```

## License

MIT

**Free Software, Hell Yeah!**