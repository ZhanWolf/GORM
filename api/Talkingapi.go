package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/dao"
	"message-board/service"
	"strconv"
)

func Talkingapi(c *gin.Context) {
	username := Getusernamefromtoken(c)
	id, err := dao.Queryusername(username)
	if err != nil {
		fmt.Println(err)
		return
	}
	content := c.PostForm("content")
	title := c.PostForm("title")
	movie_id := c.PostForm("movie_id")
	movie_id2, _ := strconv.Atoi(movie_id)
	flag := service.Talkingsvs(username, id, content, title, movie_id2, c)
	if flag {
		c.JSON(200, gin.H{
			"code":     200,
			"msg":      "插入讨论成功",
			"username": username,
			"userid":   id,
			"movie_id": movie_id,
			"content":  content,
		})
	}
}

func Talkingcmapi(c *gin.Context) {
	username := Getusernamefromtoken(c)
	id, err := dao.Queryusername(username)
	if err != nil {
		fmt.Println(err)
		return
	}
	talkingid := c.PostForm("talking_id")
	talkingid2, _ := strconv.Atoi(talkingid)
	content := c.PostForm("content")
	flag := service.Cmintalkingsvs(talkingid2, content, username, id, c)
	if flag {
		c.JSON(200, gin.H{
			"code":       200,
			"msg":        "插入评论成功",
			"username":   username,
			"userid":     id,
			"talking_id": talkingid2,
			"content":    content,
		})
	}
}

func Talkingchcmapi(c *gin.Context) {
	username := Getusernamefromtoken(c)
	id, err := dao.Queryusername(username)
	if err != nil {
		fmt.Println(err)
		return
	}
	pid := c.PostForm("pid")
	pid2, _ := strconv.Atoi(pid)
	content := c.PostForm("content")
	flag := service.ChCmintalkingsvs(pid2, content, username, id, c)
	if flag {
		c.JSON(200, gin.H{
			"code":     200,
			"msg":      "插入子评论成功",
			"username": username,
			"userid":   id,
			"pid":      pid,
			"content":  content,
		})
	}
}

func Listtalkinginmovieapi(c *gin.Context) {
	movieid := c.Query("movie_id")
	movieid2, _ := strconv.Atoi(movieid)
	talkings := service.Listalksvs(movieid2)
	c.JSON(200, talkings)
}

func Listonetaking(c *gin.Context) {
	takingid := c.Query("talking_id")
	movieid2, _ := strconv.Atoi(takingid)
	talkings := service.Lisonetalksvs(movieid2)
	c.JSON(200, talkings)

}
