package domain

type User struct {
	ID             int64  `json:"id"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	HashedPassword string `json:"-"`
	ApiToken       string `json:"apiToken"`
}

type Username struct {
	Username string `json:"username"`
}

type JwtToken struct {
	Token string `json:"token"`
}
