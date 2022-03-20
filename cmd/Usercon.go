package cmd

import (
	"github.com/gin-gonic/gin"
	"message-board/api"
	"message-board/jwt"
)

func Userroute(r *gin.Engine) {
	us := r.Group("/user")
	{
		us.POST("/login", api.Login)
		us.GET("/oauth", api.Oauth)
		us.POST("/loginbygithub", api.Loginbygithub)
		us.POST("/Singup", api.Singup)
		us.POST("/Reset", api.Reset)
		us.POST("/QueryProtectionQ", api.QueryprotectionQ)
		us.GET("/clock", jwt.JWTAuth(), api.Clock)
		us.GET("/imfor", jwt.JWTAuth(), api.Userimfor)
		us.POST("/change", jwt.JWTAuth(), api.Setuserintroduction)
		us.POST("/otherimfor", api.OtherUserimfor)
	}
}
