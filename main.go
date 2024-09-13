package main

import (
	"context"
	"log"
)




func main(){
	ctx:=context.Background()
	store,err:=GetStore(ctx)
	if err!=nil{
		log.Fatal(err)
	}
	srv:=GetServer(": 5050",store)
	srv.Run()
}
