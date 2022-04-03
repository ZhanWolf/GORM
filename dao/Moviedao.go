package dao

import (
	"fmt"
	"message-board/Struct"
	"strconv"
)

func QueryMovieimfor(id int) *Struct.MovieinWeb {

	M := new(Struct.MovieinWeb)
	var M1 Struct.Movie
	var psid int
	var persons Struct.Actorinmovie
	var time1 []uint8
	//err := Db.QueryRow("select id,moviename,yyear,introduction,ddate,posterurl,length,area,type,releasing,feature,score,language from movie where id = ?;", id).Scan(&M.Id, &M.Moviename, &M.Year, &M.Introduction, &time1, &M.Poster, &M.Length, &M.Area, &M.Type, &M.Releasing, &M.Feature, &M.Score, &M.Language)
	res := Db.First(&Struct.Movie{}).Where("id=?", id).Scan(&M1)
	if res.Error != nil {
		fmt.Println(res.Error)
		return nil
	}
	M = &Struct.MovieinWeb{
		Id:           M1.Id,
		Introduction: M1.Introduction,
		Poster:       M1.Poster,
		Year:         M1.Year,
		Date:         M1.Date,
		Moviename:    M1.Moviename,
		Score:        M1.Score,
		Language:     M1.Language,
		Length:       M1.Length,
		Area:         M1.Area,
		Type:         M1.Type,
		Feature:      M1.Feature,
		URL:          M1.URL,
		Releasing:    M1.Releasing,
	}
	Id := strconv.Itoa(M1.Id)
	M.URL = "http://119.91.20.70:6060/object?movie_id=" + Id
	M.Date = utos(time1)
	//sqlStr := "select personid from record_direct where pid=?;"
	rows, err := Db.Model(&Struct.Record_direct{}).Where("movie_id=?", id).Select("person_id").Rows()
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
		//Db.QueryRow("select id,chinesename from person where id=?;", psid).Scan(&persons.Id, &persons.Name)
		res = Db.First(&Struct.Person{}).Where("id=?", psid).Scan(&persons)
		Id = strconv.Itoa(persons.Id)
		persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
		M.Director = append(M.Director, persons)
	}
	//sqlStr2 := "select personid from record_act where pid=?;"
	rows2, err := Db.Model(&Struct.Record_act{}).Where("movie_id=?", id).Select("person_id").Rows()
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
		//Db.QueryRow("select id,chinesename from person where id=?;", psid).Scan(&persons.Id, &persons.Name)
		res = Db.First(&Struct.Person{}).Where("id=?", psid).Scan(&persons)
		Id = strconv.Itoa(persons.Id)
		persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
		M.Actor = append(M.Actor, persons)
	}
	//sqlStr3 := "select personid from record_script where pid=?;"
	rows3, err := Db.Model(&Struct.Record_script{}).Where("movie_id=?", id).Select("person_id").Rows()
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
		res = Db.First(&Struct.Person{}).Where("id=?", psid).Scan(&persons)
		Id = strconv.Itoa(persons.Id)
		persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
		M.Scriptwriter = append(M.Scriptwriter, persons)
	}
	M.Director = M.Director[1:]
	M.Actor = M.Actor[1:]
	M.Scriptwriter = M.Scriptwriter[1:]
	return M
}

