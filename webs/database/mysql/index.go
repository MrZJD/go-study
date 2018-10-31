package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// CREATE TABLE `userinfo` (
// 	`uid` INT(10) NOT NULL AUTO_INCREMENT,
// 	`username` VARCHAR(64) NULL DEFAULT NULL,
// 	`departname` VARCHAR(64) NULL DEFAULT NULL,
// 	`created` DATE NULL DEFAULT NULL,
// 	PRIMARY KEY (`uid`)
// );

// CREATE TABLE `userdetail` (
// 	`uid` INT(10) NOT NULL DEFAULT '0',
// 	`intro` TEXT NULL,
// 	`profile` TEXT NULL,
// 	PRIMARY KEY (`uid`)
// )

func main() {
	db, err := sql.Open("mysql", "root:mysql123@/gotest?charset=utf8")
	checkErr(err)
	defer db.Close()

	// insert
	stmt, err := db.Prepare("INSERT userinfo SET username=?, departname=?, created=?")
	checkErr(err)

	res, err := stmt.Exec("肖三金", "WEB", "2018-10-31")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("Inserted: [id] ", id)

	// update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("肖三金2", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// select
	rows, err := db.Query("SELECT * from userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var departname string
		var created string

		err = rows.Scan(&uid, &username, &departname, &created)
		checkErr(err)

		fmt.Println("Rows: [uid username departname created]", uid, username, departname, created)
	}

	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(1)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
