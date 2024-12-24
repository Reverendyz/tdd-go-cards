package db

import (
	"errors"
	"fmt"
	"log"
	"reverendyz/tdd-go-cards/pkg/common"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	Client   *mongo.Client
	Database *mongo.Database
	uri      string = fmt.Sprintf("mongodb://%s:%s@%s:%s",
		common.GetEnvOrFallback("MONGO_USER", "root"),
		common.GetEnvOrFallback("MONGO_PASS", "root"),
		common.GetEnvOrFallback("MONGO_HOST", "127.0.0.1"),
		common.GetEnvOrFallback("MONGO_PORT", "27017"),
	)
)

func GetClient() *mongo.Client {
	fmt.Print(uri)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		log.Fatal(errors.New("could not connect to database. check if database is currently running"))
	}
	return client
}
