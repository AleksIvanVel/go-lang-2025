package models

type User struct {
	Id       int
	Username string
	Password string
	Email    string
}

const UserShema string = `CREATE TABLE users(
	id INT(6) UNSIGNEd AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`
