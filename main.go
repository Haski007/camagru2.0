package main

import (
	"./config"
	"./logger"

	"fmt"
	"log"
	"time"
	"net/http"

)

func main() {
	s := http.Server{
		Addr:           config.HostName + config.Port,
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/", logger.JSON(startHandler))

	fmt.Println("Listerning on port", config.Port)
	log.Fatal(s.ListenAndServe())
}

func startHandler(w http.ResponseWriter, r *http.Request) {

}