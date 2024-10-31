# **Crapsolver backend API**

## **Running SurrealDB**

### Docker

```
docker run --rm --pull always -p 8000:8000 -v /mydata:/mydata surrealdb/surrealdb:1.0.0-beta.9-20230402 start --log trace --user root --pass rootnikoontoplmao5245
```

### Windows
```
surreal start --user root --pass rootnikoontoplmao52 --bind 0.0.0.0:8080 file:mydatabase.db
```

My usage:
on a admin powershell run: 
docker run --rm --pull always -p 8000:8000 -v /mydata:/mydata surrealdb/surrealdb:1.0.0-beta.9-20230402 start --log trace --user root --pass rootnikoontoplmao5245

then just do go run . to run main.go

still have to figure out long run api calls