package main

import (
	"p1-lc03.com/app"
	l "p1-lc03.com/utils/logger"

	"p1-lc03.com/repo"
)


func main(){
	l.NewLogger("games.log")

	host := "127.0.0.1"
	port := 3306
	username := "ayaka"
	password := "ayakaPa$$word2781"
	dbName := "db_games"

	err := repo.ConnectDB(host, username, password, dbName, port)

	if err != nil{
		panic(err)
	}

	app.Run()

}