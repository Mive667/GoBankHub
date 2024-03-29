# GoBankHub

Tools:
dbdiagram.io
Docker

## Database Configuration


### Database Schema
![DB Schema](assets/DB_Schema.svg)

### Postgres Docker
[Documents of Postgres](https://hub.docker.com/_/postgres?uuid=3f9a1944-f4da-4357-ae4d-9ba762d1c63d%0A)
Use the command to get a lightweight image, it is good enough for this mini-project.
```zsh
docker pull postgres:16-alpine
``` 

Then start a postgres instance by:
```zsh
docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root_user -e POSTGRES_PASSWORD=123456 -d postgres:16-alpine
```

```zsh
docker exec -it postgres16 psql -U root_user
```

### DB migration in Go
```zsh
brew install golang-migrate
```

``` zsh
migrate create -ext sql -dir db/migration -seq init_schema
```

I have created a `Makefile` for convenience, by runnnig command below(in the root directory) to run a contanier and create db:
``` zsh
make postgres
```
```zsh
make createdb
```

then run the command to make a migration:
``` zsh
migrate -path db/migration -database "postgresql://root_user:123456@localhost:5432/gobankhub?sslmode=disable" -verbose up
```

`go mod tidy` to update the dependencies.

## CI/CD
### Postgres service container
[Docs](https://docs.github.com/en/actions/using-containerized-services/creating-postgresql-service-containers)

## RESTful Http Json API

### Viper
[Docs](https://pkg.go.dev/github.com/spf13/viper)
Configur the environment files.

### mockgen
mockgen -package mockdb -destination db/mock/store.go github.com/Mive667/GoBankHub/db/sqlc Store