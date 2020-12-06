package model

type User struct {
	BaseModel
	Username string
	Name     string
	Email    string
	Password string
	IsActive bool
}
