package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tanimutomo/sqlfile"
	"log"
)

// seed を db に流し込む
// seed の sql ファイルは ./build 内に配置している。
// ローカルホストのみ対応する。
func main() {

	instanceConn := "localhost"
	dbName := "webapp_localhost"
	dbUser := "user"
	dbPwd := "pass"

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true", dbUser, dbPwd, instanceConn, dbName)

	// Get a database handler
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize SqlFile
	s := sqlfile.New()

	// Load input file and store queries written in the file
	if err := s.Directory("./build/seeds"); err != nil {
		log.Fatal(err)
	}

	db.Exec("BEGIN;")

	// Execute the stored queries
	// transaction is used to execute queries in Exec()
	res, err := s.Exec(db)
	if err != nil {
		db.Exec("ROLLBACK;")
		log.Fatal(err)
	}

	db.Exec("COMMIT;")

	log.Println(res)
	log.Println("seed done")

}
