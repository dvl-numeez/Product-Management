package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
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
	VerifyUser(context.Context,string)(string,error)
}
type ProductManagement interface{
	GetProducts(context.Context)([]Product,error)
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
	db:=client.Database("Product-Management")
	return &MongoStore{
		datbase: db,
	},nil
}

func(store *MongoStore)InserUser(ctx context.Context,user *User)error{
	_,err:=store.datbase.Collection("Users").InsertOne(ctx,user)
	if err!=nil{
		return err
	}
	return nil

}
func(store *MongoStore)VerifyUser(ctx context.Context,email string)(string,error){
	var user User
	coll:=store.datbase.Collection("Users")
	result:=coll.FindOne(ctx,bson.M{
		"email":email,
	})
	err:=result.Decode(&user)
	if err !=nil{
		return "",nil
	}
	return user.Password,nil

}

func(store *MongoStore)GetProducts(ctx context.Context)([]Product,error){
	var products []Product
	coll:=store.datbase.Collection("Products")
	result,err:=coll.Find(ctx,bson.M{})
	if err!=nil{
		return nil,err
	}
	for result.Next(ctx){
		var product Product
		err:=result.Decode(&product)
		if err!=nil{
			return nil,err
		}
		products = append(products, product)
	}
	return products,nil
}