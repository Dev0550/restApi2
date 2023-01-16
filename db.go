package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type movie struct {
	id        int
	movieid   string
	moviename string
}

func main() {

	connStr := "user=postgres password=123456 dbname=testDB sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//insert query
	// result, err := db.Exec("insert into movies (id, movieid, moviename) values ('4', '4', 'Batman')")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(result.LastInsertId()) // не поддерживается
	// fmt.Println(result.RowsAffected()) // количество добавленных строк

	//select query
	rows, err := db.Query("select * from movies where id >= 4")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	movies := []movie{}

	for rows.Next() {
		p := movie{}
		err := rows.Scan(&p.id, &p.movieid, &p.moviename)
		if err != nil {
			fmt.Println(err)
			continue
		}
		movies = append(movies, p)
	}
	for _, p := range movies {
		fmt.Println(p.id, p.movieid, p.moviename)
	}
}
