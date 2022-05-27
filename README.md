# service-golang

This is a simple example to build a microservice in Go. It has an HTTP server and RabbitMQ client packages.

## Running it

The service needs MongoDB and RabbitMQ to run succesfully. Let's use `docker-compose` to run those servers.

```bash
docker-compose up -d mongodb rabbitmq
```

And now you can build and run the server:

```bash
go build && ./service-golang
```

## Creating and image with docker

To run as a service we can create an image and use docker-compose to start the containers. First the image:

```bash
docker build . -t service:latest
```

Now we can run all the servers together using `docker-compose`.

```bash
docker-compose up -d
```

## Accessing the API

To see a Hello World: [localhost:8000](http://localhost:8000).

To check the bitcoin variation on a period: [localhost:8000/bitcoin/startdate/YYYY-MM-DD/enddate/YYYY-MM-DD](http://localhost:8000/bitcoin/startdate/2018-11-01/enddate/2018-11-30)

## Stopping the services

Again with `docker-compose`:

```bash
docker-compose down
```
