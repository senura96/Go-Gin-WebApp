package models

type User struct {
	Id        int64  `db:"ID" json:"id"`
	Username  string `db:"Username" json:"username"`
	Password  string `db:"Password" json:"password"`
	Firstname string `db:"Firstname" json:"firstname"`
	Lastname  string `db:"Lastname" json:"lastname"`
}

type Posts struct {
	Userid int64  `db:"Userid" json:"userid"`
	Postid int64  `db:"Postid" json:"postid"`
	Title  string `db:"Title" json:"title"`
	Body   string `db:"Body" json:"body"`
}
