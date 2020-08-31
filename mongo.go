package main

import (
	"gopkg.in/mgo.v2"
)

func main() {

	// открываем соединение
	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		panic(err)
	}
	println("Connection...")
	defer session.Close()


}