# Inaccess Assignment

By Theodore Mandilaras

- The JSON/HTTP service was implemented using the [mux](https://github.com/gorilla/mux) package.
- For the prediction of the matching timestamps of the periodic tasks the `time` package was used. 

## Project Structure

- `cmd`: Contains the files that are executed by the command line.
    - `main`: Main package.
        - `main.go`: The main file of the application. It starts the JSON/HTTP server. As the default port, it uses the `8080`.
    - `test`: Test package.
        - `test.go`: The implementation of the tests. 
 - `pkg`: Contains the packages that were implemented for the application needs.
    - `structs`: This package has all the required structs.
        - `error.go`: A struct with the error response that is returned to the client on the bad requests.
        - `tasks.go`: Contains the structure and some methods of the periodic tasks.
        - `testTask.go`: Contains the task structure that is used for the testing.
    - `utils`: Utilities Package
        - `utils.go`: Contains some helper functions.
- `Dockerfile`: Contains the instructions for building the Docker Image of the application.

## Setup

To start the application run the below commands:
```shell script
// download dependencies
go mod download

// execute either with 
go run cmd/main/main.go

//or by building it first
go build cmd/main/main.go
./main
```

By default the application listens to http://localhost:8080. To change it provide the bellow arguments:

- `-port`: to set a different port.
- `-addr`: to set a different address. However, since there is no webserver to set any domain, **only localhost and relevant IPs will work**.

```shell script
go run cmd/main/main.go -port <port> -addr http://0.0.0.0

// or

./main -port <port>
```

### Setup with Docker

For easier deployment, a docker-approach has been developed. You can setup the application by:

```shell script
// build
docker build --tag go-inaccess .

// run
docker run -p <port>:8080 --rm --name go-inaccess go-inaccess
```

Just replace the `<port>` with the desired port to listen the application.

## Testing

To test the application runs:

```
//without docker 
go run cmd/test/test.go

//or with docker
docker exec -it go-inaccess go run cmd/test/test.go 
```

If the application runs in a different port from the default, you can change it by providing the new port as a command line argument.

However, this cannot happen in the docker-approach, since in the docker container the port is fixed at 8080.

