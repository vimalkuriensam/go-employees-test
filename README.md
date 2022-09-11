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


