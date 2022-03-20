package service

import (
	"fmt"
	"message-board/Struct"
	"message-board/dao"
)

func Movieinfor(id int) *Struct.Movie {
	err := dao.OpenDb()

	if err != nil {
		fmt.Println(err)
	}
	dao.Scoredao(id)
	M := dao.QueryMovieimfor(id)
	return M
}

func Personinfor(id int) *Struct.Person {
	err := dao.OpenDb()

	if err != nil {
		fmt.Println(err)
	}

	M := dao.QueryPersonimfor(id)

	return M
}

func Checkmoviealiveser(id int) bool {
	dao.OpenDb()
	flag := dao.Querymovie(id)
	if flag == false {
		fmt.Println("未找到")
		return flag
	}
	return flag
}

func Moviepicsvs(id int) []string {
	dao.OpenDb()
	P := dao.QueryMoviepic2(id)

	return P
}

func Personpicsvs(id int) []string {
	dao.OpenDb()
	P := dao.QueryPersonpic(id)

	return P
}

func Copersonsvs(id int) []Struct.Coperson {
	dao.OpenDb()
	C := dao.QueryCooperation(id)

	return C
}

func HotMovieinfor() []Struct.Movie {
	err := dao.OpenDb()

	if err != nil {
		fmt.Println(err)
	}
	M := dao.QueryHotmovie()
	return M
}

func RealeasingMovieinfor() []Struct.Movie {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
	}
	M := dao.QueryReleasingmovie()
	return M
}

func RecommendMovieinfor() []Struct.Movie {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
	}
	M := dao.QueryNewhotmovie()
	return M
}

func ClassHotMovieinfor(ty string, area string, yaer string, feature string) []Struct.Movie {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
	}
	M := dao.ClassificationListmovie(ty, area, yaer, feature)
	return M
}

func ClassMovieinfor(ty string, area string, yaer string, feature string) []Struct.Movie {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
	}
	M := dao.Classificationmovie(ty, area, yaer, feature)
	return M
}

func Qerymovie(stuff string) ([]Struct.Movie, []Struct.Person) {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
	}
	M, P := dao.Querystuff(stuff)
	return M, P
}
