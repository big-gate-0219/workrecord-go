# workrecord-go

## データベースの起動

```bash
docker pull mysql:5.7

docker run --name mysqld \
           -e MYSQL_DATABASE=workrecord \
           -e MYSQL_USER=default \
           -e MYSQL_PASSWORD=password  \
           -e MYSQL_ROOT_PASSWORD=mysql \
           -p 3306:3306 \
           -d \
           mysql:5.7 \
           --character-set-server=utf8mb4 \
           --collation-server=utf8mb4_general_ci
```

## データベースの初期化

``` bash
cd src
go run ./tools/migration.go
```

## workrecord-goの起動

```bash
{$GOPATH}/bin/fresh
```

## Memo

```bash
docker image build -t workrecord-go:0.1.0 .
docker run -it --rm --name go-realizer -p 8080:8080 -v ./src:/go/src/app workrecord-go:1.0
```
