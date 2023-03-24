package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
)

// GetAllUser...
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	query := "SELECT * FROM users"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	var user User
	var users []User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Email, &user.Password, &user.UserType); err != nil {
			log.Println(err)
			return
		} else {
			users = append(users, user)
		}
	}

	var response UsersResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = users

	w.Header().Set("Content-Type", "application/")
	err2 := json.NewEncoder(w).Encode(response)
	if err2 != nil {
		log.Println(err2)
		fmt.Println(err2)
	}
}

// Insert User...
func InsertUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	// Read from Request Body
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	name := r.Form.Get("name")
	age, _ := strconv.Atoi(r.Form.Get("age"))
	address := r.Form.Get("address")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	_, errQuery := db.Exec("INSERT INTO users(name, age, address, email, password) values (?,?,?,?,?)",
		name,
		age,
		address,
		email,
		password,
	)

	var response ErrorResponse
	if age <= 0 {
		if errQuery != nil {
			fmt.Println(errQuery)
			response.Status = 400
			response.Message = "Insert Failed!"
		}
	} else {
		response.Status = 200
		response.Message = "Success"
	}

	w.Header().Set("Content-Type", "application/json")
	err2 := json.NewEncoder(w).Encode(response)
	if err2 != nil {
		log.Println(err2)
		fmt.Println(err2)
	}
}

// Update User...
func UpdateUser(params martini.Params, w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	// Read from Request Body
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}

	userId := params["idUser"]

	name := r.Form.Get("name")
	age, _ := strconv.Atoi(r.Form.Get("age"))
	address := r.Form.Get("address")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	sqlStatement := `
		UPDATE users 
		SET name = ?, age = ?, address =  ?, email = ?, password = ?
		WHERE id = ?`

	_, errQuery := db.Exec(sqlStatement,
		name,
		age,
		address,
		email,
		password,
		userId,
	)

	var response ErrorResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "Update Failed!"
	}

	w.Header().Set("Content-Type", "application/json")
	err2 := json.NewEncoder(w).Encode(response)
	if err2 != nil {
		log.Println(err2)
		fmt.Println(err2)
	}
}

// Delete User
func DeleteUser(params martini.Params, w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	userId := params["idUser"]

	err := r.ParseForm()
	if err != nil {
		return
	}

	_, errQuery := db.Exec("DELETE FROM users WHERE id=?",
		userId,
	)

	var response ErrorResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "Delete Failed!"
	}

	w.Header().Set("Content-Type", "application/json")
	err2 := json.NewEncoder(w).Encode(response)
	if err2 != nil {
		log.Println(err2)
		fmt.Println(err2)
	}
}

func sendErrorResponse(w http.ResponseWriter, message string) {
	var response ErrorResponse
	response.Status = 400
	response.Message = message

	w.Header().Set("Content-Type", "application/json")
	err2 := json.NewEncoder(w).Encode(response)
	if err2 != nil {
		log.Println(err2)
		fmt.Println(err2)
	}
}
