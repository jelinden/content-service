package db

import (
	"fmt"
	"log"
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
	space := AddSpace(domain.User{ID: 1, Username: email, Password: password, HashedPassword: password, ApiToken: apiToken}, spaceName)
	assert.True(t, space[0].Name == spaceName)
	defer RemoveSpace(space[0].ID)

	AddContent(space[0].ID, "key", "value")
	content, err := GetSpaceContentWithUserID(space[0].ID, 1)
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	log.Println(content)
	assert.True(t, len(content.Content) > 0)
	assert.Equal(t, content.Content[0].Name, "key")
	assert.Equal(t, content.Content[0].Value, "value")
}

func postTestSpace() {
	removeDB(dbFileName)
}
