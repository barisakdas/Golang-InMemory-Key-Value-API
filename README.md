# Golang In Memory Key-Value Store API

[![go-github docker-image (latest SemVer)](https://img.shields.io/github/v/release/google/go-github?sort=semver)](https://hub.docker.com/repository/docker/barisakdas/golang-key-value-app)

An api project that handles the key-value data type in memory.
Almost all of the project was written using standard packages.
You can see, add and delete all keys and values using api endpoints.

The data added to the memory within the project are kept by writing to the files at certain intervals.
At the same time, triggered handler functions and other function logs are written into separate files.

When the project runs for the first time, it processes the data saved in `Data.txt` and adds it to the memory.

The project works as an api and the functions that are triggered at a certain time interval work in the background. These simultaneous functions do not harm each other's work cycles.
Only 3rd party applications are used to write functions that need to run at certain intervals.

```go
module key-value-store

go 1.16

require (
github.com/jasonlvhit/gocron v0.0.1 // indirect
)
```

## Installation

The program can be run easily by downloading the publicly published image on `hub.docker.com`:

```bash
docker pull barisakdas/golang-key-value-app
```

If you want to run the codes by getting the codes on github, the project can be reached by pulling the repository as follows.

```go
git clone https://github.com/barisakdas/ys.git
```

## Usage

When the API is run, detailed information about all open endpoints can be obtained by sending a request to the `http://localhost:8080/` address first.

`https://golang.org/doc/` has been used as the main auxiliary source for the developments within the project.

We tried to use the `dependency injection design pattern` in the project.
The application has been containerized by creating a docker image.

```bash
curl http://localhost:8080/
```

Afterwards, the relevant endpoints can be called and the desired operations can be performed from the API.

## Future

By adding a database in addition to the application, the data can be saved in the database instead of files.
For logs, a 'fluent bit' application and an 'elasticsearch' database can be added to ensure a healthier operation.
In the second version, we will add the above plugins to this project and share the `docker-compose.yaml` file.
