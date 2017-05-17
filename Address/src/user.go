package main

type User struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
	PhoneNo      int    `json:"phoneNo"`
}

type Users []User
