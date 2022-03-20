package dao

import (
	"fmt"
	"message-board/Struct"
	"strconv"
	"time"
)

func QueryMovieimfor(id int) *Struct.Movie {

	M := new(Struct.Movie)
	var psid int
	var persons Struct.Actorinmovie
	var time1 []uint8
	err := Db.QueryRow("select id,moviename,yyear,introduction,ddate,posterurl,length,area,type,releasing,feature,score,language from movie where id = ?;", id).Scan(&M.Id, &M.Moviename, &M.Year, &M.Introduction, &time1, &M.Poster, &M.Length, &M.Area, &M.Type, &M.Releasing, &M.Feature, &M.Score, &M.Language)
	if err != nil {
		fmt.Println("查询movie出错", err)
		return nil
	}
	Id := strconv.Itoa(M.Id)
	M.URL = "http://119.91.20.70:6060/object?movie_id=" + Id
	M.Date = utos(time1)
	sqlStr := "select personid from record_direct where pid=?;"
	rows, err := Db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	M.Director = make([]Struct.Actorinmovie, 1)
	for rows.Next() {
		err := rows.Scan(&psid)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		Db.QueryRow("select id,chinesename from person where id=?;", psid).Scan(&persons.Id, &persons.Name)
		Id = strconv.Itoa(persons.Id)
		persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
		M.Director = append(M.Director, persons)
	}
	sqlStr2 := "select personid from record_act where pid=?;"
	rows2, err := Db.Query(sqlStr2, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	M.Actor = make([]Struct.Actorinmovie, 1)
	for rows2.Next() {
		err := rows2.Scan(&psid)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		Db.QueryRow("select id,chinesename from person where id=?;", psid).Scan(&persons.Id, &persons.Name)
		Id = strconv.Itoa(persons.Id)
		persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
		M.Actor = append(M.Actor, persons)
	}
	sqlStr3 := "select personid from record_script where pid=?;"
	rows3, err := Db.Query(sqlStr3, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	M.Scriptwriter = make([]Struct.Actorinmovie, 1)
	for rows3.Next() {
		err := rows3.Scan(&psid)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		Db.QueryRow("select id,chinesename from person where id=?;", psid).Scan(&persons.Id, &persons.Name)
		Id = strconv.Itoa(persons.Id)
		persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
		M.Scriptwriter = append(M.Scriptwriter, persons)
	}
	M.Director = M.Director[1:]
	M.Actor = M.Actor[1:]
	M.Scriptwriter = M.Scriptwriter[1:]
	return M
}

func QueryPersonimfor(id int) *Struct.Person {
	P := new(Struct.Person)
	var mvid int
	var mvs Struct.Movieinactor
	var time1 []uint8
	err := Db.QueryRow("select id,introduction,birthday,Constellations,chinesename,englishname,birthplace,jobs,posterurl from person where id = ?;", id).Scan(&P.Id, &P.Introduction, &time1, &P.Constellations, &P.Chinesename, &P.Englishname, &P.Birthplace, &P.Jobs, &P.Poster)
	if err != nil {
		fmt.Println("查询movie出错", err)
		return nil
	}
	Id := strconv.Itoa(P.Id)
	P.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
	P.Birthday = utos(time1)
	sqlStr := "select pid from record_all where personid=?;"
	rows, err := Db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	P.Works = make([]Struct.Movieinactor, 1)
	for rows.Next() {
		err := rows.Scan(&mvid)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		Db.QueryRow("select id,moviename from movie where id=?;", mvid).Scan(&mvs.Id, &mvs.Name)
		Id = strconv.Itoa(mvs.Id)
		mvs.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
		P.Works = append(P.Works, mvs)
	}
	P.Works = P.Works[1:]
	return P
}

func QueryCooperation(id int) []Struct.Coperson {
	var copersonid int
	sqlStr := "select personid from record_all where personid in (select personid from record_all where pid in(select pid from record_all where personid=?)) and  personid in (select personid from record_all group by personid having count(personid)>2) and personid<>?;"
	rows, err := Db.Query(sqlStr, id, id)
	copersonidslice := make([]int, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	Coperson := make([]Struct.Coperson, 1)
	var Coperson2 Struct.Coperson
	for rows.Next() {
		err := rows.Scan(&copersonid)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		copersonidslice = append(copersonidslice, copersonid)
	}
	copersonidslice = removeDuplicateValues(copersonidslice)
	for i := 0; i < len(copersonidslice); i++ {
		Db.QueryRow("select id,chinesename from person where id =?", copersonidslice[i]).Scan(&Coperson2.Id, &Coperson2.Name)
		Id := strconv.Itoa(Coperson2.Id)
		Coperson2.URL = "http://119.91.20.70:6060/celebrity?person_id=" + Id
		Coperson = append(Coperson, Coperson2)
	}
	Coperson = Coperson[1:]
	rows.Close()
	return Coperson
}

func removeDuplicateValues(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func QueryMoviepic2(id int) []string {
	Moviepiclice := make([]string, 1)
	var moviepicurl string
	sqlStr := "select url from moviepic where pid =?;"
	rows, err := Db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}

	for rows.Next() {
		err := rows.Scan(&moviepicurl)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		Moviepiclice = append(Moviepiclice, moviepicurl)
	}
	Moviepiclice = Moviepiclice[1:]
	rows.Close()

	return Moviepiclice
}

func QueryPersonpic(id int) []string {
	Personpicslice := make([]string, 1)
	var personpicurl string
	sqlStr := "select url from personpic where pid =?;"
	rows, err := Db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}

	for rows.Next() {
		err := rows.Scan(&personpicurl)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		Personpicslice = append(Personpicslice, personpicurl)

	}
	Personpicslice = Personpicslice[1:]
	rows.Close()
	return Personpicslice
}

func Querymovie(id int) bool {
	var moviname string
	err := Db.QueryRow("select moviename from movie where id = ?;", id).Scan(&moviname)
	if err != nil {
		fmt.Println("查询错误", err)
		return false
	}
	return true
}

func QueryReleasingmovie() []Struct.Movie {
	M1 := make([]Struct.Movie, 1)
	var M Struct.Movie
	var psid int
	var persons Struct.Actorinmovie
	var time1 []uint8
	sqlStr0 := "select id,moviename,yyear,introduction,ddate,posterurl,length,area,type,releasing,feature,score,language from movie where releasing = 1;"
	rows0, err := Db.Query(sqlStr0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	for rows0.Next() {
		err := rows0.Scan(&M.Id, &M.Moviename, &M.Year, &M.Introduction, &time1, &M.Poster, &M.Length, &M.Area, &M.Type, &M.Releasing, &M.Feature, &M.Score, &M.Language)
		Id := strconv.Itoa(M.Id)
		M.URL = "http://119.91.20.70:6060/object?movie_id=" + Id
		M.Date = utos(time1)
		sqlStr := "select personid from record_direct where pid=?;"
		rows, err := Db.Query(sqlStr, M.Id)
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M.Director = make([]Struct.Actorinmovie, 1)
		for rows.Next() {
			err := rows.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.QueryRow("select id,chinesename from person where id=?;", psid).Scan(&persons.Id, &persons.Name)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M.Director = append(M.Director, persons)
		}
		sqlStr2 := "select personid from record_act where pid=?;"
		rows2, err := Db.Query(sqlStr2, M.Id)
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M.Actor = make([]Struct.Actorinmovie, 1)
		for rows2.Next() {
			err := rows2.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.QueryRow("select id,chinesename from person where id=?;", psid).Scan(&persons.Id, &persons.Name)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M.Actor = append(M.Actor, persons)
		}
		sqlStr3 := "select personid from record_script where pid=?;"
		rows3, err := Db.Query(sqlStr3, M.Id)
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M.Scriptwriter = make([]Struct.Actorinmovie, 1)
		for rows3.Next() {
			err := rows3.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.QueryRow("select id,chinesename from person where id=?;", psid).Scan(&persons.Id, &persons.Name)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M.Scriptwriter = append(M.Scriptwriter, persons)
		}
		M.Director = M.Director[1:]
		M.Actor = M.Actor[1:]
		M.Scriptwriter = M.Scriptwriter[1:]
		M1 = append(M1, M)
	}
	M1 = M1[1:]
	return M1
}

func QueryHotmovie() []Struct.Movie {
	M1 := make([]Struct.Movie, 1)
	var M Struct.Movie
	var psid int
	var persons Struct.Actorinmovie
	var time1 []uint8
	t1 := time.Now()
	sqlStr0 := "select id,moviename,yyear,introduction,ddate,posterurl,length,area,type,releasing,feature,score,language from movie order by timestampdiff(day,?,ddate )*0.6+score*40 desc ;"
	rows0, err := Db.Query(sqlStr0, t1)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	for rows0.Next() {
		err := rows0.Scan(&M.Id, &M.Moviename, &M.Year, &M.Introduction, &time1, &M.Poster, &M.Length, &M.Area, &M.Type, &M.Releasing, &M.Feature, &M.Score, &M.Language)
		Id := strconv.Itoa(M.Id)
		M.URL = "http://119.91.20.70:6060/object?movie_id=" + Id
		sqlStr := "select personid from record_direct where pid=?;"
		rows, err := Db.Query(sqlStr, M.Id)
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M.Date = utos(time1)
		M.Director = make([]Struct.Actorinmovie, 1)
		for rows.Next() {
			err := rows.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.QueryRow("select id,chinesename from person where id=?;", psid).Scan(&persons.Id, &persons.Name)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M.Director = append(M.Director, persons)
		}
		sqlStr2 := "select personid from record_act where pid=?;"
		rows2, err := Db.Query(sqlStr2, M.Id)
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M.Actor = make([]Struct.Actorinmovie, 1)
		for rows2.Next() {
			err := rows2.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.QueryRow("select id,chinesename from person where id=?;", psid).Scan(&persons.Id, &persons.Name)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M.Actor = append(M.Actor, persons)
		}
		sqlStr3 := "select personid from record_script where pid=?;"
		rows3, err := Db.Query(sqlStr3, M.Id)
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M.Scriptwriter = make([]Struct.Actorinmovie, 1)
		for rows3.Next() {
			err := rows3.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.QueryRow("select id,chinesename from person where id=?;", psid).Scan(&persons.Id, &persons.Name)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M.Scriptwriter = append(M.Scriptwriter, persons)
		}
		M.Director = M.Director[1:]
		M.Actor = M.Actor[1:]
		M.Scriptwriter = M.Scriptwriter[1:]
		M1 = append(M1, M)
	}
	M1 = M1[1:]
	return M1
}

func Querystuff(stuff string) ([]Struct.Movie, []Struct.Person) {
	M := make([]Struct.Movie, 1)
	var M1 Struct.Movie
	P := make([]Struct.Person, 1)
	var P1 Struct.Person
	var time []uint8
	sqlStr := "select id, moviename, yyear, introduction, ddate, posterurl, length, area, type, feature, releasing, score,language from movie where moviename like  CONCAT('%',?,'%') ;" //遍历写给登录用户的评论
	rows, err := Db.Query(sqlStr, stuff)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil, nil
	}

	for rows.Next() {
		err := rows.Scan(&M1.Id, &M1.Moviename, &M1.Year, &M1.Introduction, &time, &M1.Poster, &M1.Length, &M1.Area, &M1.Type, &M1.Feature, &M1.Releasing, &M1.Score, &M1.Language)
		Id := strconv.Itoa(M1.Id)
		M1.URL = "http://119.91.20.70:6060/object?movie_id=" + Id
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil, nil
		}
		M1.Date = utos(time)
		M = append(M, M1)
	}
	rows.Close()
	M = M[1:]
	sqlStr2 := "select id, introduction, birthday, Constellations, chinesename, englishname, birthplace,posterurl from person where chinesename like  CONCAT('%',?,'%') ;" //遍历写给登录用户的评论
	rows2, err := Db.Query(sqlStr2, stuff)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return M, nil
	}

	for rows2.Next() {
		err := rows2.Scan(&P1.Id, &P1.Introduction, &time, &P1.Constellations, &P1.Chinesename, &P1.Englishname, &P1.Birthplace, &P1.Poster)
		Id := strconv.Itoa(P1.Id)
		P1.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return M, nil
		}
		P1.Birthday = utos(time)
		P = append(P, P1)
	}
	rows2.Close()

	P = P[1:]
	return M, P
}

