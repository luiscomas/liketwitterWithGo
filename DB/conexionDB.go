package db

import (
	"context"
	"fmt"

	"github.com/luiscomas/liketwitterWithGo/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCn *mongo.Client
var DatabaseName string

func MongoConnect(ctx context.Context) error {
	user := ctx.Value(models.Key("user")).(string)
	password := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)
	connstr := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, password, host)

	var clientOptions = options.Client().ApplyURI(connstr)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Connected to MongoDB")
	MongoCn = client

	db := ctx.Value(models.Key("database")).(string)
	DatabaseName = db

	return nil

}

func ConnectedDatabase() bool {
	err := MongoCn.Ping(context.TODO(), nil)
	return err == nil
}
