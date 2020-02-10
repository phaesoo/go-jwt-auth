# Go JWT Auth

JWT auth backend project with Golang


## Debugging/Testing

Run debugging server

```console
$ go run cmd/auth/main.go
```

Unit test

```console
$ go test ./... -v
```


## Modules in use

- Router: "github.com/gorilla/mux"
- ORM: "github.com/jinzhu/gorm"


## Test environment

- OS: Ubuntu 18.04
- DB: sqlite3


## References

REST API
- https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj
- https://www.golangprograms.com/golang-restful-api-using-grom-and-gorilla-mux.html
JWT
- https://bourbonkk.tistory.com/60