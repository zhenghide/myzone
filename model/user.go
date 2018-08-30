package model

type User struct {
	Id         string `json:"id"`
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
	NickName   string `json:"nick_name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}

