package dao

import (
	"fmt"
	"message-board/Struct"
)

func Insertcomment(cm Struct.Comment) error {
	//_, err := Db.Exec("insert into comment(from_username, from_id, Content, theday, usenum, unusenum, score,movie_id) values(?,?,?,?,0,0,?,?);", cm.From_username, cm.From_id, cm.Content, time, cm.Score, cm.Movieid)
	res := Db.Create(cm)
	Scoredao(cm.Movieid)
	if res.Error != nil {
		fmt.Println(res.Error)
		return res.Error
	}
	return res.Error
}

func Insertchcomment(chcm Struct.Childcomment) bool {
	var use int
	var unuse int
	if chcm.Useful == 1 {
		//err := Db.QueryRow("select usenum from  comment where id=?;", pid).Scan(&use)
		res := Db.Select("useful").Where("id=?", chcm.Pid).First(&Struct.Comment{}).Scan(&use)
		if res.Error != nil {
			fmt.Println(res.Error)
			return false
		}
		use++
		//_, err = Db.Exec("update comment set usenum=? where id=?;", use, pid)
		res = Db.Model(&Struct.Comment{}).Update("useful", use).Where("id=?", chcm.Pid)
		if res.Error != nil {
			fmt.Println(res.Error)
			return false
		}
	} else if chcm.Useful == 0 {
		//err := Db.QueryRow("select unusenum from  comment where id=?;", pid).Scan(&unuse)
		res := Db.Select("unuseful").First(&Struct.Comment{}).Where("id=?", chcm.Pid).Scan(&unuse)
		if res.Error != nil {
			fmt.Println(res.Error)
			return false
		}
		unuse++
		//= Db.Exec("update comment set unusenum=? where id=?;", unuse, pid)
		res = Db.Model(&Struct.Comment{}).Set("unuseful", unuse).Where("id=?", chcm.Pid)
		if res.Error != nil {
			fmt.Println(res.Error)
			return false
		}
	} else {
		fmt.Println("用户绕过客户端！！！，本次操作不予执行")
		return false
	}

	//_, err := Db.Exec("insert into childcomment(pid, from_id, from_username, content, theday, Useful) values(?,?,?,?,?,?);", pid, from_id, from_username, content, time, useful)
	Db.Create(&chcm)
	return true
}

func Queryusermoviecm(movieid int) ([]Struct.CommentinWeb, error) {
	cm := make([]Struct.CommentinWeb, 1)
	var cm1 Struct.Comment
	var cm2 Struct.CommentinWeb
	var chcm2 Struct.Childcomment
	//sqlStr := "select id,from_username, from_id,Content, theday, usenum, unusenum,score,movie_id from comment where movie_id = ? order by theday desc ;" //遍历写给登录用户的评论
	rows, err := Db.Model(&Struct.Comment{}).Where("movie_id=?", movieid).Order("created_at desc").Rows()
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		cm2.Child = make([]Struct.Childcomment, 0)
		err = rows.Scan(&cm1)
		cm2 = Struct.CommentinWeb{
			Id:            cm1.Id,
			From_id:       cm1.From_id,
			From_username: cm1.From_username,
			Content:       cm1.Content,
			Score:         cm1.Score,
			Useful:        cm1.Useful,
			Unuseful:      cm1.Unuseful,
			Movieid:       cm1.Movieid,
			CreatedAt:     cm1.CreatedAt,
			UpdatedAt:     cm1.UpdatedAt,
		}
		cm2.Child = make([]Struct.Childcomment, 0)
		if err != nil {
			fmt.Println("scan failed, err:%v\n", err)
			return nil, err
		}
		//okk, _ := Db.Query("select id, pid, from_id, from_username, content, theday, Useful from childcomment where pid=?;", cm2.Id)
		rows2, _ := Db.Model(&Struct.Childcomment{}).Where("pid=?", cm2.Id).Rows()
		for rows2.Next() {
			err = rows2.Scan(&chcm2)
			if err != nil {
				fmt.Println("scan failed, err:%v\n", err)
				return nil, err
			}
			cm2.Child = append(cm2.Child, chcm2)
		}
		rows2.Close()
		cm = append(cm, cm2)

	}

	rows.Close()
	cm = cm[1:]
	return cm, err
}

