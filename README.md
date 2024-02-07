# api

# To run localy

1. run MongoDB localy
```
docker run -d  --name mongo  -p 27888:27017 -e MONGO_INITDB_ROOT_USERNAME=phoqer -e MONGO_INITDB_ROOT_PASSWORD=change-me mongo
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
