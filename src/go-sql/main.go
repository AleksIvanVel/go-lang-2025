package main

import (
	"fmt"
	"gomyslq/db"
)

func main() {
	db.Connect()
	fmt.Println(db.ExistTable("users"))
	// db.CrateTable(models.UserShema, "users")

	db.Close()
}
