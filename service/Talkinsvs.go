package service

import (
	"github.com/gin-gonic/gin"
	"message-board/Struct"
	"message-board/dao"
)

func Talkingsvs(username string, userid int, contetnt string, title string, movieid int, c *gin.Context) bool {
	dao.OpenDb()
	flag := dao.Inserttalking(username, userid, contetnt, title, movieid)
	if flag == false {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "Insert Wrong",
		})
		return false
	}
	return true
}

func Cmintalkingsvs(talkingid int, content string, username string, userid int, c *gin.Context) bool {
	dao.OpenDb()
	err := dao.Inserttkcomment(talkingid, content, username, userid)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "Insert Wrong",
		})
		return false
	}
	return true
}

func ChCmintalkingsvs(pid int, content string, username string, userid int, c *gin.Context) bool {
	dao.OpenDb()
	err := dao.Inserttkchcomment(pid, content, username, userid)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "Insert Wrong",
		})
		return false
	}
	return true
}

func Listalksvs(movieid int) []Struct.Tkinmovie {
	dao.OpenDb()
	Talkings := dao.Querytalkinginmovie(movieid)
	return Talkings
}

func Lisonetalksvs(Talkingid int) Struct.Talking {
	dao.OpenDb()
	Talkings := dao.Querytalking(Talkingid)
	return Talkings
}
