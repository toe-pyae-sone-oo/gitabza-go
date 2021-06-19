package mongodb

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const mongoConnectTimeout = 10 * time.Second

var (
	db          *mongo.Client
	connectOnce sync.Once
)

func Connect(ctx context.Context, addr string) (*mongo.Client, error) {
	var err error

	connectOnce.Do(func() {
		c, cerr := connect(ctx, addr)
		if err != nil {
			err = cerr
			return
		}

		db = c
	})

	return db, err
}

func connect(ctx context.Context, addr string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, mongoConnectTimeout)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(addr))
	if err != nil {
		return nil, err
	}

	log.Println("connecting to mongodb...")

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	log.Println("mongodb successfully connected")

	return client, err
}

func GetDB() *mongo.Client {
	if db == nil {
		panic("db is not connected yet")
	}
	return db
}

func Disconnect(ctx context.Context) error {
	if db == nil {
		return errors.New("db already disconnceted")
	}

	if err := db.Disconnect(ctx); err != nil {
		return err
	}

	db = nil

	return nil
}
