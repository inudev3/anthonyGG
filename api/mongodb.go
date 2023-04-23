package main

type MongoDb struct{}

func (s *MongoDb) Get(id int) *User {
	return &User{
		ID:   id,
		Name: "foo",
	}
}
func NewMongoStorage() *MongoDb {
	return &MongoDb{}
}
