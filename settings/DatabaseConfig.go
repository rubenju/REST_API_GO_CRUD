package settings

import (
	"database/sql"
	"fmt"
	"os"
	"standarisasigoapi/constant"

	_ "github.com/lib/pq"
)

func AccessDatabase() *sql.DB {
	host := os.Getenv(constant.DBHost)
	port := os.Getenv(constant.DBPort)
	user := os.Getenv(constant.DBUser)
	password := os.Getenv(constant.DBPaswword)
	dbname := os.Getenv(constant.DBName)

	psqlcom := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlcom)

	if err != nil {
		panic(err)
	}

	return db
}
