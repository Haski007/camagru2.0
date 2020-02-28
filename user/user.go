package user

import (
	"fmt"
	"net/http"

	// "github.com/Haski007/camagru2.0/database"
)

// User, a head user's struct!
type User struct {
	ID       string `json:"id" bson:"id"`
	Email    string `json:"email" bson:"email"`
	UserName string `json:"user_name" bson:"user_name"`
	Password string `json:"password" bson:"password"`
	Fname    string `json:"first_name" bson:"first_name"`
	Lname    string `json:"last_name" bson:"last_name"`
	Error    error
}

// UsersStore, all users struct
type UsersStore struct {
	users []User
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		handleGetUser(w, r)
	} else if r.Method == http.MethodPost {
		handleAddUser(w, r)
	}
}

func GetBookByID(id string) (*User, error) {
	
	var err error = fmt.Errorf("User with id: [%s], not found!", id)
	return nil, err
}


