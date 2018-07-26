package gMysql

import (
	"database/sql"
	"fmt"
	"github.com/lucky-lee/gutil/gStr"
)

type DbUpdate struct {
	DbBase
	setMap map[string]interface{}
}

func NewUpdate(db *sql.DB) *DbUpdate {
	var u DbUpdate

	u.db = db

	return &u
}

//set table
func (u *DbUpdate) Table(s string) *DbUpdate {
	u.table = s

	return u
}

//set
func (u *DbUpdate) Set(key string, val interface{}) *DbUpdate {
	if u.setMap == nil {
		u.setMap = make(map[string]interface{})
	}

	u.setMap[key] = val

	return u
}

//set where
func (u *DbUpdate) Where(key string, val interface{}) *DbUpdate {
	return u.WhereSymbol(key, "=", val)
}

//set where by symbol
func (u *DbUpdate) WhereSymbol(key string, symbol string, val interface{}) *DbUpdate {
	if u.whereMap == nil {
		u.whereMap = make(map[string]Where)
	}

	u.whereMap[key] = NewWhere(key,symbol,val)

	return u
}

//update and return bool
func (u *DbUpdate) Do() bool {
	return execEasy(u.ToSql(), u.db)
}

func (u *DbUpdate) ToSql() string {
	setStr := pubSetStr(u.setMap)
	whereStr := pubWhereStr(u.whereMap)

	return fmt.Sprintf(
		`update %s set %s where %s`,
		u.table, setStr, whereStr)
}

//increase
func (u *DbUpdate) Incr() bool {
	for key, val := range u.setMap {
		u.setMap[key] = fmt.Sprintf("%s+%s", key, gStr.FormatAny(val))
	}

	return u.Do()
}

//decrease
func (u *DbUpdate) Decr() bool {
	for key, val := range u.setMap {
		u.setMap[key] = fmt.Sprintf("%s-%s", key, gStr.FormatAny(val))
	}

	return u.Do()
}

//修改-单条(事务)
func UpdateByStrWithTx(tx *sql.Tx, table string, setStr string, whereStr string) bool {
	//if setStr == "" || whereStr == "" {
	//	return false
	//}
	//
	//sqlStr := UpdateSql(table, setStr, whereStr)
	//
	//return UpdateWithTx(tx, sqlStr)
	return false
}

//update with tx
func updateWithTx(tx *sql.Tx, sqlStr string) bool {

	result, err := tx.Exec(sqlStr)

	if err != nil {
		return false
	}

	rAffected, err1 := result.RowsAffected()

	if err1 != nil {
		return false
	}

	if rAffected > 0 { //update success
		return true
	}

	return false
}
