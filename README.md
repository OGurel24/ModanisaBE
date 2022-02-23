# What is the project
It is the backend of the assigment project, Golang back end of the do list with Go standard http module

## Project setup (Dockerized)
https://hub.docker.com/r/onurgurel24/backend

After pulling the image run given command (please DO NOT change ports)
```
docker run -d -p 8081:8081 onurgurel24/backend
```

## Project setup (Manual Version)
Install all dependencies with given command
```
go mod download
go run .
```

### Run all the tests
```
go test ./... -cover
```
