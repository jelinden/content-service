package db

import (
	"log"

	"github.com/jelinden/content-service/app/domain"
)

func AddSpace(user domain.User, name string) []domain.Space {
	db := DB()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert into space(user_id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID, name)
	if err != nil {
		log.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}
	stmt.Close()
	spaces, err := GetSpacesWithUserID(user.ID)
	if err != nil {
		log.Println(err)
	}
	return spaces
}

func GetSpacesWithUserID(userID int64) ([]domain.Space, error) {
	db := DB()
	rows, err := db.Query("select id, name from space where user_id = ?", userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var spaces []domain.Space
	for rows.Next() {
		var space domain.Space
		if err := rows.Scan(&space.ID, &space.Name); err != nil {
			return spaces, err
		}
		spaces = append(spaces, space)
	}
	if err = rows.Err(); err != nil {
		return spaces, err
	}
	return spaces, nil
}

func GetSpaceContentWithUserID(spaceID, userID int64) ([]domain.Content, error) {
	db := DB()
	rows, err := db.Query(`
		select c.key, c.value
		from space s, content c
		where s.user_id = ?
		and s.id = ?
		and c.space_id = s.id`, userID, spaceID,
	)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var content []domain.Content
	for rows.Next() {
		var c domain.Content
		if err := rows.Scan(&c.Name, &c.Value); err != nil {
			return content, err
		}
		content = append(content, c)
	}
	if err = rows.Err(); err != nil {
		return content, err
	}
	return content, nil
}

func GetSpace(id int64) domain.Space {
	db := DB()
	stmt, err := db.Prepare("select id, name from space where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var space domain.Space
	err = stmt.QueryRow(id).Scan(&space.ID, &space.Name)
	if err != nil {
		log.Println(err)
	}
	stmt.Close()
	return space
}

func RemoveSpace(id int64) bool {
	db := DB()
	stmt, err := db.Prepare("delete from space where id = ?")
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
