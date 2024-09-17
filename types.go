package main

import "go.mongodb.org/mongo-driver/bson/primitive"


type User struct{
	Email string`json:"email"`
	Password string `json:"password"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

type ApiError struct {
	Error string
}

type HomePageInfo struct{
	Username string
	Products []Product
}
type Product struct{
	Id primitive.ObjectID `json:"id" bson:"_id"`
	ProductName string `json:"productName" bson:"product_name"`
	ProductType string `json:"productType" bson:"product_type"`
	ProductMaterial string `json:"productMaterial" bson:"product_material"`

}