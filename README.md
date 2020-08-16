# go-rest-gorm-postgres

## Project Structure
- The basic structure layout is based on the example shown below. \
    https://github.com/golang-standards/project-layout
- Additionally apply Go adapter pattern

|Directory  |Description                                            |
|-----------|-------------------------------------------------------|
|adapter    |External application code.                             |
|app        |Internal application code.                             |
|cmd        |Main application for this project.                     |
|internal   |Private application and library code.                  |
|model      |Data structure and method of database.                 |
|pkg        |Library code that's ok to use by external applications.|
|vendor     |Application dependencies.                              |


## Port
``` :7481 ```

## Env
- Notes: <a href="./.env.example">Env file</a>

|환경변수 이름        |값              |
|-----------------|----------------|
|VERSION          |0.1.0           |
|BIND_ADDRESS     |localhost:7481  |
|POSTGRES_HOST    |localhost       |
|POSTGRES_PORT    |5432            |
|POSTGRES_USER    |paul            |
|POSTGRES_DB_NAME |postgres        |
|POSTGRES_PASSWORD|password        |
|ALLOW_ORIGINS    |http://localhost:3000 |
|ALLOW_HEADERS    |Origin,Authorization,Content-Type,Content-Length|
|ALLOW_METHODS    |GET,POST,PATCH,OPTIONS|

## API
- Notes: <a href="./API.md">API 문서</a>

## How to run
- local
    1. git clone https://github.com/SungminSo/go-rest-gorm-postgres.git
    1. ``` cd cmd/project ```
    1. ``` go run main.go ```