func QueryNewhotmovie() []Struct.Movie {
	M1 := make([]Struct.Movie, 1)
	var M Struct.Movie
	var psid int
	var persons Struct.Actorinmovie
	var time1 []uint8
	t1 := time.Now()
	sqlStr0 := "select id,moviename,yyear,introduction,ddate,posterurl,length,area,type,releasing,feature,score,language from movie where timestampdiff(month ,?,ddate )<3 order by score  desc;"
	rows0, err := Db.Query(sqlStr0, t1)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	for rows0.Next() {
		err := rows0.Scan(&M.Id, &M.Moviename, &M.Year, &M.Introduction, &time1, &M.Poster, &M.Length, &M.Area, &M.Type, &M.Releasing, &M.Feature, &M.Score, &M.Language)
		Id := strconv.Itoa(M.Id)
		M.URL = "http://119.91.20.70:6060/object?movie_id=" + Id
		sqlStr := "select personid from record_direct where pid=?;"
		rows, err := Db.Query(sqlStr, M.Id)
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M.Date = utos(time1)
		M.Director = make([]Struct.Actorinmovie, 1)
		for rows.Next() {
			err := rows.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.QueryRow("select id,chinesename from person where id=?;", psid).Scan(&persons.Id, &persons.Name)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M.Director = append(M.Director, persons)
		}
		sqlStr2 := "select personid from record_act where pid=?;"
		rows2, err := Db.Query(sqlStr2, M.Id)
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M.Actor = make([]Struct.Actorinmovie, 1)
		for rows2.Next() {
			err := rows2.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.QueryRow("select id,chinesename from person where id=?;", psid).Scan(&persons.Id, &persons.Name)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M.Actor = append(M.Actor, persons)
		}
		sqlStr3 := "select personid from record_script where pid=?;"
		rows3, err := Db.Query(sqlStr3, M.Id)
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M.Scriptwriter = make([]Struct.Actorinmovie, 1)
		for rows3.Next() {
			err := rows3.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.QueryRow("select id,chinesename from person where id=?;", psid).Scan(&persons.Id, &persons.Name)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M.Scriptwriter = append(M.Scriptwriter, persons)
		}
		M.Director = M.Director[1:]
		M.Actor = M.Actor[1:]
		M.Scriptwriter = M.Scriptwriter[1:]
		M1 = append(M1, M)
	}
	M1 = M1[1:]
	return M1

}

