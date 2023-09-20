package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {
	PORT_NUM := 3000
	// The below will take the index.html file
	fileServer := http.FileServer(http.Dir("./public"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello" , helloHandler)

	fmt.Printf("Starting server at port number %v",PORT_NUM)
	if err := http.ListenAndServe(":3000" , nil ); err != nil {
		log.Fatal(err)
	}

}

// request is something send by the user to the server and response is something that is send from the server to the user(client)
func helloHandler(w http.ResponseWriter , r *http.Request){
 if r.URL.Path != "/hello"{
	http.Error(w , "404 not found" , http.StatusNotFound)
	return
 }

 if r.Method != "GET" {
	http.Error(w , "Other method than GET is not supported" , http.StatusNotFound)
	return
 }

 fmt.Fprint(w,"Hello World")

}

func formHandler(w http.ResponseWriter , r *http.Request){
	if r.Method != "POST"{
		http.Error(w , "It is not a GET method" , http.StatusNotFound)
		return
	}

	if err := r.ParseForm() ; err != nil{
		fmt.Fprintf(w,"ParseForm error %v" , err)
		return
	}

	// Where with the request we can access the formValue with the corresponding element name
	name := r.FormValue("name")
	number := r.FormValue("number")

	fmt.Fprintf(w , "The name is %v and the number is %v" , name , number)
}
