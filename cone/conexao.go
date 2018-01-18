package cone

import (
	"github.com/jmoiron/sqlx"
	/**
	github.com/go-sql-driver/mysql not is used in apllication directamente
	*/
	_ "github.com/go-sql-driver/mysql"
)

//Db recebe um ponteiro de sqlx.DB
var Db *sqlx.DB

//Connection abre uma conexão com banco de dados
func Connection() (err error) {
	err = nil

	Db, err = sqlx.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/cms")
	if err != nil {
		return
	}
	err = Db.Ping()
	if err != nil {
		return
	}
	return
}