func Classificationmovie(ty string, area string, year string, feature string) []Struct.Movie {
	M1 := make([]Struct.Movie, 1)
	var M Struct.Movie
	var psid int
	var persons Struct.Actorinmovie
	var time1 []uint8
	var year3 interface{}
	year2, err := strconv.Atoi(year)
	if err != nil {
		fmt.Println(err)
		year3 = "all"
	}
	year3 = year2
	sqlStr0 := "select id,moviename,yyear,introduction,ddate,posterurl,length,area,type,releasing,feature,score,language from movie where type=? or area=? or feature =? or yyear=?;"
	rows0, err := Db.Query(sqlStr0, ty, area, feature, year3)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	for rows0.Next() {
		err := rows0.Scan(&M.Id, &M.Moviename, &M.Year, &M.Introduction, &time1, &M.Poster, &M.Length, &M.Area, &M.Type, &M.Releasing, &M.Feature, &M.Score, &M.Language)
		Id := strconv.Itoa(M.Id)
		M.URL = "http://119.91.20.70:6060/object?movie_id=" + Id
		sqlStr := "select personid from record_direct where pid=?;"
		rows, err := Db.Query(sqlStr, M.Id)
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M.Date = utos(time1)
		M.Director = make([]Struct.Actorinmovie, 1)
		for rows.Next() {
			err := rows.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.QueryRow("select id,chinesename from person where id=?;", psid).Scan(&persons.Id, &persons.Name)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M.Director = append(M.Director, persons)
		}
		sqlStr2 := "select personid from record_act where pid=?;"
		rows2, err := Db.Query(sqlStr2, M.Id)
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M.Actor = make([]Struct.Actorinmovie, 1)
		for rows2.Next() {
			err := rows2.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.QueryRow("select id,chinesename from person where id=?;", psid).Scan(&persons.Id, &persons.Name)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M.Actor = append(M.Actor, persons)
		}
		sqlStr3 := "select personid from record_script where pid=?;"
		rows3, err := Db.Query(sqlStr3, M.Id)
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M.Scriptwriter = make([]Struct.Actorinmovie, 1)
		for rows3.Next() {
			err := rows3.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.QueryRow("select id,chinesename from person where id=?;", psid).Scan(&persons.Id, &persons.Name)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M.Scriptwriter = append(M.Scriptwriter, persons)
		}
		M.Director = M.Director[1:]
		M.Actor = M.Actor[1:]
		M.Scriptwriter = M.Scriptwriter[1:]
		M1 = append(M1, M)
	}
	M1 = M1[1:]
	return M1
}

