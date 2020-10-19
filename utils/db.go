package utils

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var ID []string

type DB struct {
	client *mongo.Client
}

func (C *DB) Access(url string) {
	log.Println("Starting DB")
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Error while connecting DB : %v", err)
		return
	}
	C.client = client

}

func (C *DB) GetAllhash() {
	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()
	collection := C.client.Database(GetDbName()).Collection(GetDbCollection())
	all, err := collection.Find(ctx, bson.D{})
	if err != nil {
		fmt.Println(err)
	}
	defer all.Close(ctx)
	for all.Next(ctx) {
		var result bson.M
		err = all.Decode(&result)
		if err != nil {
			log.Println(err)
		} else {
			if result["hash"] != nil {
				hash := result["hash"]
				ID = append(ID, fmt.Sprintf("%v", hash))
			}
		}
	}

}

func (C *DB) Insert(hash string) {
	collection := C.client.Database(GetDbName()).Collection(GetDbCollection())
	j := bson.M{"hash": hash}
	_, err := collection.InsertOne(context.Background(), j)
	if err != nil {
		log.Printf("Error inserting the id into db %s", err)
		return
	}
	ID = append(ID, hash)
	log.Printf("Added hash to DB : %v", hash)

}

func CheckValid(hash string) bool {
	for _, i := range ID {
		if i == hash {
			return true
		}
	}
	return false
}

func getIdIndex(userId string) (int, bool) {
	for i, j := range ID {
		if j == userId {
			return i, true
		}
	}
	return 0, false
}

func (C *DB) Delete(hash string) bool {
	index, found := getIdIndex(hash)
	if found {
		ID[index] = ID[len(ID)-1] // Copy last element to index i.
		ID[len(ID)-1] = ""        // Erase last element (write zero value).
		ID = ID[:len(ID)-1]       // Truncate slice.
		return false
	}
	collection := C.client.Database(GetDbName()).Collection(GetDbCollection())
	_, err := collection.DeleteOne(context.Background(), bson.M{"hash": hash})
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
