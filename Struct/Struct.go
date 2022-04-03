package Struct

import (
	"time"
)

type User struct {
	Id          int
	Username    string
	Password    string
	ProtectionQ string
	ProtectionA string
	Money       int
}

type Message struct {
	Id             int
	Tousername     string
	Fromusername   string
	Time           []uint8
	Messagecontent string
}

type Money struct {
	Id       int
	Touser   string
	Fromuser string
	Time     []uint8
	Howmuch  int
}

type Movie struct {
	Id           int
	Introduction string
	Poster       string
	Year         int
	Date         string
	Moviename    string
	Score        float64
	Language     string
	Length       string
	Area         string
	Type         string
	Feature      string
	URL          string
	Releasing    bool
}

type MovieinWeb struct {
	Id           int
	Introduction string
	Poster       string
	Year         int
	Date         string
	Moviename    string
	Score        float64
	Language     string
	Length       string
	Area         string
	Type         string
	Feature      string
	URL          string
	Releasing    bool
	Actor        []Actorinmovie
	Director     []Actorinmovie
	Scriptwriter []Actorinmovie
}

type Person struct {
	Id             int
	Introduction   string
	Birthday       string
	Constellations string
	Chinesename    string
	Englishname    string
	Birthplace     string
	Jobs           string
	URl            string
	Poster         string
}

type PersoninWeb struct {
	Id             int
	Introduction   string
	Birthday       string
	Constellations string
	Chinesename    string
	Englishname    string
	Birthplace     string
	Jobs           string
	URl            string
	Works          []Movieinactor `gorm:"embedded"`
	Poster         string
}

type Record struct {
	Id       int
	Pid      int
	Personid int
}

type Moviepic struct {
	Id  int
	Pid int
	URL string
}

func (Moviepic) TableName() string {
	return "moviepic"
}

type Comment struct {
	Id            int
	From_id       int
	From_username string
	Content       string
	Score         float64
	Useful        int
	Unuseful      int
	Movieid       int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
type CommentinWeb struct {
	Id            int
	From_id       int
	From_username string
	Content       string
	Score         float64
	Useful        int
	Unuseful      int
	Movieid       int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Child         []Childcomment
}

type Childcomment struct {
	Id            int
	Pid           int
	From_id       int
	From_username string
	Content       string
	Useful        int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Actorinmovie struct {
	Id   int
	Name string
	URl  string
}

type Movieinactor struct {
	Id   int
	Name string
	URl  string
}

type Coperson struct {
	Id   int
	Name string
	URL  string
}

type Shortcomment struct {
	Id            int
	From_id       int
	From_username string
	Content       string
	Lorw          int
	Score         float64
	Useful        int
	Unuseful      int
	Movieid       int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Userimfor struct {
	Id           int
	Username     string
	Introduction string
	Scm          []Scminuser
	Cm           []Cminuser
	Looked       []Movie
	Wanted       []Movie
}

type Scminuser struct {
	Id            int
	From_id       int
	From_username string
	Content       string
	Theday        string
	Lorw          int
	Score         float64
	Usenum        int
	Noues         int
	Movie_id      int
	Movieurl      string
	Moviename     string
}

type Cminuser struct {
	Id            int
	From_id       int
	From_username string
	Content       string
	Score         float64
	Time          string
	Useful        int
	Unuseful      int
	Movieid       int
	Child         []Childcomment `gorm:"embedded"`
	Movieurl      string
	Moviename     string
}

type Talking struct {
	Id       int
	MovieID  int
	Title    string
	Content  string
	Username string
	Userid   int
	Theday   string
	Url      string
	Cm       []Tkcm
}

type Tkcm struct {
	Id         int
	Talking_id int
	Username   string
	Userid     int
	Content    string
	Theday     string
	Child      []Tkchild
}

type Tkchild struct {
	Id       int
	Pid      int
	Username string
	Userid   int
	Content  string
	Theday   string
}

type Tkinmovie struct {
	Id       int
	MovieID  int
	Title    string
	Username string
	Userid   int
	Thedat   string
	URL      string
}

type Record_direct struct {
	Id        int
	Movie_id  int `gorm:"column=movie_id"`
	Person_id int `gorm:"column=person_id"`
}

func (Record_direct) TableName() string {
	return "record_direct"
}

type Record_script struct {
	Id        int
	Movie_id  int `gorm:"column=movie_id"`
	Person_id int `gorm:"column=person_id"`
}

func (Record_script) TableName() string {
	return "record_direct"
}

type Record_all struct {
	Id        int
	Movie_id  int `gorm:"column=movie_id"`
	Person_id int `gorm:"column=person_id"`
}

func (Record_all) TableName() string {
	return "record_direct"
}

type Record_act struct {
	Id        int
	Movie_id  int `gorm:"column=movie_id"`
	Person_id int `gorm:"column=person_id"`
}

func (Record_act) TableName() string {
	return "record_direct"
}
