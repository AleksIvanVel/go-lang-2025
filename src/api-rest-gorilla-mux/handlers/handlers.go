package handlers

import (
	"apirest-gorillamux/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	if users, err := models.ListUsers(); err != nil {
		models.SendNoFound(rw)
	} else {
		models.SendData(rw, users)
	}
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	if user, err := getUserByRequest(r); err != nil {
		models.SendNoFound(rw)
	} else {
		models.SendData(rw, user)
	}
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {

	// Obtiene el registro
	user := models.User{}

	// Recupera el body de la peticion
	decoder := json.NewDecoder(r.Body)

	// Convierte el json al objeto User
	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		user.Save()
		models.SendData(rw, user)
	}
}

func UpdateUSer(rw http.ResponseWriter, r *http.Request) {

	var userId int64

	if user, err := getUserByRequest(r); err != nil {
		models.SendNoFound(rw)
	} else {
		userId = user.Id
	}

	user := models.User{}
	// Recupera el body de la peticion
	decoder := json.NewDecoder(r.Body)

	// Convierte el json al objeto User
	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		user.Id = userId
		user.Save()
		models.SendData(rw, user)
	}
}
func DeleteUSer(rw http.ResponseWriter, r *http.Request) {
	if user, err := getUserByRequest(r); err != nil {
		models.SendNoFound(rw)
	} else {
		user.Delete()
		models.SendData(rw, user)
	}
}

func getUserByRequest(r *http.Request) (models.User, error) {
	//Obtenemos el ID
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	if user, err := models.GetUser(userId); err != nil {
		return *user, err
	} else {
		return *user, nil
	}
}
