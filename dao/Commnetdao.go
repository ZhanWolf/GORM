package dao

import (
	"fmt"
	"message-board/Struct"
	"time"
)

func Insertcomment(cm Struct.Comment) error {
	time := time.Now()
	_, err := Db.Exec("insert into comment(from_username, from_id, Content, theday, usenum, unusenum, score,movie_id) values(?,?,?,?,0,0,?,?);", cm.From_username, cm.From_id, cm.Content, time, cm.Score, cm.Movieid)
	Scoredao(cm.Movieid)
	return err
}

func Insertchcomment(pid int, from_id int, from_username string, content string, useful int) bool {
	time := time.Now()
	var use int
	var unuse int
	if useful == 1 {
		err := Db.QueryRow("select usenum from  comment where id=?;", pid).Scan(&use)
		use++
		_, err = Db.Exec("update comment set usenum=? where id=?;", use, pid)
		if err != nil {
			fmt.Println(err)
			return false
		}
	} else if use == 0 {
		err := Db.QueryRow("select unusenum from  comment where id=?;", pid).Scan(&unuse)
		unuse++
		_, err = Db.Exec("update comment set unusenum=? where id=?;", unuse, pid)
		if err != nil {
			fmt.Println(err)
			return false
		}
	} else {
		fmt.Println("用户绕过客户端！！！，本次操作不予执行")
		return false
	}

	_, err := Db.Exec("insert into childcomment(pid, from_id, from_username, content, theday, Useful) values(?,?,?,?,?,?);", pid, from_id, from_username, content, time, useful)
	if err != nil {
		fmt.Println(err)
	}
	return true
}

func Queryusermoviecm(movieid int) ([]Struct.Comment, error) {
	cm := make([]Struct.Comment, 1)
	var cm2 Struct.Comment
	var chcm2 Struct.Childcomment
	var time1 []uint8
	var time2 []uint8

	sqlStr := "select id,from_username, from_id,Content, theday, usenum, unusenum,score,movie_id from comment where movie_id = ? order by theday desc ;" //遍历写给登录用户的评论
	rows, err := Db.Query(sqlStr, movieid)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		cm2.Child = make([]Struct.Childcomment, 0)
		err = rows.Scan(&cm2.Id, &cm2.From_username, &cm2.From_id, &cm2.Content, &time1, &cm2.Useful, &cm2.Unuseful, &cm2.Score, &cm2.Movieid)
		if err != nil {
			fmt.Println("scan failed, err:%v\n", err)
			return nil, err
		}
		cm2.Time = utos(time1)
		okk, _ := Db.Query("select id, pid, from_id, from_username, content, theday, Useful from childcomment where pid=?;", cm2.Id)
		for okk.Next() {
			err = okk.Scan(&chcm2.Id, &chcm2.Pid, &chcm2.From_id, &chcm2.From_username, &chcm2.Content, &time2, &chcm2.Useful)
			if err != nil {
				fmt.Println("scan failed, err:%v\n", err)
				return nil, err
			}
			chcm2.Time = utos(time2)
			cm2.Child = append(cm2.Child, chcm2)
		}
		okk.Close()
		cm = append(cm, cm2)

	}

	rows.Close()
	cm = cm[1:]
	return cm, err
}

func QueryusermoviecmbyUse(movieid int) ([]Struct.Comment, error) {
	cm := make([]Struct.Comment, 1)
	var cm2 Struct.Comment
	var chcm2 Struct.Childcomment
	var time1 []uint8
	var time2 []uint8

	sqlStr := "select id,from_username, from_id,Content, theday, usenum, unusenum,score,movie_id from comment where movie_id = ? order by usenum desc ;" //遍历写给登录用户的评论
	rows, err := Db.Query(sqlStr, movieid)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		cm2.Child = make([]Struct.Childcomment, 0)
		err = rows.Scan(&cm2.Id, &cm2.From_username, &cm2.From_id, &cm2.Content, &time1, &cm2.Useful, &cm2.Unuseful, &cm2.Score, &cm2.Movieid)
		if err != nil {
			fmt.Println("scan failed, err:%v\n", err)
			return nil, err
		}
		cm2.Time = utos(time1)
		okk, _ := Db.Query("select id, pid, from_id, from_username, content, theday, Useful from childcomment where pid=?;", cm2.Id)
		for okk.Next() {
			err = okk.Scan(&chcm2.Id, &chcm2.Pid, &chcm2.From_id, &chcm2.From_username, &chcm2.Content, &time2, &chcm2.Useful)
			if err != nil {
				fmt.Println("scan failed, err:%v\n", err)
				return nil, err
			}
			chcm2.Time = utos(time2)
			cm2.Child = append(cm2.Child, chcm2)
		}
		okk.Close()
		cm = append(cm, cm2)

	}

	rows.Close()
	cm = cm[1:]
	return cm, err
}

func QueryusermoviecmbyUsebyLimit(movieid int) ([]Struct.Comment, error) {
	cm := make([]Struct.Comment, 1)
	var cm2 Struct.Comment
	var chcm2 Struct.Childcomment
	var time1 []uint8
	var time2 []uint8
	sqlStr := "select id,from_username, from_id,Content, theday, usenum, unusenum,score,movie_id from comment where movie_id = ? order by usenum desc limit 5;" //遍历写给登录用户的评论
	rows, err := Db.Query(sqlStr, movieid)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		cm2.Child = make([]Struct.Childcomment, 0)
		err = rows.Scan(&cm2.Id, &cm2.From_username, &cm2.From_id, &cm2.Content, &time1, &cm2.Useful, &cm2.Unuseful, &cm2.Score, &cm2.Movieid)
		if err != nil {
			fmt.Println("scan failed, err:%v\n", err)
			return nil, err
		}
		cm2.Time = utos(time1)
		okk, _ := Db.Query("select id, pid, from_id, from_username, content, theday, Useful from childcomment where pid=?;", cm2.Id)
		for okk.Next() {
			err = okk.Scan(&chcm2.Id, &chcm2.Pid, &chcm2.From_id, &chcm2.From_username, &chcm2.Content, &time2, &chcm2.Useful)
			if err != nil {
				fmt.Println("scan failed, err:%v\n", err)
				return nil, err
			}
			chcm2.Time = utos(time2)
			cm2.Child = append(cm2.Child, chcm2)
		}
		okk.Close()
		cm = append(cm, cm2)

	}

	rows.Close()
	cm = cm[1:]
	return cm, err
}

