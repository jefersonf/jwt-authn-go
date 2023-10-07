package env

import (
	"fmt"

	"github.com/jefersonf/jwt-authn-go/models"
)

func SyncDatabase(dbHandler *ConnectionHandler) {
	err := dbHandler.DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println(err)
	}
}
