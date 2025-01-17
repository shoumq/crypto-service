Команда для запуска:
```
docker-compose up --build
```

[POST] /currency/add - в тело передаем json по типу:
```
{
    "coin": "BTC"
}
```

[DELETE] /currency/remove - в тело передаем json по типу:
```
{
    "coin": "BTC"
}
```

[POST] /currency/price - в тело передаем json по типу:
```
{
    "coin": "BTC",
    "timestamp": 1737135391
}
```
