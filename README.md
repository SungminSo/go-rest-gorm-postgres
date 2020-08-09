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
