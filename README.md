# Stache
A remote LRU cache written in Go.

## Usage
### Building and running server:
```
$ go build
$ ./stache
```

### Using Docker:
This will run a server on `localhost:8080`. 
```
$ docker build --tag=stache .
$ docker run -p 8080:8080 stache
```

A docker-compose file is also provided for convenience.


### API Endpoints
Endpoints are currently available in Insomnia format in `Insomnia.yaml`

### Running unit tests
```
$ go test
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)