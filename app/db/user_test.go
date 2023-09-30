package db

import (
	"fmt"
	"testing"

	"github.com/jelinden/content-service/app/domain"
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
	const password = "hashedpassword"

	user := RegisterUser(domain.User{Username: email, Password: password})
	assert.True(t, user.Username == email)
	assert.True(t, user.Password == password)
	defer RemoveUser(email)
}

func postTests() {
	removeDB(dbFileName)
}
