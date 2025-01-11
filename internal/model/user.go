package model

type User struct {
	Name     string `json:"name" gorm:"not null;type:text"`
	Email    string `json:"email" gorm:"primaryKey;type:text"`
	Password string `json:"-" gorm:"not null;type:text"`
}