func QueryusermoviecmbyUse(movieid int) ([]Struct.CommentinWeb, error) {
	cm := make([]Struct.CommentinWeb, 1)
	var cm1 Struct.Comment
	var cm2 Struct.CommentinWeb
	var chcm2 Struct.Childcomment

	//sqlStr := "select id,from_username, from_id,Content, theday, usenum, unusenum,score,movie_id from comment where movie_id = ? order by usenum desc ;" //遍历写给登录用户的评论
	rows, err := Db.Model(&Struct.Comment{}).Where("movieid=?", movieid).Order("useful desc").Rows()
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err = rows.Scan(&cm1)
		if err != nil {
			fmt.Println("scan failed, err:%v\n", err)
			return nil, err
		}
		cm2 = Struct.CommentinWeb{
			Id:            cm1.Id,
			From_id:       cm1.From_id,
			From_username: cm1.From_username,
			Content:       cm1.Content,
			Score:         cm1.Score,
			Useful:        cm1.Useful,
			Unuseful:      cm1.Unuseful,
			Movieid:       cm1.Movieid,
			CreatedAt:     cm1.CreatedAt,
			UpdatedAt:     cm1.UpdatedAt,
		}
		cm2.Child = make([]Struct.Childcomment, 0)
		//rows2, _ := Db.Query("select id, pid, from_id, from_username, content, theday, Useful from childcomment where pid=?;", cm2.Id)
		rows2, _ := Db.Model(&Struct.Childcomment{}).Where("pid=?", cm1.Id).Rows()
		for rows2.Next() {
			err = rows2.Scan(&chcm2)
			if err != nil {
				fmt.Println("scan failed, err:%v\n", err)
				return nil, err
			}
			cm2.Child = append(cm2.Child, chcm2)
		}
		rows2.Close()
		cm = append(cm, cm2)

	}

	rows.Close()
	cm = cm[1:]
	return cm, err
}

func QueryusermoviecmbyUsebyLimit(movieid int) ([]Struct.CommentinWeb, error) {
	cm := make([]Struct.CommentinWeb, 1)
	var cm1 Struct.Comment
	var cm2 Struct.CommentinWeb
	var chcm2 Struct.Childcomment
	//sqlStr := "select id,from_username, from_id,Content, theday, usenum, unusenum,score,movie_id from comment where movie_id = ? order by usenum desc limit 5;" //遍历写给登录用户的评论
	rows, err := Db.Model(&Struct.Comment{}).Where("movieid=?", movieid).Rows()
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {

		err = rows.Scan(&cm1)
		if err != nil {
			fmt.Println("scan failed, err:%v\n", err)
			return nil, err
		}
		cm2 = Struct.CommentinWeb{
			Id:            cm1.Id,
			From_id:       cm1.From_id,
			From_username: cm1.From_username,
			Content:       cm1.Content,
			Score:         cm1.Score,
			Useful:        cm1.Useful,
			Unuseful:      cm1.Unuseful,
			Movieid:       cm1.Movieid,
			CreatedAt:     cm1.CreatedAt,
			UpdatedAt:     cm1.UpdatedAt,
		}
		cm2.Child = make([]Struct.Childcomment, 0)
		//rows2, _ := Db.Query("select id, pid, from_id, from_username, content, theday, Useful from childcomment where pid=?;", cm2.Id)
		rows2, _ := Db.Model(&Struct.Childcomment{}).Where("pid=?", cm2.Id).Rows()
		for rows2.Next() {
			err = rows2.Scan(&chcm2)
			if err != nil {
				fmt.Println("scan failed, err:%v\n", err)
				return nil, err
			}
			cm2.Child = append(cm2.Child, chcm2)
		}
		rows2.Close()
		cm = append(cm, cm2)

	}

	rows.Close()
	cm = cm[1:]
	return cm, err
}

func Querycomment(pid int) bool {
	var content string
	//err := Db.QueryRow("select Content from comment where id = ?;", pid).Scan(&content)
	res := Db.Model(&Struct.Comment{}).Select("content").Where("id=?", pid).Scan(content)
	if res.Error != nil {
		fmt.Println("查询错误", res.Error)
		return false
	}
	return true
}

func Insertshortcomment(scm Struct.Shortcomment) bool {
	//_, err := Db.Exec("insert into shortcomment(from_username, from_id, content, theday, lorw, score, movie_id,usenum,nouse) values(?,?,?,?,?,?,?,0,0);", from_username, from_id, content, time, lorw, score, movie_id)
	res := Db.Create(&scm)
	if res.Error != nil {
		fmt.Println(res.Error)
		return false
	}
	Scoredao(scm.Movieid)
	return true
}

