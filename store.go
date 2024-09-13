package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


type Store interface{
	UserManagement
	ProductManagement
	
}

type UserManagement interface{
	InserUser(context.Context,*User)error
}
type ProductManagement interface{

}

type MongoStore struct{
	datbase *mongo.Database
}

func GetStore(ctx context.Context)(*MongoStore,error){
	godotenv.Load()
	url:=os.Getenv("MONGO_URL")
	client,err:=mongo.Connect(ctx,options.Client().ApplyURI(url))
	if err!=nil{
		return nil,err
	}
	err=client.Ping(ctx,readpref.Primary())
	if err!=nil{
		return nil,err
	}
	db:=client.Database("Product Management")
	return &MongoStore{
		datbase: db,
	},nil
}

func(store *MongoStore)InserUser(ctx context.Context,user *User)error{
	return nil
}