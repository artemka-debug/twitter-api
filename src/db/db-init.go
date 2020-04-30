package db

import (
	"database/sql"
	"fmt"
	"github.com/artemka-debug/twitter-api/src/env"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB, _            = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", env.MysqlUsername, env.MysqlPassword, env.MysqlHost, env.MysqlDatabase))
)
