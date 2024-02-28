package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	db, error := gorm.Open(mysql.Open("root:root@/go_admin"), &gorm.Config{})

	if error != nil {
		panic("Could not connect to the database")
	}

	fmt.Println(db)
}
