# api

# To run localy

1. run Mysql DB localy
```
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql
```
if you use DBeaver Community -> open new script, following Banking/dataForDB.txt and copy to clipboard code, paste and run with DBeaver
```
Banking/dataForDB.txt
```
2. install dependencies
```
go mod tidy
```
3. run categories service 
```
go run main.go
```
