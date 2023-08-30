package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/go-sql-driver/mysql"
	migrate "github.com/rubenv/sql-migrate"
)

// mysql data migration for cloud sql

var (
	migrations = &migrate.FileMigrationSource{
		Dir: "./migrations",
	}
)

func main() {

	log.Println("start migrate...!")

	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Fatal Error in connect_connector.go: %s environment variable not set.", k)
		}
		return v
	}

	env := mustGetenv("GO_ENV")
	instanceConn := mustGetenv("DB_HOST")
	dbName := mustGetenv("DB_NAME")
	dbUser := mustGetenv("MYSQL_USER_NAME")
	dbPwd := mustGetenv("MYSQL_USER_PASS")

	var dbURI string
	if env == "localhost" || env == "" {
		dbURI = getDBURIForLocalhost(instanceConn, dbUser, dbPwd, dbName)
	} else {
		dbURI = getDBURIForCloudSQL(instanceConn, dbUser, dbPwd, dbName)
	}

	log.Println("dbURI: ", dbURI)

	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatalf("[sql.Open]: %v", err)
	}
	defer dbPool.Close()

	appliedCount, err := migrate.Exec(dbPool, "mysql", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("[migrate.Exec]: %v", err)
		return
	}
	log.Printf("Applied %v migrations", appliedCount)
}

func getDBURIForLocalhost(instanceConn, dbUser, dbPwd, dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true", dbUser, dbPwd, instanceConn, dbName)
}

func getDBURIForCloudSQL(instanceConn, dbUser, dbPwd, dbName string) string {

	d, err := cloudsqlconn.NewDialer(context.Background())
	if err != nil {
		log.Fatalf("[cloudsqlconn.NewDialer]: %v", err)
	}
	var opts []cloudsqlconn.DialOption
	//if usePrivate != "" {
	//	opts = append(opts, cloudsqlconn.WithPrivateIP())
	//}
	mysql.RegisterDialContext("cloudsqlconn",
		func(ctx context.Context, addr string) (net.Conn, error) {
			return d.Dial(ctx, instanceConn, opts...)
		})

	return fmt.Sprintf("%s:%s@cloudsqlconn(localhost:3306)/%s?parseTime=true", dbUser, dbPwd, dbName)
}