func QueryPersonimfor(id int) *Struct.PersoninWeb {
	P := new(Struct.PersoninWeb)
	var mvid int
	var mvs Struct.Movieinactor
	var time1 []uint8
	//err := Db.QueryRow("select id,introduction,birthday,Constellations,chinesename,englishname,birthplace,jobs,posterurl from person where id = ?;", id).Scan(&P.Id, &P.Introduction, &time1, &P.Constellations, &P.Chinesename, &P.Englishname, &P.Birthplace, &P.Jobs, &P.Poster)
	res := Db.First(&Struct.Person{}).Where("id=?", id).Scan(&P)
	if res.Error != nil {
		fmt.Println("查询movie出错", res.Error)
		return nil
	}
	Id := strconv.Itoa(P.Id)
	P.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
	P.Birthday = utos(time1)
	//sqlStr := "select pid from record_all where personid=?;"
	rows, err := Db.First(&Struct.Record_all{}).Where("person_id=?", id).Rows()
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
		//Db.QueryRow("select id,moviename from movie where id=?;", mvid).Scan(&mvs.Id, &mvs.Name)
		res = Db.First(&Struct.Movie{}).Where("id=?", id).Scan(&mvs)
		Id = strconv.Itoa(mvs.Id)
		mvs.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
		P.Works = append(P.Works, mvs)
	}
	P.Works = P.Works[1:]
	return P
}

func QueryCooperation(id int) []Struct.Coperson {
	var copersonid int
	//sqlStr := "select personid from record_all where personid in (select personid from record_all where pid in(select pid from record_all where personid=?)) and  personid in (select personid from record_all group by personid having count(personid)>2) and personid<>?;"
	rows, err := Db.Raw("select person_id from record_all where person_id in (select person_id from record_all where pid in(select pid from record_all where person_id=?)) and  person_id in (select person_id from record_all group by person_id having count(person_id)>2) and person_id<>?;").Rows()
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
		//Db.QueryRow("select id,chinesename from person where id =?", copersonidslice[i]).Scan(&Coperson2.Id, &Coperson2.Name)
		Db.First(&Struct.Coperson{}).Where("id=?", copersonidslice[i]).Scan(&Coperson2)
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
	//sqlStr := "select url from moviepic where pid =?;"
	rows, err := Db.Model(&Struct.Moviepic{}).Where("pid=?", id).Select("url").Rows()
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
	//sqlStr := "select url from personpic where pid =?;"
	rows, err := Db.Raw("select url from personpic where pid =?;").Rows()
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
	//err := Db.QueryRow("select moviename from movie where id = ?;", id).Scan(&moviname)
	err := Db.First(&Struct.Movie{}).Select("moviename").Where("id=?", id).Scan(&moviname)
	if err != nil {
		fmt.Println("查询错误", err)
		return false
	}
	return true
}

func QueryReleasingmovie() []Struct.MovieinWeb {
	M1 := make([]Struct.MovieinWeb, 1)
	var M2 Struct.Movie
	var M Struct.MovieinWeb
	var psid int
	var persons Struct.Actorinmovie
	var time1 []uint8
	//sqlStr0 := "select id,moviename,yyear,introduction,ddate,posterurl,length,area,type,releasing,feature,score,language from movie where releasing = 1;"
	rows0, err := Db.Model(&Struct.Movie{}).Where("releasing=?", 1).Rows()
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	for rows0.Next() {
		err := rows0.Scan(&M2)
		Id := strconv.Itoa(M2.Id)
		M.URL = "http://119.91.20.70:6060/object?movie_id=" + Id
		M.Date = utos(time1)
		M = Struct.MovieinWeb{
			Id:           M2.Id,
			Introduction: M2.Introduction,
			Poster:       M2.Poster,
			Year:         M2.Year,
			Date:         M2.Date,
			Moviename:    M2.Moviename,
			Score:        M2.Score,
			Language:     M2.Language,
			Length:       M2.Length,
			Area:         M2.Area,
			Type:         M2.Type,
			Feature:      M2.Feature,
			URL:          M2.URL,
			Releasing:    M2.Releasing,
		}
		//sqlStr := "select personid from record_direct where pid=?;"
		rows, err := Db.First(&Struct.Record_direct{}).Where("pid=?", M.Id).Rows()
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
			//Db.QueryRow("select id,chinesename from person where id=?;", psid).Scan(&persons.Id, &persons.Name)
			Db.First(&Struct.Person{}).Where("id=?", psid).Scan(&persons)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M.Director = append(M.Director, persons)
		}
		//sqlStr2 := "select personid from record_act where pid=?;"
		//rows2, err := Db.Query(sqlStr2, M.Id)
		rows2, err := Db.First(&Struct.Record_act{}).Where("pid=?", M.Id).Rows()
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
			Db.First(&Struct.Person{}).Where("id=?", psid).Scan(&persons)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M.Actor = append(M.Actor, persons)
		}
		//sqlStr3 := "select personid from record_script where pid=?;"
		rows3, err := Db.First(&Struct.Record_script{}).Where("pid=?", M.Id).Rows()
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
			Db.First(&Struct.Person{}).Where("id=?", psid).Scan(&persons)
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

