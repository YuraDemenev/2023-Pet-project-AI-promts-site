package site

type User struct {
	Id       int    `json:"-" db:"id"`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
