package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)


type apiFunc func(w http.ResponseWriter, r *http.Request)error

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
	fmt.Println("Server is running on the port",s.address)
	router.HandleFunc("/template/home",makeHttpHandlers(s.HandlerHomePage))
	router.HandleFunc("/template/signIn",makeHttpHandlers(s.HandleSignInPage))
	if err:=http.ListenAndServe(s.address,router);err!=nil{
		log.Fatal("Error : ",err)
	}
}

func(s *Server)HandlerHomePage(w http.ResponseWriter,r *http.Request)error{
	tmpl,err:=makeTemplate("templates/home.html")
		if err!=nil{
			return err
		}
		err=tmpl.Execute(w,nil)
		if err!=nil{
			return err
		}
	return nil
}
func(s *Server)HandleSignInPage(w http.ResponseWriter,r *http.Request)error{
	tmpl,err:=makeTemplate("templates/sign-in.html")
	if err!=nil{
		return err
	}
	err=tmpl.Execute(w,nil)
	if err!=nil{
		return err
	}
	return nil
}

func makeHttpHandlers(function apiFunc)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		err:=function(w,r)
		if err!=nil{
			WriteError(w,err)
		}
	}
}

func makeTemplate(templateFileLocation string)(*template.Template,error){
	tmpl,err:=template.ParseFiles(templateFileLocation,"templates/top.html","templates/bottom.html")
	if err!=nil{
		return nil,err
	}
	return tmpl,nil
}
func WriteError(w http.ResponseWriter, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ApiError{Error: err.Error()})
}

func WriteJson(w http.ResponseWriter, status int, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ResponseMessage{Message: message})
}
