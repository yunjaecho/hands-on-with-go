package database

import (
	"fmt"
	"database/sql"
)

func init() {
	fmt.Println("===== Deleting Data from Database =====")

	db, err := sql.Open("sqlite3", "personal.db")
	checkError(err)

	statement, err := db.Prepare("delete from Profile where ProfileId = ?")
	checkError(err)

	statement.Exec(3)

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
