PG_PORT=5433
MS_PORT=3307
PG_DB_URL=postgresql://root:secret@localhost:$(PG_PORT)/social_media?sslmode=disable
MS_DB_URL=mysql://root:secret@tcp(localhost:$(MS_PORT))/social_media

postgres:
		docker run --name postgres12 -p $(PG_PORT):5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

mysql:
		docker run --name mysql8 -p $(MS_PORT):3306 -e MYSQL_ROOT_USER=root -e MYSQL_ROOT_PASSWORD=secret -d mysql:8.0 

createpgdb:
		docker exec -it postgres12 createdb --username=root --owner=root social_media

createmsdb:
		docker exec -it mysql8 mysql --user=root --password=secret -e "CREATE SCHEMA social_media;"

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