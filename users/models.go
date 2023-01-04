package users

import (
	"errors"
	"gin-todo/v2/common"

	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	ID           uint    `gorm:"primary_key"`
	Username     string  `gorm:"column:username"`
	Email        string  `gorm:"column:email;unique_index"`
	Bio          string  `gorm:"column:bio;size:1024"`
	Image        *string `gorm:"column:image"`
	PasswordHash string  `gorm:"column:password;not null"`
}

func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&UserModel{})

}

func (u *UserModel) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty!")
	}
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}
