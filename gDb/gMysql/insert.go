package gMysql

import (
	"database/sql"
	"fmt"
	"strings"
	"github.com/lucky-lee/gutil/gStr"
	"reflect"
	"github.com/lucky-lee/gutil/gFmt"
)

type DbInsert struct {
	DbBase
	keyArr   []string
	valueMap []map[string]interface{}
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

func (i *DbInsert) Bean(b interface{}) *DbInsert {
	i.appendValueMap(b)
	return i
}

func (i *DbInsert) Beans(bs []interface{}) *DbInsert {
	for _, val := range bs {
		i.appendValueMap(val)
	}
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
	var valArr []string

	for _, valMap := range i.valueMap {
		var valueArr []string

		for index := 0; index < len(i.keyArr); index++ {
			value := valMap[i.keyArr[index]]
			switch value.(type) {
			case string:
				valueArr = append(valueArr, fmt.Sprintf("'%s'", value))
			default:
				valueArr = append(valueArr, gStr.FormatAny(value))
			}
		}

		valArr = append(valArr, gStr.Merge("(", strings.Join(valueArr, ","), ")"))
	}

	keyStr := strings.Join(i.keyArr, ",")
	valStr := strings.Join(valArr, ",")
	sqlStr := fmt.Sprintf(
		"insert into %s (%s) values %s",
		i.table, keyStr, valStr)

	gFmt.Println(sqlStr)

	return sqlStr
}

func (i *DbInsert) appendValueMap(bean interface{}) {

	values := make(map[string]interface{})
	value := reflect.ValueOf(bean)

	for index := 0; index < value.NumField(); index++ {
		t := value.Type().Field(index)
		v := value.Field(index)

		values[t.Tag.Get(DB_TAG)] = v.Interface()
		//str := fmt.Sprintf("name:%v,type:%v,fieldName:%v,value:%v", t.Name, t.Type, t.Tag.Get(DB_TAG), v)
		//gFmt.Println(str)

		if len(i.valueMap) == 0 {
			i.keyArr = append(i.keyArr, t.Tag.Get(DB_TAG))
		}
	}

	i.valueMap = append(i.valueMap, values)
}
