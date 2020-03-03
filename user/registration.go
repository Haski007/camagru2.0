package user

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"../database"
)

var userID = 0

func handleAddUser(w http.ResponseWriter, r *http.Request) {
	var user User

	user.ID = strconv.Itoa(getLastID())
	user.Email = r.FormValue("email")
	user.Login = r.FormValue("login")
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

// Registration draw registration page
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

func getLastID() int {
	var foundUser User

	summ, err := database.UsersCollection.Count()
	if err != nil {
		log.Fatal(err)
	} else if summ == 0 {
		return 0
	}

	err = database.UsersCollection.Find(nil).Skip(summ - 1).One(&foundUser)
	if err != nil {
		log.Println(err)
	}
	res, err := strconv.Atoi(foundUser.ID)
	if err != nil {
		log.Println(err)
	}
	return res + 1
}
