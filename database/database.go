package database

import (
	"log"

	"github.com/Haski007/camagru2.0/config"
	"github.com/globalsign/mgo"
)

// UsersCollection main users table.
var UsersCollection *mgo.Collection

// InitDB, initialises mongoDB with your configurations at config/config.go.
func InitDB() {
	session, err := mgo.Dial(config.DataBaseHost)
	if err != nil {
		log.Fatal(err)
	}

	UsersCollection = session.DB("Camagru").C("users")

	if err = session.Ping(); err != nil {
		log.Fatal(err)
	}
}