package user

import (
	"net/http"

	"../database"
	"github.com/globalsign/mgo/bson"
	// "github.com/Haski007/camagru2.0/database"
)

// User - a head user's struct!
type User struct {
	ID       string `json:"id" bson:"id"`
	Email    string `json:"email" bson:"email"`
	Login    string `json:"login" bson:"login"`
	Password string `json:"password" bson:"password"`
	Fname    string `json:"first_name" bson:"first_name"`
	Lname    string `json:"last_name" bson:"last_name"`
	Error    error
}

// UsersStore , all users struct
type UsersStore struct {
	users []User
}

// Handler handle GET resp /user/.
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		handleGetUser(w, r)
	} else if r.Method == http.MethodPost {
		handleAddUser(w, r)
	}
}

// GetByLogin return user with your id.
func GetByLogin(login string) (User, error) {
	var user User

	err := database.UsersCollection.Find(bson.M{"login" : login}).One(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}