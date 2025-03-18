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
	user.Delete()

	fmt.Println(models.ListUsers())

	db.Close()
}
