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

## workrecord-go起動（Local）

```bash
{$GOPATH}/bin/fresh
```

## workrecord-go起動（Docker）

```bash
docker image build -t workrecord-go:0.2.0 .
docker run -it --rm -p 8080:8080 --link mysqld --name workrecord-go workrecord-go:0.2.0
```
