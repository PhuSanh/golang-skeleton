package entity

type UserID int64

type User struct {
	UserID   uint64
	Username string
	Name     string
	Email    string
	Password string
	IsActive bool
}
