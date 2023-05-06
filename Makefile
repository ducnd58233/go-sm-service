PG_DB_URL=postgresql://root:secret@localhost:5432/social_media?sslmode=disable
MS_DB_URL=mysql://root:secret@tcp(localhost:3306)/social_media?sslmode=disable

postgres:
		docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

mysql:
		docker run --name mysql8 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql:8 

createpgdb:
		docker exec -it postgres createdb --username=root --owner=root social_media

createmsdb:
		docker exec -it mysql createdb --username=root --owner=root social_media

migratepgup:
		migrate -path db/migration/postgresql -database "$(PG_DB_URL)" -verbose up

migratepgdown:
		migrate -path db/migration/postgresql -database "$(PG_DB_URL)" -verbose down

migratemsup:
		migrate -path db/migration/mysql -database "$(MS_DB_URL)" -verbose up

migratemsdown:
		migrate -path db/migration/mysql -database "$(MS_DB_URL)" -verbose down

test:
		go test -v -cover ./...