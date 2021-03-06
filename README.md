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
export BODY_OK2='{"title":"Black Panther","year":2018,"runtime":"134 mins","genres":["action","adventure"]}'
export BODY_OK3='{"title":"Deadpool","year":2016, "runtime":"108 mins","genres":["action","comedy"]}'
export BODY_OK4='{"title":"The Breakfast Club","year":1986, "runtime":"96 mins","genres":["drama"]}'

export BODY_OK2a='{"title":"Black Panther","year":2018,"runtime":"134 mins","genres":["sci-fi","action","adventure"]}'

curl -i -d "$BODY" localhost:4000/v1/movies

curl -X DELETE localhost:4000/v1/movies/3

curl -w '\nTime: %{time_total}s \n' localhost:4000/v1/movies/1

BODY_USER='{“name": "Alice Smith", "email": "alice@example.com", "password": "pa55word"}'

curl -i -d "$BODY_USER" http://localhost:4000/v1/users

BODY_USER_INVALID='{"name": "", "email": "bob@invalid.", "password": "pass"}'

curl -i -d "$BODY_USER_INVALID" http://localhost:4000/v1/users

BODY_USER_DUPLICATE='{"name": "Alice Jones", "email": "alice@example.com", "password": "pa55word"}'
```

### DB Migrations

```
migrate create -seq -ext=.sql -dir=./migrations create_movies_table

migrate create -seq -ext=.sql -dir=./migrations add_movies_check_constraints

migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up

migrate create -seq -ext .sql -dir ./migrations add_movies_indexes
migrate -path ./migrations -database $GREENLIGHT_DB_DSN up
```
