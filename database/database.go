package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"gogql/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	client *mongo.Client
}

func Connect(dbUrl string) *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		client: client,
	}
}

func (db *DB) InsertProductById(product model.NewProduct) *model.Product {
	productColl := db.client.Database("graphql-mongodb-api-db").Collection("product")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	inserg, err := productColl.InsertOne(ctx, bson.D{{Key: "name", Value: product.Name}})

	if err != nil {
		log.Fatal(err)
	}

	insertedID := inserg.InsertedID.(primitive.ObjectID).Hex()
	returnProduct := model.Product{ID: insertedID, Name: product.Name}

	return &returnProduct
}

func (db *DB) FindProductById(id string) *model.Product {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	productColl := db.client.Database("graphql-mongodb-api-db").Collection("product")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res := productColl.FindOne(ctx, bson.M{"_id": ObjectID})

	product := model.Product{ID: id}

	res.Decode(&product)

	return &product
}

func (db *DB) All() []*model.Product {
	productColl := db.client.Database("graphql-mongodb-api-db").Collection("product")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := productColl.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var products []*model.Product
	for cur.Next(ctx) {
		sus, err := cur.Current.Elements()
		fmt.Println(sus)
		if err != nil {
			log.Fatal(err)
		}

		product := model.Product{ID: (sus[0].String()), Name: (sus[1].String())}

		products = append(products, &product)
	}
	return products
}
