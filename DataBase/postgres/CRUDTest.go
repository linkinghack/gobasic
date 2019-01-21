package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres","user=linking password=123456 dbname=chitchat sslmode=disable")
	if(err != nil){
		panic(err)
	}
}
func main(){
	statement := `SELECT * FROM posts`
	stmt,err := Db.Prepare(statement)
	stmt.
}


