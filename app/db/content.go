package db

import (
	"log"

	"github.com/jelinden/content-service/app/domain"
)

func AddContent(spaceID int64, name string, value string) domain.Content {
	db := DB()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert into content(space_id, key, value) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(spaceID, name, value)
	if err != nil {
		log.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}
	stmt.Close()
	content, err := GetContentWithID(spaceID)
	if err != nil {
		log.Println(err)
	}
	return content
}

func GetContentWithSpaceID(spaceID int64) ([]domain.Content, error) {
	db := DB()
	rows, err := db.Query("select id, space_id, key as name, value from content where space_id = ?", spaceID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var contents []domain.Content
	for rows.Next() {
		var content domain.Content
		if err := rows.Scan(&content.ID, &content.SpaceID, &content.Name, &content.Value); err != nil {
			return contents, err
		}
		contents = append(contents, content)
	}
	if err = rows.Err(); err != nil {
		return contents, err
	}
	return contents, nil
}

func GetContentWithID(ID int64) (domain.Content, error) {
	db := DB()
	rows, err := db.Query("select id, space_id, key as name, value from content where id = ?", ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var content domain.Content

	for rows.Next() {
		if err := rows.Scan(&content.ID, &content.SpaceID, &content.Name, &content.Value); err != nil {
			return content, err
		}
	}
	return content, nil
}

func RemoveContent(id int64) bool {
	db := DB()
	stmt, err := db.Prepare("delete from content where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	stmt.Close()
	return true
}