func Querycomment(pid int) bool {
	var content string
	err := Db.QueryRow("select Content from comment where id = ?;", pid).Scan(&content)
	if err != nil {
		fmt.Println("查询错误", err)
		return false
	}
	return true
}

func Insertshortcomment(from_username string, from_id int, content string, lorw int, score float64, movie_id int) bool {
	time := time.Now()
	_, err := Db.Exec("insert into shortcomment(from_username, from_id, content, theday, lorw, score, movie_id,usenum,nouse) values(?,?,?,?,?,?,?,0,0);", from_username, from_id, content, time, lorw, score, movie_id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	Scoredao(movie_id)
	return true
}

func QueryshortcommentbyTime(movieid int) []Struct.Shortcomment {
	scm := make([]Struct.Shortcomment, 1)
	var scm1 Struct.Shortcomment
	var time1 []uint8
	sqlStr := "select id,from_username, from_id,Content, theday, usenum,nouse,score,movie_id from shortcomment where movie_id = ? order by theday desc ;" //遍历写给登录用户的评论
	rows, err := Db.Query(sqlStr, movieid)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		err := rows.Scan(&scm1.Id, &scm1.From_username, &scm1.From_id, &scm1.Content, &time1, &scm1.Usenum, &scm1.Noues, &scm1.Score, &scm1.Movie_id)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		scm1.Theday = utos(time1)
		scm = append(scm, scm1)
	}
	scm = scm[1:]
	return scm
}

func QueryshortcommentbyUse(movieid int) []Struct.Shortcomment {
	scm := make([]Struct.Shortcomment, 1)
	var scm1 Struct.Shortcomment
	var time1 []uint8
	sqlStr := "select id,from_username, from_id,Content, theday, usenum,nouse,score,movie_id from shortcomment where movie_id = ? order by usenum desc ;" //遍历写给登录用户的评论
	rows, err := Db.Query(sqlStr, movieid)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		err := rows.Scan(&scm1.Id, &scm1.From_username, &scm1.From_id, &scm1.Content, &time1, &scm1.Usenum, &scm1.Noues, &scm1.Score, &scm1.Movie_id)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		scm1.Theday = utos(time1)
		scm = append(scm, scm1)
	}
	scm = scm[1:]
	return scm
}

func QueryshortcommentbyUsebyLimit(movieid int) []Struct.Shortcomment {
	scm := make([]Struct.Shortcomment, 1)
	var scm1 Struct.Shortcomment
	var time1 []uint8
	sqlStr := "select id,from_username, from_id,Content, theday, usenum,nouse,score,movie_id from shortcomment where movie_id = ? order by usenum desc limit 5;" //遍历写给登录用户的评论
	rows, err := Db.Query(sqlStr, movieid)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		err := rows.Scan(&scm1.Id, &scm1.From_username, &scm1.From_id, &scm1.Content, &time1, &scm1.Usenum, &scm1.Noues, &scm1.Score, &scm1.Movie_id)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		scm1.Theday = utos(time1)
		scm = append(scm, scm1)
	}
	scm = scm[1:]
	return scm
}

func QuerycommentwithoutChild(movieid int) []Struct.Comment {
	cm := make([]Struct.Comment, 1)
	var cm1 Struct.Comment
	var time1 []uint8
	sqlStr := "select id, from_username, from_id, Content, theday, usenum, unusenum, score, movie_id from comment where movie_id = ? order by usenum desc limit 10;"
	rows, err := Db.Query(sqlStr, movieid)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		err := rows.Scan(&cm1.Id, &cm1.From_username, &cm1.From_id, &cm1.Content, &time1, &cm1.Useful, &cm1.Unuseful, &cm1.Score, &cm1.Movieid)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		cm1.Time = utos(time1)
		cm = append(cm, cm1)
	}
	cm = cm[1:]
	return cm

}

func Scoredao(id int) {
	var score1 float64
	var score2 float64
	err := Db.QueryRow("select AVG(score) from comment where movie_id = ?;", id).Scan(&score1)
	if err != nil {
		score1 = 0.0
	}
	err = Db.QueryRow("select AVG(score) from shortcomment where movie_id = ?;", id).Scan(&score2)
	if err != nil {
		score2 = 0.0
	}
	Score := (score2 + score1) / 2
	err2, _ := Db.Exec("update movie set score=? where id =?;", Score, id)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
}

func utos(u []uint8) string {
	by := []byte{}
	for _, b := range u {
		by = append(by, b)
	}
	return string(by)
}

func Queryshortusenum(id int) (int, int) {
	var use int
	var nouse int
	Db.QueryRow("select usenum,nouse from shortcomment where id=?;", id).Scan(&use, &nouse)
	return use, nouse
}

func Updateshortuse(id int, usenum int) bool {
	usenum++
	_, err := Db.Exec("update shortcomment set usenum=? where id =?;", usenum, id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func Updateshortnouse(id int, nouse int) bool {
	nouse++
	_, err := Db.Exec("update shortcomment set nouse=? where id =?;", nouse, id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
