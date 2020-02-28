package database

import (
	"log"
	// "fmt"

	"github.com/Haski007/camagru2.0/config"
	"github.com/globalsign/mgo"
)

var usersCollection *mgo.Collection

// InitDB, initialises mongoDB with your configurations at config/config.go.
func InitDB() {
	session, err := mgo.Dial(config.DataBaseHost)
	if err != nil {
		log.Fatal(err)
	}

	usersCollection = session.DB("Camagru").C("users")
	
}