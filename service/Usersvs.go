package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/Struct"
	"message-board/dao"
	"net/http"
)

func UserLoginser(username string, password string) string {
	dao.OpenDb()
	turepassword := dao.Queryuserpassword(username)
	if turepassword != password {
		return ""
	}

	return username
}

func Checkuseraliveser(username string) error {
	dao.OpenDb()
	_, err := dao.Queryusername(username)
	return err
}

func UserSingup(username string, password string, passwordagain string, protectionQ string, protectionA string) bool {
	dao.OpenDb()

	if passwordagain != password {
		return true
	}

	err := dao.Insertuser(username, password, protectionQ, protectionA)
	if err != nil {
		fmt.Println("注册错误", err)
	}

	return false
}

func PasswordReset(c *gin.Context, username string) (bool, string) {
	dao.OpenDb()
	protectionQ, trueprotectionA := dao.Queryprotection(username)
	id, err := dao.Queryusername(username)
	if err != nil {
		fmt.Println(err)
		return false, "none"
	}
	if protectionQ == "" {
		c.JSON(403, gin.H{
			"code":     403,
			"id":       id,
			"username": username,
			"reason":   "该用户未设置密保",
		})
		return false, "none"
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":        200,
			"username":    username,
			"protectionQ": protectionQ,
		})

	}
	return true, trueprotectionA
}

func PasswordReset2(c *gin.Context, username string, password string, protectionA string, passwordagain string) {
	_, trueprotectionA := dao.Queryprotection(username)
	id, _ := dao.Queryusername(username)
	if trueprotectionA == protectionA && password == passwordagain {
		dao.Updatepassword(password, username)
		c.JSON(http.StatusOK, gin.H{
			"code":        200,
			"id":          id,
			"username":    username,
			"performance": "修改密码成功",
		})
	} else if passwordagain != password && trueprotectionA == protectionA {
		c.JSON(403, gin.H{
			"code":     403,
			"id":       id,
			"username": username,
			"reason":   "两次密码不相同",
		})
	} else if trueprotectionA != protectionA {
		c.JSON(403, gin.H{
			"code":     403,
			"id":       id,
			"username": username,
			"reason":   "密保答案有误",
		})
	}
}

func Listuserimfor(username string, c *gin.Context) (Struct.Userimfor, bool) {
	dao.OpenDb()
	var U Struct.Userimfor
	var err error
	U.Id, err = dao.Queryusername(username)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "无该用户",
		})
		return U, false
	}
	U.Username = username
	U.Introduction, err = dao.Queryintroducton(U.Id)
	U.Cm = dao.QueryUsercm(U.Id)
	U.Scm = dao.QueryUserscm(U.Id)
	U.Looked = dao.Looked(U.Id)
	U.Wanted = dao.Wanted(U.Id)
	return U, true
}

func Setintroduction(username string, introduction string) error {
	dao.OpenDb()
	Id, err := dao.Queryusername(username)
	if err != nil {
		fmt.Println(err)
		return err
	}
	dao.UpdateIntroduction(introduction, Id)
	return nil
}

func UserLoginser2(username string, password string) bool {
	dao.OpenDb()
	turepassword := dao.Queryuserpassword(username)
	if turepassword != password {
		return false
	}

	return true
}
