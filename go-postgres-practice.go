package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"

)

const (
	DB_USER = "postgres"
	DB_NAME = "go-first-db"
)

func main() {
	connStr := "user=postgres dbname=go-first-db password=123"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected successfully!")
	}
	defer db.Close()

	// Insert

	var lastInsertedId int
	err = db.QueryRow("INSERT INTO student VALUES($1, $2, $3, $4);", 9331069, "alis", "Computer Eng", 12.96).Scan(&lastInsertedId)
	fmt.Println("last inserted ID:", lastInsertedId)

	stmnt, err := db.Prepare("update student set average=$1 where id=$2;")
	if err != nil {
		panic(err)
	}
	
	_, err = stmnt.Exec(14.20, 9331055)
	if err != nil {
		panic(err)
	}

	//Delete

	stmnt, err = db.Prepare("delete from student where id=$1;")
	if err != nil {
		panic(err)
	}

	_, err = stmnt.Exec(9331055)
	if err != nil {
		panic(err)
	}

	// Query
	rows, err := db.Query("select * from student;")
	if err != nil {
		panic(err)
	}
	 defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var major string
		var avg float32
		err = rows.Scan(&id, &name, &major, &avg)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, name, major, avg)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

}