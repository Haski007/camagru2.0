package user

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	structer "../r_structer"
)

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	user, err := GetByLogin(r.FormValue("login"))
	if err != nil {
		log.Println(err)
	}

	fmt.Println("YEaws" )

	t := template.New("t").Funcs(structer.TemplateFuncs)
	t, err = t.ParseFiles("user/get_user.html")
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err.Error())
	}

	err = t.ExecuteTemplate(w, "user", user)
	if err != nil {
		log.Fatal(err)
		fmt.Fprintln(w, err.Error())
	}
}
