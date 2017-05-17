package main

import (
	"fmt"
	"strings"
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

var currentId int

var users Users

// Give us some seed data
func init() {
	RepoCreateUser(User{FirstName: "Prasad", LastName: "Eduballi", EmailAddress: "prasad@gmail.com", PhoneNo: 1233131})
	RepoCreateUser(User{FirstName: "Prasad1", LastName: "Eduballi", EmailAddress: "prasad1@gmail.com", PhoneNo: 1233131})
}

func RepoFindUser(name string) User {
	for _, t := range users {
		if strings.Contains(t.FirstName, name) {
			return t
		}
	}
	// return empty User if not found
	return User{}
}

//this is bad, I don't think it passes race condtions
func RepoCreateUser(t User) User {
	users = append(users, t)
	return t
}

func RepoDestroyUser(name string) error {
	for i, t := range users {
		if strings.Contains(t.FirstName, name) {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find User with id of %d to delete", name)
}

func Export() {
	file, err := os.Create("users.csv")
	checkError("Cannot create file", err)
	defer file.Close()
	
	w := csv.NewWriter(file)

	for _, record := range users {
		var writingRecord []string
		writingRecord = append(writingRecord,record.FirstName)
		writingRecord = append(writingRecord,record.LastName)
		writingRecord = append(writingRecord,record.EmailAddress)
		writingRecord = append(writingRecord,strconv.Itoa(record.PhoneNo))
		if err := w.Write(writingRecord); err != nil {
			log.Fatalln("error writing record to csv:", err)
			fmt.Println(err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	defer w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	
	fmt.Println("done with")
}


func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}