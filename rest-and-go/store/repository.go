package store

import (
	"fmt"
	//"log"
	"gopkg.in/mgo.v2"
	// "strings"
	// "gopkg.in/mgo.v2/bson"
)

type Repository struct{}

const SERVER = "mongodb://sean:test1test@ds117145.mlab.com:17145/golanglearn"

const DBNAME = "golanglearn"

const COLLECTION = "store"

func (r Repository) GetProducts() Products{
	session, err := mgo.Dial(SERVER)

	if err != nil {
		fmt.Println("Failed to establish mongo connection", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	results := Products{}

	if err := c.Find(nil).All(&results); err != nil{
		fmt.Println("Failed to write results:", err)
	}

	return results
}