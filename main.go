package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB
var err error

func initDB() {
	DB, err = gorm.Open("sqlite3", "emp.db")
	if err != nil {
		panic(err.Error())
	}
	DB.AutoMigrate(&Employee{}) //creating tables from struct
}

func initRouter() {
	//mux router initialization
	r := mux.NewRouter()
	r.HandleFunc("/employee", GetAllEmps).Methods("GET")
	r.HandleFunc("/employee/{id}", GetAllEmp).Methods("GET")
	r.HandleFunc("/employee", CreateEmp).Methods("POST")
	r.HandleFunc("/employee/{id}", UpdateEmp).Methods("PUT")
	r.HandleFunc("/employee/{id}", DeleteEmp).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func main() {
	initDB()     //database setup method
	initRouter() //routes and handle function method
}
