package main

import (
	"context"
	"testing"
)


func TestVerifyUser(t *testing.T) {
	ctx:=context.Background()
	store,err:=GetStore(ctx)
	if err!=nil{
		t.Error(err)
	}
	cases:=[]struct{
		email string
		password string
	}{
		{email: "numeez@gmail.com",
	password:"123456"},
	{email: "koko@gmail.com",
password:"1234567890"},
	}
	for _,c:=range cases{
		t.Run("Verifying user password",func(t *testing.T){
			pass,err:=store.VerifyUser(ctx,c.email)
			if err!=nil{
				t.Error(err)
			}
			if pass!=c.password{
				t.Errorf("Expected : %s Actual : %s",c.password,pass)
			}

		})
	}
	
	t.Run("Passing an invalid email",func(t *testing.T){
		pass,err:=store.VerifyUser(ctx,"abc@fmail.com")
			if err!=nil{
				t.Error(err)
			}
			if pass!=""{
				t.Error("Expected empty string did not get it")
			}
	})
	
}

func TestInsertUser(t *testing.T) {
	ctx:=context.Background()
	store,err:=GetStore(ctx)
	if err!=nil{
		t.Error(err)
	}
	cases:= []User{
		{Email: "tonystark@gmail.com",
	Password: "Pepper",
},
{Email: "chris@gmail.com",
Password: "Elsa",
},
{Email: "Hulk@gmail.com",
Password: "black widow",
},
{Email: "loki@gmail.com",
Password: "thor",
},
}

for _,c:=range cases{
	t.Run("Testing inserting user",func(t *testing.T){
		err:=store.InserUser(ctx,&c)
		if err!=nil{
			t.Error("Did not expected an error but an error occured")
		}
	})
}
	
}
func TestGetProducts(t *testing.T) {
	ctx:=context.Background()
	store,err:=GetStore(ctx)
	if err!=nil{
		t.Error(err)
	}
	products,err:=store.GetProducts(ctx)
	if err!=nil{
		t.Error(err)
	}
	if len(products)==0{
		t.Error("products should not be empty")
	}
}