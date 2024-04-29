Add module to go.mod
```
go get github.com/joho/godotenv
```

Add code to vendor directory
```
go mod vendor
```

Install Tools related to DB 
```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Running a migration
```
cd sql/schema
goose postgres postgres://admin:test@localhost:5432/test up
```
2024/04/28 18:04:54 OK   001_users.sql (5.94ms)
2024/04/28 18:04:54 goose: successfully migrated database to version: 1
```

Using SQLC , we would generate go code for the query in sql/queries/users.sql  (Raw SQL)
```
sqlc generate
```
The go code lives in internal/database
