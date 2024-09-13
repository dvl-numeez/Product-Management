package main




func main(){
	srv:=GetServer(":5050",struct{}{})
	srv.Run()
}
