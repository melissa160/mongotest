package utils

import (
	"testing"

	"github.com/liftitapp/mongotest/models"
	"github.com/stretchr/testify/assert"
)

func TestCrudUser(t *testing.T) {
	db := SetupDatabaseTest()
	defer db.Close()
	assert.NotEqual(t, db, nil)

	// Create user
	var user0 = models.UserExample{
		Name:     "Jonathan",
		Email:    "dayessisanchez@gmail.com",
		Password: EncryptPassword([]byte("saraylizzy"))}

	_, err := CreateUser(db, user0)
	assert.Equal(t, err, nil, "user created")
}

func TestInvalidUser(t *testing.T) {
	db := SetupDatabaseTest()
	defer db.Close()
	assert.NotEqual(t, db, nil)

	// Create user
	var user1 = models.UserExample{
		Name:     "Jonathan1",
		Email:    "dayessisanchez1@gmail.com",
		Password: EncryptPassword([]byte("saraylizzy"))}

	var _, err = CreateUser(db, user1)
	assert.Equal(t, err, nil, "user created2")

	_, err = CreateUser(db, user1)
	assert.NotEqual(t, err, nil, "Error duplicate user mail")
}
