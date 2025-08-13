# Gator (RSS Feed)
This is a project I had to make during my Boot.Dev courses. It's a terminal based RSS fetcher

# Dependices
- Postgresql
- [Goose](https://pressly.github.io/goose/))
    - Install with go install: `go install github.com/pressly/goose/v3/cmd/goose@latest`
- [sqlc](https://sqlc.dev)
    - Install with go install: `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`
- Google's UUID library
    - Install with go get: `go get github.com/google/uuid`
- Go Postgresdriver:
    - Install with go get: `go get github.com/lib/pq`

# Instructions
I assume when you develop on this project that the dependecies are installed, with a database called gator 
First off in the user's home directory there ought to be a file called .gatorconfig.json
It should contain the following, provided you have:
```json
{"db_url":"postgres:///[userspec@][hostspec][/dbname][?paramspec]/gator?","name":""}
```
[Check this section of the psql docs](https://www.postgresql.org/docs/12/libpq-connect.html#id-1.7.3.8.3.6) to learn more about the URI connector
<br>
Navigate to the sql/schema directory and run goose to have the schematic up in the database:
```sh
goose postgres "postgres://username:password@ip/gator" up
```

Once you have written new queries in sql/queries, generate the proper model files by using
```sh
sqlc generate
```