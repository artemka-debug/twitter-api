package env

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	_                = godotenv.Load()
	MysqlUsername, _ = os.LookupEnv("DB_MYSQL_USERNAME")
	MysqlPassword, _ = os.LookupEnv("DB_MYSQL_PASSWORD")
	MysqlHost, _     = os.LookupEnv("DB_MYSQL_HOST")
	MysqlDatabase, _ = os.LookupEnv("DB_MYSQL_DATABASE")
	SmtpHost, _      = os.LookupEnv("SMTP_HOST")
	Email, _         = os.LookupEnv("EMAIL")
	Password, _      = os.LookupEnv("PASSWORD")
)
