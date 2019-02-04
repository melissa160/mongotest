package services

import (
	"testing"

	"github.com/liftitapp/mongotest/models"
	"github.com/liftitapp/mongotest/utils"

	"github.com/stretchr/testify/assert"
)

func init() {
	utils.SetupErrorCode()
	utils.SetTestEnviroment()
}

func TestWritteUserExample(t *testing.T) {
	utils.CleanDBTest()
	utils.MigrateDBTest()
	var user1 = models.UserExample{
		Name:     "user3",
		Email:    "Email3@mail.com",
		Password: utils.EncryptPassword([]byte("hola")),
		Role:     "user"}

	resp := writteUserExample(user1)
	assert.Equal(t, resp.Message, "ok")
}

func TestWritteUserExampleFail(t *testing.T) {
	utils.CleanDBTest()
	utils.MigrateDBTest()
	var user1 = models.UserExample{
		Name:     "user1",
		Email:    "Email1@mail.com",
		Password: utils.EncryptPassword([]byte("hola")),
		Role:     "user"}

	resp := writteUserExample(user1)
	resp = writteUserExample(user1)
	assert.Equal(t, resp.Message, utils.Error7002)
}
