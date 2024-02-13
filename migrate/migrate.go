package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Database connection closed")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
