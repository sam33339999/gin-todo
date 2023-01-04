package users

import (
	"gin-todo/v2/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UsersRegister(router *gin.RouterGroup) {
	router.POST("/", UsersRegistration)
	router.POST("/login", UserLogin)
}

func UsersRegistration(c *gin.Context) {
	userModelValidator := NewUserModelValidator()

	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
}

func UserLogin(c *gin.Context) {

}

func UserRegister(router *gin.RouterGroup) {
	router.GET("/", UserRetrieve)
	router.PUT("/", UserUpdate)
}

func UserRetrieve(c *gin.Context) {

}

func UserUpdate(c *gin.Context) {

}

func ProfileRegister(router *gin.RouterGroup) {
	router.GET("/:username", ProfileRetrieve)
	router.POST("/:username/follow", ProfileFollow)
	router.DELETE("/:username/follow", ProfileUnfollow)
}

func ProfileRetrieve(c *gin.Context) {

}

func ProfileFollow(c *gin.Context) {

}

func ProfileUnfollow(c *gin.Context) {

}
