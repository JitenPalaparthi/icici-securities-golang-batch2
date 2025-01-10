go get -u github.com/gin-gonic/gin


docker network create demo-network

docker run --name pg -d -p 5432:5432  -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin123 -e POSTGRES_DB=userdb --network=demo-network postgres  

docker run --name mysql -d -p 3306:3306 -e MYSQL_USER=admin -e MYSQL_PASSWORD=admin123 -e MYSQL_ROOT_PASSWORD=admin123 -e MYSQL_DATABASE=userdb --network=demo-network mysql

docker run --name dbui -d -p 8089:8080 --network=demo-network adminer


### to run all test in a package

```bash
go test ./...
```

```
go test -timeout 30s -run ^TestValidateUserSuccess$ demo/models
go test demo/models -v -cover
go test demo/models --coverprofile=cover.o
```