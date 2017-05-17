package main

import (
	"log"
	"net/http"
)

//http://localhost:8080/users GET
//http://localhost:8080/users POST
//JSON: 
/*{
    "firstName": "Prasad2",
    "lastName": "Eduballi2",
    "emailAddress": "prsddasad2@gmail.com",
    "phoneNo": 1233131
  }*/
//http://localhost:8080/users/Prasad GET
//http://localhost:8080/users/Sandesh DELETE
//http://localhost:8080/usersexport GET

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}