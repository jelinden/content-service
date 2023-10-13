package domain

type Space struct {
	ID     int64  `json:"id"`
	UserID int64  `db:"user_id"`
	Name   string `db:"name"`
}
