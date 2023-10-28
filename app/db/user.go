package db

import (
	"fmt"
	"log"

	"github.com/jelinden/content-service/app/domain"
)

var dbFileName string

func Init() {
	fmt.Println(dbFileName)
	if dbFileName == "" {
		dbFileName = "./content-service.db"
	}
	fmt.Println(dbFileName)
	OpenDB(dbFileName)
}

func GetUser(email string) domain.User {
	db := DB()
	log.Println("getting user", email)
	stmt, err := db.Prepare("select id, email as username, password, apitoken from user where email = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var user domain.User
	err = stmt.QueryRow(email).Scan(&user.ID, &user.Username, &user.Password, &user.ApiToken)
	if err != nil {
		log.Println(err)
	}
	stmt.Close()
	log.Println("got user", user.Username)
	return user
}

func RegisterUser(user domain.User) domain.User {
	db := DB()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert into user(email, password, apitoken) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.HashedPassword, user.ApiToken)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	stmt.Close()
	return GetUser(user.Username)
}

func RemoveUser(email string) bool {
	db := DB()
	log.Println("removing user", email)
	stmt, err := db.Prepare("delete from user where email = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(email)
	if err != nil {
		log.Fatal(err)
	}

	stmt.Close()

	log.Println("removed user", email)
	return true
}

func GetUserWithToken(token string) domain.User {
	db := DB()
	log.Println("getting user with token", token)
	stmt, err := db.Prepare("select id, email as username, password, apitoken from user where apitoken = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var user domain.User
	err = stmt.QueryRow(token).Scan(&user.ID, &user.Username, &user.Password, &user.ApiToken)
	if err != nil {
		log.Println(err)
	}
	stmt.Close()
	log.Println("got user", user.Username, "with token", token)
	return user
}
