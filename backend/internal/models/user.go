package models

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Username  string `gorm:"uniqueIndex:idx_username;not null" json:"username"`
	Email     string `gorm:"uniqueIndex:idx_email;not null" json:"email"`
	Password  string `gorm:"not null" json:"password"`
}
