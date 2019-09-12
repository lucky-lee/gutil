package gMysql

import (
	"database/sql"
	"github.com/lucky-lee/gutil/gJson"
	"github.com/lucky-lee/gutil/gLog"
)

//exec easy
func ExecEasy(sqlStr string, db *sql.DB, args ...interface{}) bool {
	gLog.Sql("execEasySql", sqlStr) //add sql log

	return exec(sqlStr, db, args...)
}

//exec easy and return last id
func ExecEasyLastId(sqlStr string, db *sql.DB, args ...interface{}) int64 {
	gLog.Sql("execEasyLastIdSql", sqlStr) //add sql log

	return execLastId(sqlStr, db, args...)
}

//exec
func exec(sqlStr string, db *sql.DB, args ...interface{}) bool {
	result, err := db.Exec(sqlStr, args...)

	if err != nil {
		gLog.E("sqlErr", err)
		gLog.E("sqlErrSql", sqlStr)
		gLog.E("sqlErrParams", gJson.Encode(args))
		return false
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		gLog.E("sqlErr", err)
		gLog.E("sqlErrSql", sqlStr)
		gLog.E("sqlErrParams", gJson.Encode(args))
		return false
	}

	if rowsAffected > 0 {
		return true
	}

	return false
}

//exec and return last last id
func execLastId(sqlStr string, db *sql.DB, args ...interface{}) int64 {
	result, err := db.Exec(sqlStr, args...)

	if err != nil {
		gLog.E("sqlErr", err)
		gLog.E("sqlErrSql", sqlStr)
		gLog.E("sqlErrParams", gJson.Encode(args))
		return 0
	}

	id, err := result.LastInsertId()

	if err != nil {
		gLog.E("sqlErr", err)
		gLog.E("sqlErrSql", sqlStr)
		gLog.E("sqlErrParams", gJson.Encode(args))
		return 0
	}

	return id
}
