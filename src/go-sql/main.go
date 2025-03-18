package main

import (
	"fmt"
	"gomyslq/db"
	"gomyslq/models"
)

func main() {
	db.Connect()

	// user := models.CreateUser("ivan", "ivan456", "ivan@example.com")
	// fmt.Println(user)
	// fmt.Println(db.ExistTable("users"))
	// db.CrateTable(models.UserShema, "users")
	// db.TruncateTable("users")

	user := models.GetUser(2)
	fmt.Println(user)
	user.Username = "juan"
	user.Password = "juan789"
	user.Email = "juan@example.com"
	user.Save()

	fmt.Println(models.ListUsers())

	db.Close()
}
