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

//before you need set default database
func NewInsert() *DbInsert {
	var i DbInsert

	i.db = defDb

	return &i
}

//set database
func (i *DbInsert) Db(db *sql.DB) *DbInsert {
	i.db = db
	return i
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
		switch val.(type) {
		case string:
			valArr = append(valArr, fmt.Sprintf("'%s'", val))
		default:
			valArr = append(valArr, gStr.FormatAny(val))
		}

	}

	keyStr := strings.Join(keyArr, ",")
	valStr := strings.Join(valArr, ",")
	sqlStr := fmt.Sprintf(
		"insert into %s (%s) values (%s)",
		i.table, keyStr, valStr)

	return sqlStr
}
