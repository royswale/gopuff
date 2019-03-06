package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {

	connStr := "postgres://postgres:root@localhost/mylocaldb?sslmode=disable"
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	var name string
	row := db.QueryRow("SELECT name FROM test_table WHERE id = $1", 3)
	err = row.Scan(&name)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(name)
	default:
		panic(err)
	}

}
