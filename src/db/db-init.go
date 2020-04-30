package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

var (
	_ = godotenv.Load()
	MysqlUsername, _ = os.LookupEnv("DB_MYSQL_USERNAME")
	MysqlPassword, _ = os.LookupEnv("DB_MYSQL_PASSWORD")
	MysqlHost, _     = os.LookupEnv("DB_MYSQL_HOST")
	MysqlDatabase, _ = os.LookupEnv("DB_MYSQL_DATABASE")
	DB, _            = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", MysqlUsername, MysqlPassword, MysqlHost, MysqlDatabase))
)
