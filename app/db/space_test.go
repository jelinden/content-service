package db

import (
	"fmt"
	"testing"

	"github.com/jelinden/content-service/app/domain"
	"github.com/jelinden/content-service/app/util"
	"github.com/stretchr/testify/assert"
)

func beforeTest() {
	dbFileName = "./content-service-test.db"
	Init()
}

func TestRunSpace(t *testing.T) {
	beforeTest()
	testAddSpace(t)
	defer postTestSpace()
}

func testAddSpace(t *testing.T) {
	fmt.Println(dbFileName)

	const email = "test@email.localhost"
	password, err := util.HashPassword("hashedpassword")
	const spaceName = "myspace"
	if err != nil {
		assert.Fail(t, err.Error())
	}
	apiToken := util.GenerateToken(email)
	space := AddSpace(domain.User{Username: email, Password: password, HashedPassword: password, ApiToken: apiToken}, spaceName)
	assert.True(t, space[0].Name == spaceName)
	defer RemoveSpace(space[0].ID)
}

func postTestSpace() {
	removeDB(dbFileName)
}
