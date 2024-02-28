package initializers

import (
	"example/go-crud/models"
)


func MigrateDB() {

	if DB == nil {
		panic("Database connection is nil. Check the initialization.")
	}

	DB.AutoMigrate(&models.Post{})
	DB.AutoMigrate(&models.User{})

}