package utils

import (
	"errors"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //for use postgres
	_ "github.com/jinzhu/gorm/dialects/sqlite"   //for use sqlite at tests
	"github.com/liftitapp/mongotest/models"
)

var (
	ErrDBConnection = errors.New("failed to connect database")
)

//ConnectToDB func that create a DB connection
func ConnectToDB() (*gorm.DB, error) {
	db, err := gorm.Open(os.Getenv("DATABASE_TYPE"), os.Getenv("DATABASE_URL"))
	return db, err
}

// Clean data base
func CleanDB() {
	db, err := ConnectToDB()
	if err != nil {
		log.Fatal(ErrDBConnection)
	}
	db.Exec("drop table user_examples;")
	defer db.Close()
}

//Migrate database
func MigrateDB() (*gorm.DB, error) {
	db, err := ConnectToDB()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&models.UserExample{})

	var user models.UserExample
	db.First(&user, 1)
	if user == (models.UserExample{}) {
		db.Create(&models.UserExample{
			Name:       "Liftit Admin",
			Email:      "optiLiftiAdmin@liftit.co",
			Role:       "admin",
			Password:   "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918",
			WebhookURL: "http://liftit.co",
			Token:      "d033e22ae348aeb5660fc2140aec35850c4da997",
		})
	}
	return db, err
}

// CreateUser create a new user in the database
func CreateUser(db *gorm.DB, user models.UserExample) (models.UserExample, error) {
	var err error
	if db == nil {
		db, err = ConnectToDB()
		if err != nil {
			log.Println("---- error in CreateUser", err.Error())
			return user, ErrCode7001
		}
		defer db.Close()
	}

	result := db.Create(&user)
	if result.Error != nil {
		log.Println("---- error in CreateUser", result.Error.Error())
		return user, ErrCode7002
	}
	return user, nil
}

// CreateTracker returns all people from DB
func CreateTracker() []models.TrackerRegister {
	var newTracker []models.TrackerRegister
	return newTracker
}
