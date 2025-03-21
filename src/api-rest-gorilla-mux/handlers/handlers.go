package handlers

import (
	"apirest-gorillamux/db"
	"apirest-gorillamux/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	db.Connect()
	users, _ := models.ListUsers()
	db.Close()

	//Transformamos el objeto en tipo json
	output, _ := json.Marshal(users)
	fmt.Fprintln(rw, string(output))

}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	// Obtiene el Id
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	db.Connect()
	user, _ := models.GetUser(userId)
	db.Close()

	//Transformamos el objeto en tipo json
	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	// Obtiene el registro
	user := models.User{}

	// Recupera el body de la peticion
	decoder := json.NewDecoder(r.Body)

	// Convierte el json al objeto User
	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	} else {
		db.Connect()
		user.Save()
		db.Close()
	}

	//Transformamos el objeto en tipo json
	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

func UpdateUSer(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	// Obtiene el registro
	user := models.User{}

	// Recupera el body de la peticion
	decoder := json.NewDecoder(r.Body)

	// Convierte el json al objeto User
	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	} else {
		db.Connect()
		user.Save()
		db.Close()
	}

	//Transformamos el objeto en tipo json
	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}
func DeleteUSer(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	// Obtiene el Id
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	db.Connect()
	user, _ := models.GetUser(userId)
	user.Delete()
	db.Close()

	//Transformamos el objeto en tipo json
	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}
