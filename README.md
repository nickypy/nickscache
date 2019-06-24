# NC (Nick's Cache)
Actual name TBD

## Getting started
### Building and running server:
```
$ go build
$ ./main
```

Using Docker:
```
$ docker build --tag=nickscache .
$ docker run -p 8080:8080 nickscache
```

This will run a server on `localhost:8080`. 

### API Endpoints
Endpoints are currently available in Insomnia format in `Insomnia.yaml`

### Running unit tests
```
$ go test *.go
```

