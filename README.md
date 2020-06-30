# gin-practice

## 起動
```
$ docker-compose up
```

## マイグレーション
```
$ docker exec -it gin_app sh -c "cd /go/src/app; go run migrate/migrate.go"
```

## 確認
```
$ curl http://localhost:3001/ping
{"message":"pong"}
```

## 停止
```
$ docker-compose down
```

## TODO
