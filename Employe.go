package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//employee struct
type Employee struct {
	ID        int    `json:"id" gorm:"primary_key"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Address   string `json:"address"`
	Contact   int    `json:"contact"`
}

//shows all the employees
func GetAllEmps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "json/application") //sets the data format in json/apps
	var employee []Employee
	DB.Find(&employee) //show all the records form the database
	json.NewEncoder(w).Encode(employee)
}

//shows single employee using id
func GetAllEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "json/application")
	params := mux.Vars(r)
	var employee Employee
	DB.First(&employee, params["id"]) //getting the single record by using id
	json.NewEncoder(w).Encode(employee)
}

//creates employee
func CreateEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "json/application")
	var employee Employee
	json.NewDecoder(r.Body).Decode(&employee) //read data from json body
	DB.Create(&employee)                      //save data into the database
	json.NewEncoder(w).Encode(employee)       //sends data back
}

//update employee
func UpdateEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "json/application")
	params := mux.Vars(r)
	var employee Employee
	DB.First(&employee, params["id"])         //getting the record by id
	json.NewDecoder(r.Body).Decode(&employee) //reading data from the body
	DB.Save(&employee)                        //saving the data
	json.NewEncoder(w).Encode(employee)

}

//delete employee
func DeleteEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "json/application")
	params := mux.Vars(r) //taking id form the url
	var employee Employee
	json.NewDecoder(r.Body).Decode(&employee)
	DB.Delete(&employee, params["id"]) //deleting the employee by using id
	json.NewEncoder(w).Encode("user delete success")
}
