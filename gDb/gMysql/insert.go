package gMysql

import (
	"database/sql"
	"fmt"
	"strings"
	"github.com/lucky-lee/gutil/gStr"
)

type DbInsert struct {
	DbBase
	valueMap map[string]interface{}
}

//new insert struct
func NewInsert(db *sql.DB) *DbInsert {
	var d DbInsert

	d.db = db

	return &d
}

//set table
func (i *DbInsert) Table(s string) *DbInsert {
	i.table = s

	return i
}

//set value
func (i *DbInsert) Value(key string, val interface{}) *DbInsert {
	if i.valueMap == nil {
		i.valueMap = make(map[string]interface{})
	}

	i.valueMap[key] = val

	return i
}

//insert and return bool
func (i *DbInsert) Do() bool {
	sqlStr := i.ToSql()
	return execEasy(sqlStr, i.db)
}

//insert and return last id
func (i *DbInsert) DoLastId() int64 {
	sqlStr := i.ToSql()
	return execEasyLastId(sqlStr, i.db)
}

//create insert sql string
func (i *DbInsert) ToSql() string {
	var keyArr []string
	var valArr []string

	for key, val := range i.valueMap {
		keyArr = append(keyArr, key)
		valArr = append(valArr, gStr.FormatAny(val))
	}

	keyStr := strings.Join(keyArr, ",")
	valStr := strings.Join(valArr, ",")
	sqlStr := fmt.Sprintf(
		"insert into %s (%s) values (%s)",
		i.table, keyStr, valStr)

	return sqlStr
}