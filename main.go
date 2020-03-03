package main

import (
	"html/template"

	"./config"
	"./database"
	"./logger"
	"./user"

	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	s := http.Server{
		Addr:          	config.Port,
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	database.InitDB()
	http.HandleFunc("/", logger.PreLogs(startHandler))
	http.HandleFunc("/user", logger.PreLogs(user.Handler))
	http.HandleFunc("/registration", logger.PreLogs(user.Registration))



	fmt.Println("Listerning on port", config.Port)
	log.Fatal(s.ListenAndServe())
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err.Error())
		return
	}


	err = t.ExecuteTemplate(w, "main-page", nil)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err.Error())
		return
	}
}