package database

import (
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"database/sql"
)

type Profile struct {
	ProfileId int
	FirstName string
	LastName string
	Age int
}

func checkError(err error) {
	if (err != nil) {
		panic(err)
	}

}

func init() {
	fmt.Println("===== Reading Data from SQL Databases= ====")

	db, err := sql.Open("sqlite3", "personal.db")
	checkError(err)

	var profile Profile
	rows, err := db.Query("select * from Profile where FirstName =? and LastName =? and Age = ? ", "Tarik", "Guney", 31)
	checkError(err)

	for rows.Next() {
		err := rows.Scan(&profile.ProfileId, &profile.FirstName, &profile.LastName, &profile.Age)
		checkError(err)

		fmt.Println(profile)
	}

	rows.Close()
	db.Close()
}

