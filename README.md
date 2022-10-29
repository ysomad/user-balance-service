# Avito internship task - user wallet service

## Local development
1. Run all containers except applicaton one:
```sh
$ make compose-up
```

2. Compile and run the app with migrations
```sh
$ make run-migrate
```
Cmd from above will run `go run` with build tag `migrate`.

Or if you don't need run `migrate up` on startup:
```sh
$ make run
```

## Run app in docker
To run the app in fully in docker:
```sh
$ make compose-all
```
Makefile cmd from above will run all containers from `docker-compose.yml`

### Stop containers
To stop all containers:
```sh
$ make compose-down
```

### Questions
1. В каком формате хранить баланс? 
Сначала решил хранить в наименьшем эквиваленте (копейки, центы например), т.к. в go нету реализации decimal, по API принимать строку и хранить в бд в типе bigint.
Но затем нашел библиотеку github.com/shopspring/decimal реализуюущую decimal на основе big.Int и решил использовать ее.