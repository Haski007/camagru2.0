package user

import (
	"fmt"
	"log"
	"html/template"
	"net/http"
	"strconv"

	"../database"
)

var userID = 0

func handleAddUser(w http.ResponseWriter, r *http.Request) {
	var user User

	user.ID = strconv.Itoa(userID)
	userID++
	fmt.Printf("id = %s\n", strconv.Itoa(userID))
	user.Email = r.FormValue("email")
	user.UserName = r.FormValue("user_name")
	user.Password = r.FormValue("password")
	user.Fname = r.FormValue("first_name")
	user.Lname = r.FormValue("last_name")

	err := database.UsersCollection.Insert(user)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err.Error())
		return
	}
}

// Registration, adding a new one user to the mongoDB.
func Registration(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("user/registration.html")
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err.Error())
		return
	}


	err = t.ExecuteTemplate(w, "registration", nil)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err.Error())
		return
	}
}