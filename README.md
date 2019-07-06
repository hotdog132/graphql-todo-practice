# Graphql todo practice
- Create graphql (https://github.com/99designs/gqlgen) gateway to get informations from internal services and response to client
- Internal service: user-service to store and fetch user data
- Internal service: event-service to store and fetch todo data
- Use gateway implemented by graphql to gather data from internal services and response to client

## Setup & launch

### Test
go test ./...

### Install goose
```
go get -u github.com/pressly/goose/cmd/goose
```

### Build images and create containers
```
docker-compose build
```
```
docker-compose up -d
```

### DB migration
```
goose -dir db/migrations postgres "user=test password=test dbname=todo sslmode=disable" up
```

## Gateway test
- Open "http://localhost:8080" to the navigation of graphql
- Create todo item using mutation "createTodo"
- View todo item using query "findTodos" 
