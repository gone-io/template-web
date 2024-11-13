package entity

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type LoginParam struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResult struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

type RegisterParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
