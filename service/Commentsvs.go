package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/Struct"
	"message-board/dao"
)

func Setcomment(cm Struct.Comment, c *gin.Context) bool {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return false
	}
	flag := dao.Querymovie(cm.Movieid)
	if flag == false {
		c.JSON(404, gin.H{
			"code":   404,
			"reason": "没有该电影",
		})
		return false
	}
	err = dao.Insertcomment(cm)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func Setchildcomment(pid int, from_id int, from_username string, content string, useful int) {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	flag := dao.Insertchcomment(pid, from_id, from_username, content, useful)
	if flag == false {
		fmt.Println("插入评论出错")
	}
}

func ListFilmcomment(movieid int) []Struct.Comment {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cm := dao.QuerycommentwithoutChild(movieid)
	return cm
}

func ListFlimcommentbytime(movieid int) []Struct.Comment {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cm, _ := dao.Queryusermoviecm(movieid)
	return cm
}

func ListFlimcommentbyuse(movieid int) []Struct.Comment {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cm, _ := dao.QueryusermoviecmbyUse(movieid)
	return cm
}

func ListFlimshortcommentbyuselimit(movieid int) []Struct.Shortcomment {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cm := dao.QueryshortcommentbyUsebyLimit(movieid)
	return cm
}

func ListFilmshortcommentbytime(movieid int) []Struct.Shortcomment {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cm := dao.QueryshortcommentbyTime(movieid)
	return cm
}

func ListFilmshortcommentbyuse(movieid int) []Struct.Shortcomment {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cm := dao.QueryshortcommentbyUse(movieid)
	return cm
}

func ListFlimcommentwihtchild(movieid int) []Struct.Comment {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cm := dao.QuerycommentwithoutChild(movieid)
	return cm
}

func Setshortcomment(fromusername string, fromuerid int, content string, lorw int, score float64, movieid int, c *gin.Context) {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	flag := dao.Querymovie(movieid)
	if flag == false {
		c.JSON(404, gin.H{
			"code":   404,
			"reason": "没有该电影",
		})
		return
	}
	dao.Insertshortcomment(fromusername, fromuerid, content, lorw, score, movieid)
}

func Updateshortuse(id int, use int, c *gin.Context) {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	usenum, nouse := dao.Queryshortusenum(id)
	if use == 1 {
		flag := dao.Updateshortuse(id, usenum)
		if flag {
			c.JSON(200, gin.H{
				"code":            200,
				"shortcomment_id": id,
				"if use":          use,
				"msg":             "更新信息成功",
			})
		} else {
			c.JSON(500, gin.H{
				"code": 500,
				"msg":  "更新信息失败",
			})
		}
	} else if use == 0 {
		flag := dao.Updateshortnouse(id, nouse)
		if flag {
			c.JSON(200, gin.H{
				"code":            200,
				"shortcomment_id": id,
				"if use":          use,
				"msg":             "更新信息成功",
			})
		} else {
			c.JSON(500, gin.H{
				"code": 500,
				"msg":  "更新信息失败",
			})
		}
	} else {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "更新信息失败",
		})
	}

}
