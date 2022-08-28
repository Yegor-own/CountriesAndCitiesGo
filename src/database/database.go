package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"os"
)

var DB *sqlx.DB

func init() {
	fmt.Println("Reading env file")
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	path := fmt.Sprintf("%v:%v@(%v)/%v", os.Getenv("DB_LOGIN"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	fmt.Println("Connecting to db")
	DB, err = sqlx.Connect("mysql", path)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected successfully")
}
