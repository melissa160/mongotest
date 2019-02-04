package utils

import (
	"errors"
	"log"

	"github.com/liftitapp/mongotest/models"
)

var Errors []models.ErrorCode

// This constants are used for describe errors
const Error7000 = "Unexpected error"
const Error7001 = "Failed Database connection"
const Error7002 = "The record couldn't be created"
const Error7003 = "Failed in commit transaction"
const Error7004 = "Failed in publishProcessCharge"

// This var are used to create an error base on the constants described above
var (
	ErrCode7001 = errors.New(Error7001)
	ErrCode7002 = errors.New(Error7002)
	ErrCode7003 = errors.New(Error7003)
	ErrCode7004 = errors.New(Error7004)
)

// Slice with app errors used
func SetupErrorCode() {
	Errors = append(Errors, models.ErrorCode{Code: 7001, Message: Error7001})
	Errors = append(Errors, models.ErrorCode{Code: 7002, Message: Error7002})
	Errors = append(Errors, models.ErrorCode{Code: 7003, Message: Error7003})
	Errors = append(Errors, models.ErrorCode{Code: 7004, Message: Error7004})
	Errors = append(Errors, models.ErrorCode{Code: 7000, Message: Error7000})
}

// Return ErrorCode according to code
func GetError(code string) models.ErrorCode {
	for _, err := range Errors {
		if err.Message == code {
			return err
		}
	}
	log.Println("Unexpected error with code", code)
	return GetError(Error7000)
}
