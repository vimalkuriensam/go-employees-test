package config

import (
	"context"
	"fmt"
	"math"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func (app *Config) MongoConnect() error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	user := app.Env["db_user"].(string)
	pass := app.Env["db_password"].(string)
	host := app.Env["dsn"].(string)
	dsn := fmt.Sprintf("mongodb://%v:%v@%v/?maxPoolSize=20&w=majority", user, pass, host)
	var counts int
	var err error
	var client *mongo.Client
	for {
		client, err = mongo.Connect(ctx, options.Client().ApplyURI(dsn))
		if err != nil {
			app.Logger.Print("Db not connected yet! Retrying...")
			counts++
			time.Sleep(time.Duration(math.Pow(float64(counts), 2)) * time.Second)
		} else {
			app.DataBase.Client = client
			if err = app.DataBase.Client.Ping(ctx, readpref.Primary()); err != nil {
				return err
			}
			app.Logger.Print("Db connected and pinged...")
			return nil
		}
		if counts > 10 {
			return err
		}
	}
}

func (app *Config) MongoDisconnect() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	if err := app.DataBase.Client.Disconnect(ctx); err != nil {
		app.Logger.Fatal(err)
	}
}

func (app *Config) InsertMongoCollections(collections ...string) {
	for _, value := range collections {
		go InsertMongoCollection(value, app.DataBase, app.Env["db_database"].(string))
	}
}

func InsertMongoCollection(collection string, db *DataBase, dbName string) {
	col := db.Client.Database(dbName).Collection(collection)
	db.Collections[collection] = col
	fmt.Printf("Added collection %s\n", collection)
}