func QueryHotmovie() []Struct.MovieinWeb {
	M1 := make([]Struct.MovieinWeb, 1)
	var M Struct.Movie
	var M2 Struct.MovieinWeb
	var psid int
	var persons Struct.Actorinmovie
	var time1 []uint8
	rows0, err := Db.Raw("select id,moviename,yyear,introduction,ddate,posterurl,length,area,type,releasing,feature,score,language from movie order by timestampdiff(day,?,ddate )*0.6+score*40 desc ;").Rows()
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	for rows0.Next() {
		err := rows0.Scan(&M)
		Id := strconv.Itoa(M.Id)
		M.URL = "http://119.91.20.70:6060/object?movie_id=" + Id
		M2 = Struct.MovieinWeb{
			Id:           M.Id,
			Introduction: M.Introduction,
			Poster:       M.Poster,
			Year:         M.Year,
			Date:         M.Date,
			Moviename:    M.Moviename,
			Score:        M.Score,
			Language:     M.Language,
			Length:       M.Length,
			Area:         M.Area,
			Type:         M.Type,
			Feature:      M.Feature,
			URL:          M.URL,
			Releasing:    M.Releasing,
		}
		//sqlStr := "select personid from record_direct where pid=?;"
		rows, err := Db.First(&Struct.Record_direct{}).Where("pid=?", M.Id).Rows()
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M.Date = utos(time1)
		M2.Director = make([]Struct.Actorinmovie, 1)
		for rows.Next() {
			err := rows.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.First(&Struct.Person{}).Where("id=?", psid).Scan(&persons)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M2.Director = append(M2.Director, persons)
		}
		//sqlStr2 := "select personid from record_act where pid=?;"
		rows2, err := Db.First(&Struct.Record_act{}).Where("pid=?", M.Id).Rows()
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M2.Actor = make([]Struct.Actorinmovie, 1)
		for rows2.Next() {
			err := rows2.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.First(&Struct.Person{}).Where("id=?", psid).Scan(&persons)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M2.Actor = append(M2.Actor, persons)
		}
		//sqlStr3 := "select personid from record_script where pid=?;"
		rows3, err := Db.First(&Struct.Record_script{}).Where("pid=?", M.Id).Rows()
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M2.Scriptwriter = make([]Struct.Actorinmovie, 1)
		for rows3.Next() {
			err := rows3.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.First(&Struct.Person{}).Where("id=?", psid).Scan(&persons)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M2.Scriptwriter = append(M2.Scriptwriter, persons)
		}
		M2.Director = M2.Director[1:]
		M2.Actor = M2.Actor[1:]
		M2.Scriptwriter = M2.Scriptwriter[1:]
		M1 = append(M1, M2)
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
	rows, err := Db.Raw("select * from movie where moviename like  CONCAT('%',?,'%') ;", stuff).Rows()
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil, nil
	}

	for rows.Next() {
		err := rows.Scan(&M)
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
	rows2, err := Db.Raw("select * from persons like  CONCAT('%',?,'%') ;", stuff).Rows()
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

func QueryNewhotmovie() []Struct.MovieinWeb {
	M1 := make([]Struct.MovieinWeb, 1)
	var M Struct.MovieinWeb
	var M2 Struct.Movie
	var psid int
	var persons Struct.Actorinmovie
	//sqlStr0 := "select id,moviename,yyear,introduction,ddate,posterurl,length,area,type,releasing,feature,score,language from movie where timestampdiff(month ,?,ddate )<3 order by score  desc;"
	rows0, err := Db.Raw("select * from movie where timestampdiff(month ,?,ddate )<3 order by score  desc;").Rows()
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	for rows0.Next() {
		err := rows0.Scan(&M)
		Id := strconv.Itoa(M.Id)
		M = Struct.MovieinWeb{
			Id:           M2.Id,
			Introduction: M2.Introduction,
			Poster:       M2.Poster,
			Year:         M2.Year,
			Date:         M2.Date,
			Moviename:    M2.Moviename,
			Score:        M2.Score,
			Language:     M2.Language,
			Length:       M2.Length,
			Area:         M2.Area,
			Type:         M2.Type,
			Feature:      M2.Feature,
			Releasing:    M2.Releasing,
		}
		M.URL = "http://119.91.20.70:6060/object?movie_id=" + Id
		rows, err := Db.First(&Struct.Record_direct{}).Where("movie_id=?", M.Id).Rows()
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
			Db.First(&Struct.Person{}).Where("id=?", psid).Scan(&persons)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M.Director = append(M.Director, persons)
		}
		//sqlStr2 := "select personid from record_act where pid=?;"
		rows2, err := Db.First(&Struct.Record_act{}).Where("movie_id=?", M.Id).Rows()
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
			//Db.QueryRow("select id,chinesename from person where id=?;", psid).Scan(&persons.Id, &persons.Name)
			Db.First(&Struct.Person{}).Where("id=?", psid).Scan(&persons)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M.Actor = append(M.Actor, persons)
		}
		//sqlStr3 := "select personid from record_script where pid=?;"
		rows3, err := Db.First(&Struct.Record_script{}).Where("movie_id=?", M.Id).Rows()
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
			Db.First(&Struct.Person{}).Where("id=?", psid).Scan(&persons)
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

func Classificationmovie(ty string, area string, year string, feature string) []Struct.MovieinWeb {
	M1 := make([]Struct.MovieinWeb, 1)
	var M Struct.Movie
	var M2 Struct.MovieinWeb
	var psid int
	var persons Struct.Actorinmovie
	var year3 interface{}
	year2, err := strconv.Atoi(year)
	if err != nil {
		fmt.Println(err)
		year3 = "all"
	}
	year3 = year2
	//sqlStr0 := "select id,moviename,yyear,introduction,ddate,posterurl,length,area,type,releasing,feature,score,language from movie where type=? or area=? or feature =? or yyear=?;"
	rows0, err := Db.Model(&Struct.Movie{}).Where("type=? or area=? or feature=? or year=?", ty, area, year3, feature).Rows()
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	for rows0.Next() {
		err := rows0.Scan(&M)
		Id := strconv.Itoa(M.Id)
		M.URL = "http://119.91.20.70:6060/object?movie_id=" + Id
		M2 = Struct.MovieinWeb{
			Id:           M.Id,
			Introduction: M.Introduction,
			Poster:       M.Poster,
			Year:         M.Year,
			Date:         M.Date,
			Moviename:    M.Moviename,
			Score:        M.Score,
			Language:     M.Language,
			Length:       M.Length,
			Area:         M.Area,
			Type:         M.Type,
			Feature:      M.Feature,
			URL:          M.URL,
			Releasing:    M.Releasing,
		}
		//sqlStr := "select personid from record_direct where pid=?;"
		rows, err := Db.First(&Struct.Record_direct{}).Where("pid=?", M2.Id).Rows()
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M2.Director = make([]Struct.Actorinmovie, 1)
		for rows.Next() {
			err := rows.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.First(&Struct.Person{}).Where("id=?", psid).Scan(&persons)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M2.Director = append(M2.Director, persons)
		}

		rows2, err := Db.First(&Struct.Record_act{}).Where("movie_id=?", M.Id).Rows()
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M2.Actor = make([]Struct.Actorinmovie, 1)
		for rows2.Next() {
			err := rows2.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.First(&Struct.Person{}).Where("id=?", psid).Scan(&persons)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M2.Actor = append(M2.Actor, persons)
		}
		rows3, err := Db.First(&Struct.Record_script{}).Where("movie_id=?", M.Id).Rows()
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M2.Scriptwriter = make([]Struct.Actorinmovie, 1)
		for rows3.Next() {
			err := rows3.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.First(&Struct.Person{}).Where("id=?", psid).Scan(&persons)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M2.Scriptwriter = append(M2.Scriptwriter, persons)
		}
		M2.Director = M2.Director[1:]
		M2.Actor = M2.Actor[1:]
		M2.Scriptwriter = M2.Scriptwriter[1:]
		M1 = append(M1, M2)
	}
	M1 = M1[1:]
	return M1
}

func ClassificationListmovie(ty string, area string, year string, feature string) []Struct.MovieinWeb {
	M1 := make([]Struct.MovieinWeb, 1)
	var M Struct.Movie
	var M2 Struct.MovieinWeb
	var psid int
	var persons Struct.Actorinmovie
	var year3 interface{}
	year2, err := strconv.Atoi(year)
	if err != nil {
		fmt.Println(err)
		year3 = "all"
	}
	year3 = year2
	//sqlStr0 := "select id,moviename,yyear,introduction,ddate,posterurl,length,area,type,releasing,feature,score,language from movie where type=? or area=? or feature =? or yyear=? order by score desc ;"
	rows0, err := Db.Model(&Struct.Movie{}).Where("type=? or area=? or feature =? or year=?", ty, area, year3, feature).Rows()
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	for rows0.Next() {
		err := rows0.Scan(&M)
		Id := strconv.Itoa(M.Id)
		M.URL = "http://119.91.20.70:6060/object?movie_id=" + Id
		M2 = Struct.MovieinWeb{
			Id:           M.Id,
			Introduction: M.Introduction,
			Poster:       M.Poster,
			Year:         M.Year,
			Date:         M.Date,
			Moviename:    M.Moviename,
			Score:        M.Score,
			Language:     M.Language,
			Length:       M.Length,
			Area:         M.Area,
			Type:         M.Type,
			Feature:      M.Feature,
			URL:          M.URL,
			Releasing:    M.Releasing,
		}

		rows, err := Db.First(&Struct.Record_direct{}).Where("movie_id=?", M.Id).Rows()
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M2.Director = make([]Struct.Actorinmovie, 1)
		for rows.Next() {
			err := rows.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.First(&Struct.Person{}).Where("id=?", psid).Scan(&persons)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M2.Director = append(M2.Director, persons)
		}
		rows2, err := Db.First(&Struct.Record_act{}).Where("movie_id=?", M.Id).Rows()
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M2.Actor = make([]Struct.Actorinmovie, 1)
		for rows2.Next() {
			err := rows2.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.First(&Struct.Person{}).Where("id=?", psid).Scan(&persons)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M2.Actor = append(M2.Actor, persons)
		}
		rows3, err := Db.First(&Struct.Record_script{}).Where("movie_id=?", M.Id).Rows()
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil
		}
		M2.Scriptwriter = make([]Struct.Actorinmovie, 1)
		for rows3.Next() {
			err := rows3.Scan(&psid)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil
			}
			Db.First(&Struct.Person{}).Where("id=?", psid).Scan(&persons)
			Id = strconv.Itoa(persons.Id)
			persons.URl = "http://119.91.20.70:6060/celebrity?person_id=" + Id
			M2.Scriptwriter = append(M2.Scriptwriter, persons)
		}
		M2.Director = M2.Director[1:]
		M2.Actor = M2.Actor[1:]
		M2.Scriptwriter = M2.Scriptwriter[1:]
		M1 = append(M1, M2)
	}
	M1 = M1[1:]
	return M1
}
