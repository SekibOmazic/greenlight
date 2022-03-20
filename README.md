# Greenlight

Code for "Let's go further" book

# How to run?

First you'll need to export following variable to your environment:

```
export GREENLIGHT_DB_DSN=postgres://greenlight:pa55word@localhost/greenlight?sslmode=disable
```

Then start the PostgreSQL in docker container with:

```
docker-compose up -d
```

and finally start the server with:

```
go run ./cmd/api
```

## Useful variables for testing

```
export BODY='{"title":"","year":1000,"runtime":"-123 mins","genres":["sci-fi","sci-fi"]}'

export BODY_OK='{"title":"Moana","year":2016,"runtime":"107 mins","genres":["animation","adventure"]}'

curl -i -d "$BODY" localhost:4000/v1/movies
```

### DB Migrations

```
migrate create -seq -ext=.sql -dir=./migrations create_movies_table

migrate create -seq -ext=.sql -dir=./migrations add_movies_check_constraints

migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up
```
