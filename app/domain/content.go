package domain

type Data struct {
	Content []Content `json:"content"`
}
type Content struct {
	ID      int64  `json:"id,omitempty"`
	SpaceID int64  `json:"spaceID,omitempty"`
	Name    string `json:"name"`
	Value   string `json:"value"`
}
