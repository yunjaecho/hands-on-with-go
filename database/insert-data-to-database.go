package database

import (
	"fmt"
	"database/sql"
)


func init() {
	fmt.Println("===== Inserting Data to Database =====")

	db, err := sql.Open("sqlite3", "personal.db")
	checkError(err)

	statement, err := db.Prepare("insert into Profile(FirstName, LastName, Age) values (?, ?, ?)")
	checkError(err)

	statement.Exec("Jessica2", "McArthur", 30)

	var profile Profile
	rows, err := db.Query("select * from Profile")
	checkError(err)

	for rows.Next() {
		err := rows.Scan(&profile.ProfileId, &profile.FirstName, &profile.LastName, &profile.Age)
		checkError(err)

		fmt.Println(profile)
	}

	statement.Close()
	rows.Close()
	db.Close()
}
