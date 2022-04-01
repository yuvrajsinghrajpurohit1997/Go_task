package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type users struct {
	Id      int     `gorm:"primaryKey"`
	Fname   string  `json:"fname"`
	City    string  `json:"city"`
	Phone   int     `json:"phone"`
	Height  float32 `json:"height"`
	Married string  `json:"married"`
}

//GetUser function..
func GetUser(w http.ResponseWriter, r *http.Request) {
	connect()
	var user []users
	DB.Find(&user)
	json.NewEncoder(w).Encode(&user)

}

//SearchUser function..
func SearchUser(w http.ResponseWriter, r *http.Request) {
	connect()
	vars := mux.Vars(r)
	userID := vars["id"]
	b := validateId(userID)
	if b == 1 {
		var user users
		DB := DB.Where("id= ?", userID).Find(&user)
		DB.Find(&user)
		json.NewEncoder(w).Encode(&user)
	} else {
		fmt.Fprintln(w, "Please enter a valid id")
	}

}

//SearchUserGroup function..
func SearchUserGroup(w http.ResponseWriter, r *http.Request) {
	connect()
	vars := mux.Vars(r)
	var validId []int
	userID := vars["id"]
	a := strings.Split(userID, ",")
	for i := range a {
		check := validateId(a[i])
		if check == 1 {
			j, _ := strconv.Atoi(a[i])
			validId = append(validId, j)

		}
	}
	var user, arr []users
	for i := range validId {
		DB.Where("id = ?", validId[i]).Find(&user)
		arr = append(arr, user[0])
	}
	json.NewEncoder(w).Encode(arr)
}

//DB Connection funcion..
func connect() {
	dsn := "host=localhost user=postgres password=yuvi@123 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Can not connect to the database")
	}
	DB = db

}

func validateId(id string) int {
	var a int
	// That's obviously not a good validation strategy for email addresses
	// pretend a complex regex here
	if _, err := strconv.Atoi(id); err == nil {
		a = 1
	} else {
		a = 0
	}
	return a
}