func QueryshortcommentbyTime(movieid int) []Struct.Shortcomment {
	scm := make([]Struct.Shortcomment, 1)
	var scm1 Struct.Shortcomment
	//sqlStr := "select id,from_username, from_id,Content, theday, usenum,nouse,score,movie_id from shortcomment where movie_id = ? order by theday desc ;" //遍历写给登录用户的评论
	rows, err := Db.Model(&scm).Where("movieid=?", movieid).Order("createted_at desc").Rows()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		err := rows.Scan(scm1)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		scm = append(scm, scm1)
	}
	scm = scm[1:]
	return scm
}

func QueryshortcommentbyUse(movieid int) []Struct.Shortcomment {
	scm := make([]Struct.Shortcomment, 1)
	var scm1 Struct.Shortcomment
	//sqlStr := "select id,from_username, from_id,Content, theday, usenum,nouse,score,movie_id from shortcomment where movie_id = ? order by usenum desc ;" //遍历写给登录用户的评论
	rows, err := Db.Model(&Struct.Shortcomment{}).Where("movieid=?", movieid).Order("useful desc").Rows()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		err := rows.Scan(&scm1)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		scm = append(scm, scm1)
	}
	scm = scm[1:]
	return scm
}

func QueryshortcommentbyUsebyLimit(movieid int) []Struct.Shortcomment {
	scm := make([]Struct.Shortcomment, 1)
	var scm1 Struct.Shortcomment
	//sqlStr := "select id,from_username, from_id,Content, theday, usenum,nouse,score,movie_id from shortcomment where movie_id = ? order by usenum desc limit 5;" //遍历写给登录用户的评论
	rows, err := Db.Model(&Struct.Shortcomment{}).Where("movieid=?", movieid).Order("created_at desc").Limit(10).Rows()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		err := rows.Scan(&scm1)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		scm = append(scm, scm1)
	}
	scm = scm[1:]
	return scm
}

func QuerycommentwithoutChild(movieid int) []Struct.Comment {
	cm := make([]Struct.Comment, 1)
	var cm1 Struct.Comment
	//sqlStr := "select id, from_username, from_id, Content, theday, usenum, unusenum, score, movie_id from comment where movie_id = ? order by usenum desc limit 10;"
	rows, err := Db.Model(&Struct.Comment{}).Where("movieid=?", movieid).Order("useful desc").Rows()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		err := rows.Scan(&cm1)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}

		cm = append(cm, cm1)
	}
	cm = cm[1:]
	return cm

}

func Scoredao(id int) {
	var score1 float64
	var score2 float64
	//err := Db.QueryRow("select AVG(score) from comment where movie_id = ?;", id).Scan(&score1)
	res := Db.Find(&Struct.Comment{}).Where("movieid=?", id).Select("AVG(score)").Scan(&score1)
	if res.Error != nil {
		score1 = 0.0
	}
	res = Db.Find(&Struct.Shortcomment{}).Where("movieid=?", id).Select("AVG(score)").Scan(&score2)
	if res.Error != nil {
		score2 = 0.0
	}
	Score := (score2 + score1) / 2
	//err2, _ := Db.Exec("update movie set score=? where id =?;", Score, id)
	res = Db.Model(&Struct.Movie{}).Update("score", Score).Where("id=?", id)
	if res.Error != nil {
		fmt.Println(res.Error)
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
	//Db.QueryRow("select usenum,nouse from shortcomment where id=?;", id).Scan(&use, &nouse)
	var scm Struct.Shortcomment
	Db.First(&Struct.Shortcomment{}).Select("useful").Where("id=?", id).Scan(&scm)
	return scm.Useful, scm.Unuseful
}

func Updateshortuse(id int, usenum int) bool {
	usenum++
	//_, err := Db.Exec("update shortcomment set usenum=? where id =?;", usenum, id)
	res := Db.Model(&Struct.Shortcomment{}).Update("useful", usenum).Where("id=?", id)
	if res.Error != nil {
		fmt.Println(res.Error)
		return false
	}
	return true
}

func Updateshortnouse(id int, nouse int) bool {
	nouse++
	//_, err := Db.Exec("update shortcomment set nouse=? where id =?;", nouse, id)
	res := Db.Model(&Struct.Shortcomment{}).Update("unuseful", nouse).Where("id=?", id)
	if res.Error != nil {
		fmt.Println(res.Error)
		return false
	}
	return true
}
