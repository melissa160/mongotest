package main

import (
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mrkaspa/amqputils"
	"github.com/streadway/amqp"

	"github.com/liftitapp/mongotest/constants"
	"github.com/liftitapp/mongotest/services"
	"github.com/liftitapp/mongotest/utils"
)

func main() {

	// Enviroment variables validation
	err := utils.ValidateEnvVars()
	if err != nil {
		log.Fatal(err)
	}
	// Charge error code array
	utils.SetupErrorCode()

	// DB connection
	db, err := utils.MigrateDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// RabbitMQ conexion and start services
	startRabbitAndServices(db)
}

func startRabbitAndServices(db *gorm.DB) {
	conn, ch, close, err := amqputils.CreateConnection(os.Getenv("RABBIT_PATH"))
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	runExample, err := amqputils.NewServer(ch, constants.ExampleQueue, services.ExampleSaveUser)
	if err != nil {
		panic(err)
	}

	go runExample.Start()
	utils.Msg(" [*] Waiting for logs. To exit press CTRL+C", "")
	testRabbitConnection(conn)
}

func testRabbitConnection(conn *amqp.Connection) {
	for {
		ch, err := conn.Channel()
		if err != nil {
			ch.Close()
			panic(err)
		}
		ch.Close()
		time.Sleep(30 * time.Second)
	}
}
