# avito-grpc
**pet project for grpc practice**
## descripion of version
you can see 2 version of this project
1. grps version
    - on main branch
1. version on fiber
    - on fiber-app-relise

## Usage
### Запросы обрабатываемые сервисом
`POST /create body{"url": "job.ozon.ru"}`\n
`GET /get/{hash}`
### Примеры запросов
```shell
#POST
curl -X POST localhost:8080/create -H "Content-Type: application/json" -d '{"url": "job.ozon.ru"}'
#GET
curl -X POST localhost:8080/get/someHash'