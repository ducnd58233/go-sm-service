postgres:
		docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

mysql:
		docker run --name mysql8 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql:8 

test:
		go test -v -cover ./...