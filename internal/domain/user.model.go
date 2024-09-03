package domain

import "gorm.io/gorm"

//DEFINICION DE USUARIO PROVENIENTE DE BDD

type User struct {
	gorm.Model
	Username       *string   `gorm:"unique;not null" json:"username"`
	Password       *string   `json:"password"`
	Email          *string   `gorm:"unique;not null" json:"email"`
	FirstName      *string   `json:"firstName"`
	LastName       *string   `json:"lastName"`
	Documents      *[]string `gorm:"type:text[]" json:"documents"`
	Role           *string   `json:"role"`
	LastConnection *string   `json:"lastConnection"`
	Token          *string   `json:"token"`
}

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}