func ClassificationListmovie(ty string, area string, year string, feature string) []Struct.Movie {
	M1 := make([]Struct.Movie, 1)
	var M Struct.Movie
	var psid int
	var persons Struct.Actorinmovie
	var time1 []uint8
	var year3 interface{}
	year2, err := strconv.Atoi(year)
	if err != nil {
		fmt.Println(err)
		year3 = "all"
	}
	year3 = year2
	sqlStr0 := "select id,moviename,yyear,introduction,ddate,posterurl,length,area,type,releasing,feature,score,language from movie where type=? or area=? or feature =? or yyear=? order by score desc ;"
	rows0, err := Db.Query(sqlStr0, ty, area, feature, year3)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	for rows0.Next() {
		err := rows0.Scan(&M.Id, &M.Moviename, &M.Year, &M.Introduction, &time1, &M.Poster, &M.Length, &M.Area, &M.Type, &M.Releasing, &M.Feature, &M.Score, &M.Language)
		Id := strconv.Itoa(M.Id)
		M.URL = "http://119.91.20.70:6060/object?movie_id=" + Id
		sqlStr := "select personid from record_direct where pid=?;"
		rows, err := Db.Query(sqlStr, M.Id)
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M.Date = utos(time1)
		M.Director = make([]Struct.Actorinmovie, 1)
		for rows.Next() {
			err := rows.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.QueryRow("select id,chinesename from person where id=?;", psid).Scan(&persons.Id, &persons.Name)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M.Director = append(M.Director, persons)
		}
		sqlStr2 := "select personid from record_act where pid=?;"
		rows2, err := Db.Query(sqlStr2, M.Id)
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M.Actor = make([]Struct.Actorinmovie, 1)
		for rows2.Next() {
			err := rows2.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.QueryRow("select id,chinesename from person where id=?;", psid).Scan(&persons.Id, &persons.Name)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M.Actor = append(M.Actor, persons)
		}
		sqlStr3 := "select personid from record_script where pid=?;"
		rows3, err := Db.Query(sqlStr3, M.Id)
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M.Scriptwriter = make([]Struct.Actorinmovie, 1)
		for rows3.Next() {
			err := rows3.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.QueryRow("select id,chinesename from person where id=?;", psid).Scan(&persons.Id, &persons.Name)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M.Scriptwriter = append(M.Scriptwriter, persons)
		}
		M.Director = M.Director[1:]
		M.Actor = M.Actor[1:]
		M.Scriptwriter = M.Scriptwriter[1:]
		M1 = append(M1, M)
	}
	M1 = M1[1:]
	return M1
}
