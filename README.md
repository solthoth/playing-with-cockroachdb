# playing-with-cockroachdb
Sample applications interacting with a CockroachDB backend

## Sample Startup
To begin testing everything, simply run

```
$ docker-compose up -d
$ docker exec -it roach1 ./cockroach init --insecure
```

This will create the cockaroachdb cluster and misc apps to talk to it.

## Database design
For detail information on how to spin up the database in docker, follow this link: https://www.cockroachlabs.com/docs/v21.2/start-a-local-cluster-in-docker-linux