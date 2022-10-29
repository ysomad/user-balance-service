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
Т.к. в go нет decimal решил хранить в наименьшем эквиваленте  (копейки, центы), есть сторонние реализации, но решил не тащить лишнюю зависимость.
2. Откуда лучше вызывать транзакции? 
Изначально думал будет хорошим решением вызывать транзакции только из репозиториев, но хотелось добиться простоты доступа к данным (1 репозиторий - 1 таблица БД) и иметь возможность выполнять методы из разных репозиториев атомарно.
// TODO: дописать вопрос
3. Стоит ли хранить uuid в string, чтобы не нести зависимость в бизнес-логику? Если идеально следовать чистой архитектуре, то наверное да, но я в угоду удобства и в меру того, что uuid достаточно стандартизирован, решил использовать uuid тип там, где мне нужно чтобы он был.

### TODO
1. Метод резервирования средств с основного баланса на отдельном счете. Принимает id пользователя, ИД услуги, ИД заказа, стоимость.
2. Метод признания выручки – списывает из резерва деньги, добавляет данные в отчет для бухгалтерии. Принимает id пользователя, ИД услуги, ИД заказа, сумму.
3. Метод получения баланса пользователя. Принимает id пользователя.

5. Вынести сгенеренный код в отдельный пакет internal/gen/oapi/v1
6. Навести порядок в handler/v1