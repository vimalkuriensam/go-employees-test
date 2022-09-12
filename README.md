# Employee REST-API
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

This project creates a rest-api server that deals with the basic CRUD operations of an
employee. The project is created using golang and the database used is mongodb.

## Run the app

#### Run the project with go version 1.18

```bash
  go run ./cmd/main/. 
```
#### Run the project on docker compose.
```bash
  docker compose -f ./deployment/docker-compose.yml up -d --build
```
#### Run the project with makefile
Run the docker compose file
```bash
  make up
```
Remove the existing containers and rebuild the apps
```bash
  make up_build
```
Stop the running containers
```bash
  make down
```
Create a binary file
```bash
  make build_employees
```

## Environment Variables

To run this project, you will need to add the following environment variables to
either production.env or development.env in the deployment folder

`port` *-web server port*

`dsn` *-mongodb host and port eg: 127.0.0.1:27017*

`db_user` *-mongodb username*

`db_password` *-mongodb password*

`db_database` *-mongodb database name*

To run the database, you will need to add the following environment variables to
mongo.env file in deployment folder.

`MONGO_INITDB_ROOT_USERNAME` *-mongodb root username*

`MONGO_INITDB_ROOT_PASSWORD` *-mongodb root password*

`MONGO_INITDB_DATABASE` *-mongodb database name*

## Project Structure
```
├── cmd
│   ├── api
│        └── main.go // main file that initializes the app
├── pkg 
│   ├── config   
│   │   ├── config.go
|   |   ├── db.go  
│   │   └── utilities.go
|   ├── controllers
│   │   └── employees.controllers.go
|   ├── routes
│   │   ├── routes.go
│   │   └── employees.routes.go 
|   ├── models
│   │   └── employees.models.go  
│   └── services
|       ├── services.go  
│       └── employees.services.go
├── deployment
│   ├── Dockerfile.prod
│   └── docker-compose.yml
├── environment
│   ├── development.env
│   ├── production.env
│   └── mongo.env
├── go.mod
└── go.sum
```


## API Reference

#### Create an employee

```http
  POST /api/v1/employees/addEmployee
```

| Body      | Type   | Description                 |
| :-------- | :----- | :-------------------------- |
| name      | string | **Required** employees name |
| age       | int    | **Required** employees age  |
| email     | string | **Required** employees email|

#### Get employee with id

```http
  GET /api/v1/employees/getEmployee/${id}
```

| Parameter | Type     | Description                          |
| :-------- | :------- | :----------------------------------- |
| `id` | `string` | **Required**. Id of the employee to fetch |

#### Update employee with id

```http
  PATCH /api/v1/employees/updateEmployee/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

| Body      | Type   | Description                 |
| :-------- | :----- | :-------------------------- |
| name      | string | **Optional** employees name |
| age       | int    | **Optional** employees age  |
| email     | string | **Optional** employees email|

#### Delete employee with id

```http
  DELETE /api/v1/employees/deleteEmployee/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |


