package db

import (
	"fmt"
	"testing"

	"github.com/jelinden/content-service/app/domain"
	"github.com/jelinden/content-service/app/util"
	"github.com/stretchr/testify/assert"
)

func init() {
	dbFileName = "./content-service-test.db"
	Init()
}

func TestRun(t *testing.T) {
	testRegister(t)
	defer postTests()
}

func testRegister(t *testing.T) {
	fmt.Println(dbFileName)

	const email = "test@email.localhost"
	password, err := util.HashPassword("hashedpassword")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	user := RegisterUser(domain.User{Username: email, Password: password, HashedPassword: password})
	assert.True(t, user.Username == email)
	assert.True(t, user.Password == password)
	defer RemoveUser(email)
}

func postTests() {
	removeDB(dbFileName)
}
