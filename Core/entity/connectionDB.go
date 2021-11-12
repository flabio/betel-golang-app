package entity

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// setup database connection is creating a new connection to our database
func DatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}
	db.AutoMigrate(
		&User{},
		&Rol{},
		&Role{},
		&Church{},
		&Detachment{},
		&SubDetachment{},
		&Patrol{},
		&StudyCarried{},
		&Module{},
		&RoleModule{},
		&RoleChurch{},
		&ParentScout{},
	)

	return db

}

// Close database connection method is closin a connection between your app db
func CloseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()

}

//Closedb
func Closedb() {
	var db *gorm.DB = DatabaseConnection()
	CloseConnection(db)

}
