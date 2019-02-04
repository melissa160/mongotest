package utils

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/liftitapp/mongotest/models"
)

func SetupDatabaseTest() *gorm.DB {
	db, err := connectToDBTest()
	if err != nil {
		log.Fatal(err)
	}

	SetTestEnviroment()
	//if db.HasTable("user_examples") == false {
	db.AutoMigrate(&models.UserExample{})
	//}
	db.Exec("drop table user_examples;")
	db.AutoMigrate(&models.UserExample{})

	return db
}

func SetTestEnviroment() {
	err := ValidateTestEnvVars()
	if err != nil {
		log.Fatal(err)
	}
	os.Setenv("DATABASE_URL", os.Getenv("DATABASE_URL_TEST"))
}

func CleanDBTest() {
	db, err := connectToDBTest()
	if err != nil {
		log.Fatal(ErrDBConnection)
	}
	db.Exec("drop table user_examples;")
	defer db.Close()
}

//ConnectToDB func that create a DB connection
func connectToDBTest() (*gorm.DB, error) {
	db, err := gorm.Open(os.Getenv("DATABASE_TYPE"), os.Getenv("DATABASE_URL_TEST"))
	return db, err
}

func MigrateDBTest() {
	db, err := connectToDBTest()

	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.UserExample{})
	defer db.Close()
}
