package domain

type Content struct {
	ID      int64  `json:"id"`
	SpaceID int64  `json:"spaceID"`
	Name    string `json:"name"`
	Value   string `json:"value"`
}
