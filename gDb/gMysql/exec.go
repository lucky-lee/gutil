package gMysql

import (
	"database/sql"
	"github.com/lucky-lee/gutil/gLog"
)

//exec easy
func execEasy(sqlStr string, db *sql.DB) bool {
	gLog.Sql("execEasySql", sqlStr) //add sql log

	result, err := db.Exec(sqlStr)
	return exec(result, err)
}

//exec easy and return last id
func execEasyLastId(sqlStr string, db *sql.DB) int64 {
	gLog.Sql("execEasyLastIdSql", sqlStr) //add sql log

	result, err := db.Exec(sqlStr)
	return execLastId(result, err)
}

//exec
func exec(result sql.Result, err error) bool {
	if err != nil {
		return false
	}

	rowsAffected, err1 := result.RowsAffected()

	if err1 != nil {
		return false
	}

	if rowsAffected > 0 {
		return true
	}

	return false
}

//exec and return last last id
func execLastId(result sql.Result, err error) int64 {
	if err != nil {
		return 0
	}

	id, err1 := result.LastInsertId()

	if err1 != nil {
		return 0
	}

	return id
}
