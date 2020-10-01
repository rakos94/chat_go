package models

// User models
type User struct {
	ID       int64 `gorm:"primaryKey"`
	Name     string
	Password string
	Email    string
}
