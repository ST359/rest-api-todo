- [Screenshots](#Screenshots)
- [English](#Running)
  - [Running](#Running)
  - [Description](#Description)
  - [Request examples](#Request_examples)
- [Русский](#Запуск)
  - [Запуск](#Запуск)
  - [Описание](#Описание)
  - [Примеры запросов](#Примеры_запросов)

## Screenshots
[Example of task creation](assets/postman_screenhot_1.jpg "Example of task creation")
[Example of getting all tasks list](assets/postman_screenhot_2.jpg "Example of getting all tasks list")

## Running
To run the service, Docker and Docker Compose are required.
```sh
git clone https://github.com/ST359/todo-rest-api
cd todo-rest-api
docker-compose up --build
```

## Description
Sample TODO backend service with REST API, created with Fiber, Postgres and pgx

## Request_examples
`POST /tasks` creates a new task and returns created task ID
```
{
    "title": "testtask1",
    "description": "test description"
}
```
`PUT tasks/:id` updates existing task where task id=id
```
{
    "title": "changed via PUT",
    "description": "changed via PUT",
    "status": "in_progress"
}
```
`GET /tasks` returns a list of all existing tasks
`GET /tasks/:id` returns a task where task id=id
`DELETE /tasks/:id` deletes existing tasks where task id=id

## Запуск
Для запуска требуется Docker и Docker-compose
```sh
git clone https://github.com/ST359/todo-rest-api
cd todo-rest-api
docker-compose up --build
```

## Описание
Пример TODO backend сервиса с REST API, создан с использованием Fiber, Postgres и pgx

## Примеры_запросов
`POST /tasks` создает новую задачу и возвращает ID созданной задачи
```
{
    "title": "testtask1",
    "description": "test description"
}
```
`PUT tasks/:id` обновляет существующую задачу с указанным ID
```
{
    "title": "changed via PUT",
    "description": "changed via PUT",
    "status": "in_progress"
}
```
`GET /tasks` возвращает список всех существующих задач
`GET /tasks/:id` возвращает задачу с указанным ID
`DELETE /tasks/:id` удаляет задачу с указанным ID
