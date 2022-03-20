package dao

import (
	"fmt"
	"message-board/Struct"
	"strconv"
	"time"
)

func Inserttalking(username string, userid int, content string, title string, movie_id int) bool {
	time := time.Now()

	_, err := Db.Exec("insert into talking(username, userid, content, theday,title,movie_id) values(?,?,?,?,?,?);", username, userid, content, time, title, movie_id)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func Inserttkcomment(talking_id int, content string, username string, userid int) error {
	time := time.Now()

	_, err := Db.Exec("insert into comment2(talking_id, content, username, userid, theday) values(?,?,?,?,?);", talking_id, content, username, userid, time)

	return err
}

func Inserttkchcomment(pid int, content string, username string, userid int) error {
	time := time.Now()

	_, err := Db.Exec("insert into childcomment2(pid, content, username, userid, theday) values(?,?,?,?,?);", pid, content, username, userid, time)

	return err
}

func Querytalking(Id int) Struct.Talking {
	var Talking Struct.Talking
	var time []uint8
	cm := make([]Struct.Tkcm, 1)
	var cm1 Struct.Tkcm
	cm1.Child = make([]Struct.Tkchild, 0)
	ch := make([]Struct.Tkchild, 1)
	var ch1 Struct.Tkchild
	err := Db.QueryRow("select ID, MOVIE_ID, TITLE, CONTENT, THEDAY, USERNAME, USERID from talking where id =?;", Id).Scan(&Talking.Id, &Talking.MovieID, &Talking.Title, &Talking.Content, &time, &Talking.Username, &Talking.Userid)
	time2 := utos(time)
	Talking.Theday = time2
	Id2 := strconv.Itoa(Id)
	Talking.Url = "http://119.91.20.70:6060/talking?talking_id=" + Id2
	if err != nil {
		fmt.Println(err)
		return Talking
	}
	sqlStr := "select id, talking_id, content, username, userid, theday from comment2 where talking_id= ? ;" //遍历写给登录用户的评论
	rows, err := Db.Query(sqlStr, Id)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err = rows.Scan(&cm1.Id, &cm1.Talking_id, &cm1.Content, &cm1.Username, &cm1.Userid, &time)
		if err != nil {
			fmt.Println(err)
			return Talking
		}
		time2 = utos(time)
		cm1.Theday = time2
		sqlStr2 := "select id, pid,content, username, userid, theday from childcomment2 where pid= ? ;" //遍历写给登录用户的评论
		rows2, err := Db.Query(sqlStr2, cm1.Id)
		if err != nil {
			fmt.Println(err)
		}
		for rows2.Next() {
			err = rows2.Scan(&ch1.Id, &ch1.Pid, &ch1.Content, &ch1.Username, &ch1.Userid, &time)
			if err != nil {
				fmt.Println(err)
				return Talking
			}
			time2 = utos(time)
			ch1.Theday = time2
			ch = append(ch, ch1)
		}
		ch = ch[1:]
		cm1.Child = ch
		cm = append(cm, cm1)
	}
	cm = cm[1:]
	Talking.Cm = cm
	return Talking
}

func Querytalkinginmovie(movieid int) []Struct.Tkinmovie {
	talkings := make([]Struct.Tkinmovie, 1)
	var talking Struct.Tkinmovie
	sqlStr := "select id, movie_id, title, theday, username, userid from talking where movie_id= ? ;" //遍历写给登录用户的评论
	rows, err := Db.Query(sqlStr, movieid)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err = rows.Scan(&talking.Id, &talking.MovieID, &talking.Title, &talking.Thedat, &talking.Username, &talking.Userid)
		Id2 := strconv.Itoa(talking.Id)
		talking.URL = "http://119.91.20.70:6060/talking?talking_id=" + Id2
		if err != nil {
			fmt.Println(err)
		}
		talkings = append(talkings, talking)
	}
	talkings = talkings[1:]
	return talkings
}
