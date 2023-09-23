package domain

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Username struct {
	Username string `json:"username"`
}

type JwtToken struct {
	Token string `json:"token"`
}
