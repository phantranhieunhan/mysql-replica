package mysql

import (
	"database/sql"
	"log"
	"mysql-replica/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Database struct {
	ReadDB  *sql.DB
	WriteDB *sql.DB
}

func NewDatabase() Database {
	config.C.ReadURL = "my_db_user:S3cret@tcp(127.0.0.1:3307)/my_db?sslmode=disable"
	rDb, err := sql.Open("mysql", config.C.ReadURL)

	if err != nil {
		log.Fatalln(err)
	}
	config.C.WriteURL = "my_db_user:S3cret@tcp(127.0.0.1:3306)/my_db?sslmode=disable"
	wDb, err := sql.Open("mysql", config.C.WriteURL)
	if err != nil {
		log.Fatalln(err)
	}

	// boil.SetDB(db)
	boil.DebugMode = true

	return Database{
		ReadDB:  rDb,
		WriteDB: wDb,
	}
}
