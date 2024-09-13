package main

import (
	"fmt"
	"log"
	"net/http"
)


type Server struct{
	address string
	store Store
}


func GetServer(address string ,store Store)*Server{
	return  &Server{
		address:address,
		store:store,
	}
}

func(s *Server)Run(){
	router:=http.NewServeMux()
	fmt.Println("Server is running on the port : ",s.address)
	router.HandleFunc("/health",func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("This is the health end point")
	})
	if err:=http.ListenAndServe(s.address,router);err!=nil{
		log.Fatal("Error : ",err)
	}
}