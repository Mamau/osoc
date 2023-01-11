## Osoc
### Запустить:
* Скопировать .env файл
```shell
cp example.env .env
```
* Будет выполнен запуск сервисов и выполнены миграции
```shell
make start
```
* В БД есть тестовый юзе с id 1
#### Затестить так
```shell
curl 'http://localhost:8081/external/api/v1/user/1?test=wtf'
```
просто ответит json  
потом можно посмотреть на метрики

Запустится docker-compose файлик
* http://localhost:9090/graph - Прометей (fetch_user_count_total - по этому ключу можно найти сколько запросов было на роут /user/:id)
* http://localhost:9100/metrics - Url где можно увидеть свои кастомные метрики
* http://localhost:9111/metrics - Url где можно увидеть метрики node-exporter
* http://localhost:16686/ - Трейсер Джагер
* Вести разработку лучше через команду будет запущен проект в режиме live reload, при каждом изменении кода - будет перезапускаться проект
```shell
make watch
```
#### Работа с миграциями
[Смотреть](docs/migrations.md)

#### Работа с pprof
[Смотреть](docs/pprof.md)
