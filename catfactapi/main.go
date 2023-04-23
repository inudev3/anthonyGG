package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

type CatFact struct {
	Fact   string `bson:"fact" json:"fact"`
	Length int    `bson:"length" json:"length"`
}
type Storer interface {
	GetAll() ([]*CatFact, error)
	Put(*CatFact) error
}
type MongoStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func (store *MongoStore) GetAll() ([]*CatFact, error) {

	query := bson.M{}
	cursor, err := store.coll.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	results := []*CatFact{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil

}
func (store *MongoStore) Put(fact *CatFact) error {

	_, err := store.coll.InsertOne(context.TODO(), fact)
	if err != nil {
		return err
	}
	return nil
}

func NewMongoStore() (*MongoStore, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	coll := client.Database("catfact").Collection("facts")
	return &MongoStore{
		client: client,
		coll:   coll,
	}, nil
}

type Server struct {
	store Storer
}

func NewServer(store Storer) *Server {
	return &Server{store: store}
}
func (s *Server) handleGetFact(w http.ResponseWriter, r *http.Request) {
	catfacts, err := s.store.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(catfacts)
}

type CatFactWorker struct {
	store    Storer
	endpoint string
}

func NewCatFactWorker(store Storer, endpoint string) *CatFactWorker {
	return &CatFactWorker{store: store, endpoint: endpoint}
}
func (cw *CatFactWorker) Start() error {
	ticker := time.NewTicker(time.Second * 2)
	for {
		res, err := http.Get(cw.endpoint)
		if err != nil {
			return nil
		}
		var catFact *CatFact //unordered map[string]any
		if err := json.NewDecoder(res.Body).Decode(&catFact); err != nil {
			return err
		}
		if err := cw.store.Put(catFact); err != nil {
			return err
		}
		<-ticker.C
	}
}
func main() {
	mongostore, err := NewMongoStore()
	if err != nil {
		log.Fatal(err)
	}
	worker := NewCatFactWorker(mongostore, "https://catfact.ninja/fact")
	go worker.Start()
	server := NewServer(mongostore)
	http.HandleFunc("/facts", server.handleGetFact)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
