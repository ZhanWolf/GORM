package api

import (
	"github.com/gin-gonic/gin"
	"message-board/service"
	"net/http"
	"strconv"
)

func Movieimforapi(c *gin.Context) {
	movieid := c.Query("movie_id")

	movieid2, _ := strconv.Atoi(movieid)
	flag := service.Checkmoviealiveser(movieid2)
	if flag == false {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "无该电影",
		})
		c.Abort()
		return
	}
	M := service.Movieinfor(movieid2)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
	c.JSON(http.StatusOK, M)

}

func Personapi(c *gin.Context) {
	personid := c.Query("person_id")

	personid2, _ := strconv.Atoi(personid)
	P := service.Personinfor(personid2)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
	c.JSON(http.StatusOK, P)
}

func Moviepicapi(c *gin.Context) {
	movieid := c.Query("movie_id")

	movieid2, _ := strconv.Atoi(movieid)
	M := service.Moviepicsvs(movieid2)
	c.JSON(http.StatusOK, M)
}

func Personpic(c *gin.Context) {
	movieid := c.Query("person_id")

	movieid2, _ := strconv.Atoi(movieid)
	M := service.Personpicsvs(movieid2)
	c.JSON(http.StatusOK, M)
}

func Coperson(c *gin.Context) {
	movieid := c.Query("person_id")

	movieid2, _ := strconv.Atoi(movieid)
	M := service.Copersonsvs(movieid2)
	c.JSON(http.StatusOK, M)
}

func HotMovieimforapi(c *gin.Context) {
	M := service.HotMovieinfor()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
	c.JSON(http.StatusOK, M)
}

func RealeasingMovieimforapi(c *gin.Context) {
	M2 := service.RealeasingMovieinfor()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
	c.JSON(http.StatusOK, M2)
}

func Newhotlist(c *gin.Context) {
	M2 := service.RecommendMovieinfor()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
	c.JSON(http.StatusOK, M2)
}

func Searchmovie(c *gin.Context) {
	stuff := c.PostForm("stuff")
	M2, P2 := service.Qerymovie(stuff)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
	c.JSON(http.StatusOK, M2)
	c.JSON(http.StatusOK, P2)
}

func Classhotlist(c *gin.Context) {
	year := c.PostForm("year")
	ty := c.PostForm("type")
	area := c.PostForm("area")
	feature := c.PostForm("feature")
	M2 := service.ClassHotMovieinfor(ty, area, year, feature)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
	c.JSON(http.StatusOK, M2)
}

func Classmovie(c *gin.Context) {
	year := c.PostForm("year")
	ty := c.PostForm("type")
	area := c.PostForm("area")
	feature := c.PostForm("feature")
	M2 := service.ClassMovieinfor(ty, area, year, feature)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
	c.JSON(http.StatusOK, M2)
}
