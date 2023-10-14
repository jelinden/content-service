package domain

type Space struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"userID" db:"user_id"`
	Name   string `json:"spaceName" db:"name"`
}
