package main

import (
	"encoding/json"
	"fmt"
)

//  bson is a named type that declares a map of key/value pairs.
type bson map[string]interface{}

// user is a struct type that declares user information.
type user struct {
	Id   int
	Name string
}

func main() {
	// Retrieve the user profile.
	users, err := RetrieveUser()
	if err != nil {
		fmt.Println(err)
	}

	// Display each user from the slice.
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}
}

// RetrieveUser retrieves the user document for the specified
// user and returns a pointer to a user type value.
func RetrieveUser() ([]user, error) {
	// Make the web call to get the json response.
	r, err := GetUsers()
	if err != nil {
		return nil, err
	}

	// Unmarshal the json document into a value of
	// the user struct type.
	var users []user
	err = json.Unmarshal([]byte(r), &users)
	return users, err
}

// GetUsers simulates a web call that returns an array
// of json documents with users.
func GetUsers() (string, error) {
	response := `[{"id":1432, "name":"sally"},{"id":4312, "name":"bill"},{"id":6721, "name":"jane"}]`
	return response, nil
}
