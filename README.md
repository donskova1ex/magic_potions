# Go API Server for openapi

Сервис изготовления зелий

## Overview
This server was generated by the [openapi-generator]
(https://openapi-generator.tech) project.
By using the [OpenAPI-Spec](https://github.com/OAI/OpenAPI-Specification) from a remote server, you can easily generate a server stub.

To see how to make this your own, look here:

[README](https://openapi-generator.tech)

- API version: 1.0.0
- Build date: 2025-02-06T09:01:32.518904300+05:00[Asia/Yekaterinburg]
- Generator version: 7.10.0


### Running the server
To run the server, follow these simple steps:

```
go run main.go
```

The server will be available on `http://localhost:8080`.

To run the server in a docker container
```
docker build --network=host -t openapi .
```

Once image is built use
```
docker run --rm -it openapi
```
