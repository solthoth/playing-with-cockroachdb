# playing-with-cockroachdb
Sample applications interacting with a CockroachDB backend

## Sample Startup
To begin testing everything, simply run

```
$ docker-compose build
$ docker-compose up -d roach_alb
$ docker exec -it roach1 ./cockroach init --insecure
$ docker-compose up roach_test_client
```

This will create the cockaroachdb cluster and misc apps to talk to it.

### Test Connection Manually
```
$ cd app
$ go mod tidy
$ go run main.go
```

Use connection string `postgresql://root@localhost:26257/defaultdb?sslmode=disable`

## Database design
For detail information on how to spin up the database in docker, follow this link: https://www.cockroachlabs.com/docs/v21.2/start-a-local-cluster-in-docker-linux