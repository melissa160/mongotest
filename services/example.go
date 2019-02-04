package services

import (
	"encoding/json"
	"log"
	"os"

	"github.com/mrkaspa/amqputils"
	"github.com/streadway/amqp"

	"github.com/liftitapp/mongotest/constants"
	"github.com/liftitapp/mongotest/models"
	"github.com/liftitapp/mongotest/utils"
)

// Example receive rabbitmq messages and return status charge and charge validated
var ExampleSaveUser amqputils.SubscribeFunc = func(message amqp.Delivery) ([]byte, error) {
	var originalExample models.UserExample
	json.Unmarshal(message.Body, &originalExample)

	response := writteUserExample(originalExample)

	body, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	return body, nil
}

func writteUserExample(user models.UserExample) models.Response {

	db, _ := utils.ConnectToDB()
	transaction := db.Begin()

	userCreated, errUser := utils.CreateUser(transaction, user)

	if errUser != nil {
		transaction.Rollback()
		var err models.ErrorCode = utils.GetError(errUser.Error())

		return models.Response{
			Code:    err.Code,
			Message: err.Message,
			Data:    userCreated,
		}
	}

	result := transaction.Commit()
	if result.Error != nil {
		transaction.Rollback()
		var err models.ErrorCode = utils.GetError(utils.Error7003)
		return models.Response{
			Code:    err.Code,
			Message: err.Message,
			Data:    userCreated,
		}
	}

	errPublish := publishUserCreated(userCreated)
	if errPublish != nil {
		var err models.ErrorCode = utils.GetError(utils.Error7004)
		return models.Response{
			Code:    err.Code,
			Message: err.Message,
			Data:    userCreated,
		}
	}

	return models.Response{
		Code:    1,
		Message: "ok",
		Data:    userCreated,
	}

}

func publishUserCreated(user models.UserExample) error {
	body, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	amqputils.Publish(os.Getenv("RABBIT_PATH"), constants.PublishUserCreated, body)
	return err
}
