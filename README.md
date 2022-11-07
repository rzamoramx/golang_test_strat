# Golang test Strat

## Requirements

* Go 1.18+


## Steps

* Clone project
* Located in project root, run command: 
``` go mod tidy ```
* Finally run: ```go run .```


## Considerations

The sqlite users table in database is dropped on every start of the server

The endpoints to test are:
* localhost:8080/v1/register
* localhost:8080/v1/login