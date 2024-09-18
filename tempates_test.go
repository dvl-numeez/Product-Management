package main

import (
	"bytes"
	"context"
	"testing"
)


func TestSignInTemplate(t *testing.T) {
signInTemplate:=`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body> 

<h1>This is a sign in page</h1>
<form action="/signIn" method="POST">
    <label for="username">Email:</label><br>
    <input type="text" id="email" name="email" required><br><br>
    
    <label for="password">Password:</label><br>
    <input type="password" id="password" name="password" required><br><br>
    
    <input id="signIn"type="submit" value="Submit">
</form>
<script>
</script>

<footer>This is the footer</footer>
</body>
`
output:=bytes.Buffer{}
tmpl,err:=makeTemplate("templates/sign-in.html")
	if err!=nil{
		t.Error(err)
	}
	err=tmpl.Execute(&output,nil)
	if err!=nil{
	 t.Error(err)
	}
	if output.String()!=signInTemplate{
		t.Errorf("Expected : %s \n Got : %s",signInTemplate,output.String())
	}
	t.Run("Giving a wrong path",func(t *testing.T){
		_,err:=makeTemplate("abc")
		if err==nil{
			t.Error("Expected an error but did not get it")
		}
	})
}
func TestSignUpTempate(t *testing.T) {
signUpTemplate:=`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body> 

<h1>This is a sign up page</h1>
<form id="signUpForm"  method="POST" action="/signUp">
    <label for="email">Email:</label><br>
    <input type="text" id="email" name="email" required><br><br>
    
    <label for="password">Password:</label><br>
    <input type="password" id="password" name="password" required><br><br>
    
    <input id="signup"type="submit" value="Submit">
</form>

<footer>This is the footer</footer>
</body>
`
output:=bytes.Buffer{}
tmpl,err:=makeTemplate("templates/sign-up.html")
	if err!=nil{
		t.Error(err)
	}
	err=tmpl.Execute(&output,nil)
	if err!=nil{
	 t.Error(err)
	}
	if output.String()!=signUpTemplate{
		t.Errorf("Expected : %s \n Got : %s",signUpTemplate,output.String())
	}

}

func TestHomeTemplate(t *testing.T) {
homeTemplate:=`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body> 

    <h1>This is the home page</h1>
    <h3>Welcome Guest</h3>
   <table>
       <tr>
           <th>Product Name</th>
           <th>Product Type</th>
           <th>Product Material</th>
       </tr>
       
       <tr>
           <td>Eco-Friendly Yoga Mat</td>
           <td>Fitness Equipment</td>
           <td>Natural Rubber</td>
       </tr>
       
       <tr>
           <td>Stainless Steel Water Bottle</td>
           <td>Hydration</td>
           <td> Stainless Steel</td>
       </tr>
       
       <tr>
           <td>Leather Office Chair</td>
           <td>Furniture</td>
           <td>Leather</td>
       </tr>
       
   </table>
        
        <li>Eco-Friendly Yoga Mat</li>
        
        <li>Stainless Steel Water Bottle</li>
        
        <li>Leather Office Chair</li>
        
    </ul>

<footer>This is the footer</footer>
</body>
`
ctx:=context.Background()
	store,err:=GetStore(ctx)
	if err!=nil{
		t.Error(err)
	}
products,err:=store.GetProducts(ctx)
	if err!=nil{
		t.Error()
	}
	info := HomePageInfo{
		Username: "Guest",
		Products: products,
	}
output:=bytes.Buffer{}
tmpl,err:=makeTemplate("templates/home.html")
	if err!=nil{
		t.Error(err)
	}
	err=tmpl.Execute(&output,info)
	if err!=nil{
	 t.Error(err)
	}
	if output.String()!=homeTemplate{
		t.Errorf("Expected : %s \n Got : %s",homeTemplate,output.String())
	}
t.Run("Having a greeting message with name",func(t *testing.T){
homeTemplate:=`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body> 

    <h1>This is the home page</h1>
    <h3>Welcome Numeez</h3>
   <table>
       <tr>
           <th>Product Name</th>
           <th>Product Type</th>
           <th>Product Material</th>
       </tr>
       
       <tr>
           <td>Eco-Friendly Yoga Mat</td>
           <td>Fitness Equipment</td>
           <td>Natural Rubber</td>
       </tr>
       
       <tr>
           <td>Stainless Steel Water Bottle</td>
           <td>Hydration</td>
           <td> Stainless Steel</td>
       </tr>
       
       <tr>
           <td>Leather Office Chair</td>
           <td>Furniture</td>
           <td>Leather</td>
       </tr>
       
   </table>
        
        <li>Eco-Friendly Yoga Mat</li>
        
        <li>Stainless Steel Water Bottle</li>
        
        <li>Leather Office Chair</li>
        
    </ul>

<footer>This is the footer</footer>
</body>
`
	products,err:=store.GetProducts(ctx)
	if err!=nil{
		t.Error()
	}
	info := HomePageInfo{
		Username: "Numeez",
		Products: products,
	}
output:=bytes.Buffer{}
tmpl,err:=makeTemplate("templates/home.html")
	if err!=nil{
		t.Error(err)
	}
	err=tmpl.Execute(&output,info)
	if err!=nil{
	 t.Error(err)
	}
	if output.String()!=homeTemplate{
		t.Errorf("Expected : %s \n Got : %s",homeTemplate,output.String())
	}

})

}