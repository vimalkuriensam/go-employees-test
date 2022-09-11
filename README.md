# Employee REST-API

This project creates a rest-api server that deals with the basic CRUD operations of an
employee. The project is created using golang and the database used is mongodb.

## Run the app

Run the project with go version 1.18

```bash
  go run ./cmd/main/. 
```
Run the project on docker compose.
```bash
  docker compose -f ./deployment/docker-compose.yml up -d --build
```