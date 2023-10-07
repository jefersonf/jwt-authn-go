package env

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConnectionHandler struct {
	DB *gorm.DB
	On bool
}

// Conn return a new connection to database
func Conn() func() *ConnectionHandler {
	connHandler := new(ConnectionHandler)
	dsn, err := buildDSN()
	return func() *ConnectionHandler {
		if connHandler.On {
			log.Println("database connection is already established")
			return connHandler
		}
		connHandler.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("fail to connect to database")
		}
		connHandler.On = true
		log.Println("successful connected to database")
		return connHandler
	}
}
