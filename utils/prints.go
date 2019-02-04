package utils

import (
	"log"
)

// Msg func for print message
func Msg(title, a interface{}) {
	log.Printf("mongotest - %v: %#v ", title, a)
}
