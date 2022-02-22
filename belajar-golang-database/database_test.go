package belajar_golang_database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {

}

func TestConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:secret@tcp(localhost:3306)/belajar_database_go")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
