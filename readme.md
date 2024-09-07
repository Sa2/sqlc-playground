



# set up flow

```
$ go install github.com/rubenv/sql-migrate/...@latest
$ go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```


# generate codes

```
$ docker run --rm -v $PWD:/src -w /src sqlc/sqlc generate
