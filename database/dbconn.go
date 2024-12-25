package database

import (
	"fmt"
	"github.com/ArtiomStartev/jwt-auth-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "artiomstartev"
	password = "postgres"
	dbname   = "jwt-auth-api"
)

var DB *gorm.DB
var dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func DBConn() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		return
	}
	DB = db

	if err = db.AutoMigrate(&models.User{}); err != nil {
		fmt.Println("Error migrating database: ", err)
		return
	}
}
