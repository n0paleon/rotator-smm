package migrations

import (
	"RotatorSMM/pkg/models"
)

func Users() {
	DB.AutoMigrate(&models.User{})
}
