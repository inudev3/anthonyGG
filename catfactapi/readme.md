```
docker run --name mongo -p 27017:27017 -d mongo
```

```
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/bson
```

```go
client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
if err!=nil{
	panic(err)
}
```
