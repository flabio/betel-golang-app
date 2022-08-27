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
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
	// 	os.Getenv("MYSQL_USER"),
	// 	os.Getenv("MYSQL_PASSWORD"),
	// 	os.Getenv("MYSQL_HOST"),
	// 	os.Getenv("MYSQL_PORT"),

	// 	os.Getenv("MYSQL_DATABASE"),
	// )
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	// db.AutoMigrate(
	// 	&User{},
	// 	&Rol{},
	// 	&Role{},
	// 	&Church{},
	// 	&Detachment{},
	// 	&SubDetachment{},
	// 	&Patrol{},
	// 	&StudyCarried{},
	// 	&Module{},
	// 	&RoleModule{},
	// 	&RoleChurch{},
	// 	//&ParentScout{},
	// 	&MinisterialAcademy{},
	// 	&UserSubdetachement{},
	// 	&Attendance{},
	// 	&WeeksDetachment{},
	// 	&City{},
	// 	&Parent{},

	// 	&Visit{},
	// )

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

// Closedb
func Closedb() {
	var db *gorm.DB = DatabaseConnection()
	CloseConnection(db)

}
