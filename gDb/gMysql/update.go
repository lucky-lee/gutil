package gMysql

import (
	"database/sql"
	"fmt"
	"reflect"
)

type DbUpdate struct {
	DbBase
	setMap map[string]Set
}

//before you need set default database
func NewUpdate() *DbUpdate {
	var u DbUpdate

	u.db = defDb

	return &u
}

//set database
func (u *DbUpdate) Db(db *sql.DB) *DbUpdate {
	u.db = db
	return u
}

//set table
func (u *DbUpdate) Table(s string) *DbUpdate {
	u.table = s

	return u
}

func (u *DbUpdate) Bean(b interface{}) *DbUpdate {
	value := reflect.ValueOf(b)

	//if value is ptr
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	for index := 0; index < value.NumField(); index++ {
		t := value.Type().Field(index)
		v := value.Field(index)

		if t.Tag.Get(DB_TAG) != "" {
			u.Set(t.Tag.Get(DB_TAG), v.Interface())
		}
	}

	return u
}

//set
func (u *DbUpdate) Set(key string, val interface{}) *DbUpdate {
	if u.setMap == nil {
		u.setMap = make(map[string]Set)
	}

	u.setMap[key] = NewSet(key, val, true)
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

	u.whereMap[key] = NewWhere(key, symbol, val, true)

	return u
}

//update and return bool
func (u *DbUpdate) Do() bool {
	return ExecEasy(u.ToSql(), u.db)
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
	u.incr()
	return u.Do()
}

//increase sql string
func (u *DbUpdate) IncrSql() string {
	u.incr()
	return u.ToSql()
}

//decrease
func (u *DbUpdate) Decr() bool {
	u.decr()
	return u.Do()
}

//decrease sql string
func (u *DbUpdate) DecrSql() string {
	u.decr()
	return u.ToSql()
}

//increase
func (u *DbUpdate) incr() {
	for key, val := range u.setMap {
		v := fmt.Sprintf("%s+%v", key, val.Val)
		u.setMap[key] = NewSet(key, v, false)
	}
}

//decrease
func (u *DbUpdate) decr() {
	for key, val := range u.setMap {
		v := fmt.Sprintf("%s-%v", key, val.Val)
		u.setMap[key] = NewSet(key, v, false)
	}
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
