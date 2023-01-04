package users

import (
	"gin-todo/v2/common"

	"github.com/gin-gonic/gin"
)

type UserModelValidator struct {
	User struct {
		Username string `form:"username" json:"username" binding:"exists,alphanum,required,min=4,max=255"`
		Email    string `form:"email" json:"email" binding:"exists,email"`
		Password string `form:"password" json:"password" binding:"exists,required,min=8,max=255"`
		Bio      string `form:"bio" json:"bio" binding:"max=1024"`
		Image    string `form:image" json:"image" binding:"omitempty,url"`
	} `json:"user"`
	userModel UserModel `json:"-"`
}

func NewUserModelValidator() UserModelValidator {
	userModelValidator := UserModelValidator{}
	return userModelValidator
}

func (validator *UserModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, validator)

	if err != nil {
		return err
	}
	validator.userModel.Username = validator.User.Username
	validator.userModel.Email = validator.User.Email
	validator.userModel.Bio = validator.User.Bio

	if validator.User.Password != common.NBRandomPassword {
		validator.userModel.setPassword(validator.User.Password)
	}
	if validator.User.Image != "" {
		validator.userModel.Image = &validator.User.Image
	}
	return nil
}
