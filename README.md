# Inaccess Assignment

By Theodore Mandilaras

## Setup

To start the application run the below commands:
```shell script
// to download dependencies
go mod download

//to execute either with 
go run cmd/main/main.go

//or by building it first
go build cmd/main/main.go
./main
```

By default the application listens to the port 8080. To change it run like:
```shell script
go run cmd/main/main.go -port 8000

// or

./main -port 8000
```

### Setup with docker

For easier deployment, a dockerfile has been implemented. You can setup the application by

```shell script
// to build
docker build --tag go-inaccess . 

// to run
docker run -p <port>:8080 --rm --name go-inaccess go-inaccess
```

Just replace the `<port>` with the desired port to listen the application.
