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
