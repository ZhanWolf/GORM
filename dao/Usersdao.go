package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"message-board/Struct"
)

func Queryuserpassword(username string) string {
	U := new(Struct.User)
	err := Db.QueryRow("select username,password,id from user where username = ?;", username).Scan(&U.Username, &U.Password, &U.Id)
	if err != nil {
		fmt.Println("错误:", err)
	}

	return U.Password
}

func Queryusername(username string) (int, error) {
	OpenDb()
	U := new(Struct.User)
	err := Db.QueryRow("select id from user where username = ?;", username).Scan(&U.Id)
	if err != nil {
		fmt.Println("查询错误", err)
		return 0, err
	}
	return U.Id, nil
}

func Queryuserid(id int) (string, error) {
	OpenDb()
	U := new(Struct.User)
	err := Db.QueryRow("select username from user where id = ?;", id).Scan(&U.Username)
	if err != nil {
		fmt.Println("查询错误", err)
		return "", err
	}
	return U.Username, nil
}

func Insertuser(username string, password string, protectionQ string, protectionA string) error {
	_, err := Db.Exec("insert into user(username,password,protectionQ,protectionA,introduction) values (?,?,?,?,0);", username, password, protectionQ, protectionA)
	if err != nil {
		fmt.Println("插入错误", err)
	}
	return nil
}

func Updatepassword(newpassword string, username string) error {
	_, err := Db.Exec("update user set password=? where username=?;", newpassword, username)
	return err
}

func Queryprotection(username string) (string, string) {
	U := new(Struct.User)
	err := Db.QueryRow("select protectionQ,protectionA from user where username=?;", username).Scan(&U.ProtectionQ, &U.ProtectionA)
	if err != nil {
		fmt.Println(err)
	}
	return U.ProtectionQ, U.ProtectionA
}

func UpdateIntroduction(introduction string, id int) {
	_, err := Db.Exec("update  user set introduction=? where id =?;", introduction, id)
	if err != nil {
		fmt.Println("插入错误", err)
	}

	return
}
func Queryintroducton(from_id int) (string, error) {
	var introduction string
	err := Db.QueryRow("select introduction from user where id=?", from_id).Scan(&introduction)
	if err != nil {
		fmt.Println(err)
		return " ", err
	}
	return introduction, nil
}

func QueryUserscm(from_id int) []Struct.Scminuser {
	scm := make([]Struct.Scminuser, 1)
	var scm1 Struct.Scminuser
	var time1 []uint8
	sqlStr := "select id, from_username, from_id, content, theday, lorw, score, usenum, nouse, movie_id from shortcomment where from_id = ? order by theday desc ;" //遍历写给登录用户的评论
	rows, err := Db.Query(sqlStr, from_id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		err := rows.Scan(&scm1.Id, &scm1.From_username, &scm1.From_id, &scm1.Content, &scm1.Theday, &scm1.Lorw, &scm1.Score, &scm1.Usenum, &scm1.Noues, &scm1.Movie_id)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		scm1.Theday = utos(time1)
		Db.QueryRow("select moviename from movie where id =?", scm1.Movie_id).Scan(&scm1.Moviename)
		scm = append(scm, scm1)
	}
	scm = scm[1:]
	return scm
}

func QueryUsercm(from_id int) []Struct.Cminuser {
	scm := make([]Struct.Cminuser, 1)
	var scm1 Struct.Cminuser
	var time1 []uint8
	sqlStr := "select id, from_username, from_id, content, theday, score, usenum, unusenum, movie_id from comment where from_id = ? order by theday desc ;" //遍历写给登录用户的评论
	rows, err := Db.Query(sqlStr, from_id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		err := rows.Scan(&scm1.Id, &scm1.From_username, &scm1.From_id, &scm1.Content, &scm1.Time, &scm1.Score, &scm1.Useful, &scm1.Unuseful, &scm1.Movieid)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		scm1.Time = utos(time1)
		Db.QueryRow("select moviename from movie where id =?", scm1.Movieid).Scan(&scm1.Moviename)
		scm = append(scm, scm1)
	}
	scm = scm[1:]
	return scm
}

func Looked(from_id int) []Struct.Movie {
	var movieid int
	var Movie Struct.Movie
	Movieslice := make([]Struct.Movie, 1)
	sqlStr := "select movie_id from shortcomment where from_id = ? and lorw = 1;" //遍历写给登录用户的评论
	rows, err := Db.Query(sqlStr, from_id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		err := rows.Scan(&movieid)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		err = Db.QueryRow("select id, moviename, yyear, introduction, ddate, posterurl length, area, type, feature, releasing, score from movie where id =?", movieid).Scan(&Movie.Id, &Movie.Moviename, &Movie.Year, &Movie.Introduction, &Movie.Date, &Movie.Poster, &Movie.Length, &Movie.Area, &Movie.Type, &Movie.Feature, &Movie.Releasing, &Movie.Score)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		Movieslice = append(Movieslice, Movie)
	}
	Movieslice = Movieslice[1:]
	return Movieslice

}

func Wanted(from_id int) []Struct.Movie {
	var movieid int
	var Movie Struct.Movie
	Movieslice := make([]Struct.Movie, 1)
	sqlStr := "select movie_id from shortcomment where from_id = ? and lorw = 0;"
	rows, err := Db.Query(sqlStr, from_id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		err := rows.Scan(&movieid)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}

		Db.QueryRow("select id, moviename, yyear, introduction, ddate, posterurl, length, area, type, feature, releasing, score from movie where id =?", movieid).Scan(&Movie.Id, &Movie.Moviename, &Movie.Year, &Movie.Introduction, &Movie.Date, &Movie.Poster, &Movie.Length, &Movie.Area, &Movie.Type, &Movie.Feature, &Movie.Releasing, &Movie.Score)
		Movieslice = append(Movieslice, Movie)
	}
	Movieslice = Movieslice[1:]
	return Movieslice

}

func Queryuserbygit(github int) (id int, username string) {
	err := Db.QueryRow("select id, username from user where github_id =?", github).Scan(&id, &username)
	if err != nil {
		fmt.Println(err)
		return 0, ""
	}
	return id, username
}
