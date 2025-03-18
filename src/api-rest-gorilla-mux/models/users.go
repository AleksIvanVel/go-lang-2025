package models

import "apirest-gorillamux/db"

type User struct {
	Id       int64
	Username string
	Password string
	Email    string
}

type Users []User

const UserShema string = `CREATE TABLE users(
	id INT(6) UNSIGNEd AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`

// Construir usuario
func NewUser(username, password, email string) *User {
	user := &User{
		Username: username,
		Password: password,
		Email:    email,
	}

	return user
}

// Crear e insertat usuario en db
func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save()

	return user
}

// Inserta un registro en la db
func (user *User) insert() {
	sql := "INSERT users SET username=?, password=?, email=?"
	result, _ := db.Exec(sql, user.Username, user.Password, user.Email)

	user.Id, _ = result.LastInsertId()
}

// Listar registros
func ListUsers() Users {
	//Crea la consulta
	sql := "SELECT id, username, password, email FROM users"

	//Incializamos una lista que contendra los usuarios
	users := Users{}

	//guarda el resultado de la consulta
	rows, _ := db.Query(sql)

	// recorre el resultado de la consulta
	for rows.Next() {
		//COnvierte el resultado en un objeto de tipo Usuer
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)

		// Guarda el objeto de tipo User en la lista Users
		users = append(users, user)
	}
	return users
}

// Obtener un registro
func GetUser(id int) *User {
	user := NewUser("", "", "")

	sql := "SELECT id, username, password, email FROM users WHERE id=?"

	rows, _ := db.Query(sql, id)

	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	}
	return user
}

// Actualizar un registro
func (user *User) update() {
	sql := "UPDATE users SET username=?, password=?, email=? WHERE id=?"

	db.Exec(sql, user.Username, user.Password, user.Email, user.Id)
}

// guardar o editar registro
func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}
}

// Eliminar un registro
func (user *User) Delete() {
	sql := "DELETE FROM users WHERE id=?"

	db.Exec(sql, user.Id)
}
