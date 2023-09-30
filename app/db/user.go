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
	fmt.Println("getting user", email)
	stmt, err := db.Prepare("select email as username, password from user where email = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var user domain.User
	err = stmt.QueryRow(email).Scan(&user.Username, &user.Password)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Close()
	fmt.Println("got user", user.Username)
	return user
}

func RegisterUser(user domain.User) domain.User {
	db := DB()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert into user(email, password) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Password)
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
	fmt.Println("removing user", email)
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

	fmt.Println("removed user", email)
	return true
}
