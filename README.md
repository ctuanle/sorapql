## Setup

Visit and testing [sorapql](https://sorapql.fly.dev/)

```shell
go run github.com/99designs/gqlgen generate
go run server.go
```

## Deploy

```shell
go run github.com/99designs/gqlgen generate
flyctl secrets set KEY=VALUE
flyctl deploy
```
