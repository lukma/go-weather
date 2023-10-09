package sqlutil

import (
	"database/sql"
	"fmt"

	"altechomega.lukma.dev/notes/config"
	"github.com/go-sql-driver/mysql"
)

func ConnectDB(c *config.Config) *sql.DB {
	config := mysql.Config{
		User:      c.DBConfig.User,
		Passwd:    c.DBConfig.Pass,
		Net:       "tcp",
		Addr:      fmt.Sprintf("%s:%d", c.DBConfig.Host, c.DBConfig.Port),
		DBName:    c.DBConfig.Name,
		ParseTime: true,
	}
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err)
	}

	return db
}
