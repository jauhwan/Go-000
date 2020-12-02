package init

import (
	"database/sql"
	"github.com/docker/docker/daemon"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var (
	DB *sql.DB
)

func init() {
	db, err := sql.Open("mysql", "root:admin123@tcp(127.0.0.1:3306)/order")
	if err != nil {
		panic(err)
	}
	DB = db
}
