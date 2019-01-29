package gMysql

import (
	"database/sql"
	"github.com/lucky-lee/gutil/gLog"
)

//exec easy
func ExecEasy(sqlStr string, db *sql.DB, args ...interface{}) bool {
	gLog.Sql("execEasySql", sqlStr) //add sql log

	result, err := db.Exec(sqlStr, args...)
	return exec(result, err)
}

//exec easy and return last id
func ExecEasyLastId(sqlStr string, db *sql.DB, args ...interface{}) int64 {
	gLog.Sql("execEasyLastIdSql", sqlStr) //add sql log

	result, err := db.Exec(sqlStr, args...)
	return execLastId(result, err)
}

//exec
func exec(result sql.Result, err error) bool {
	if err != nil {
		gLog.E("sqlErr", err)
		return false
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		gLog.E("sqlErr", err)
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
		gLog.E("sqlErr", err)
		return 0
	}

	id, err := result.LastInsertId()

	if err != nil {
		gLog.E("sqlErr", err)
		return 0
	}

	return id
}
