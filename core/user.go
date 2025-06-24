package core

type Login struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	UserId   uint   `gorm:"column:userId;primaryKey;autoIncrement"`
	UserName string `gorm:"column:userName"`
	Account  string `gorm:"column:account"`
	Password string `gorm:"column:password"`
}
