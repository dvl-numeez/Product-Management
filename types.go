package main


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