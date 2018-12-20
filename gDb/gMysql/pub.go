package gMysql

import (
	"fmt"
	"strings"
	"database/sql"
	"github.com/lucky-lee/gutil/gStr"
)

type DbBase struct {
	db       *sql.DB
	table    string
	tableAs  string //as name
	whereMap map[string]Where
}

//type PubFunc interface {
//	ToSql() string
//}
type Set struct {
	Key         string
	Val         interface{}
	ValStrQuote bool
}

type Where struct {
	Key         string
	Symbol      string
	Val         interface{}
	ValStrQuote bool
}

const (
	PAGE_SIZE     = 10
	DB_TAG        = "field"
	DB_TAG_IFNULL = "ifnull"
)

var defDb *sql.DB //默认的数据库

//设置默认数据库
func SetDefDb(db *sql.DB) {
	defDb = db
}
func NewSet(key string, val interface{}, valQuote bool) Set {
	var s Set

	s.Key = key
	s.Val = val
	s.ValStrQuote = valQuote

	return s
}

func NewWhere(key string, symbol string, val interface{}, valQuote bool) Where {
	var w Where

	w.Key = key
	w.Symbol = symbol
	w.Val = val
	w.ValStrQuote = valQuote

	return w
}

//get set string
func pubSetStr(setMap map[string]Set) string {
	if len(setMap) == 0 {
		return ""
	}

	var sets []string

	for _, val := range setMap {
		var str string

		valStr := gStr.FormatAll(val.Val, val.ValStrQuote)

		if valStr != "" && valStr != `""` {
			str = fmt.Sprintf(`%s=%s`, val.Key, valStr)
			sets = append(sets, str)
		}
	}

	return strings.Join(sets, ",")
}

//get where string
func pubWhereStr(whereMap map[string]Where) string {
	if len(whereMap) == 0 {
		return ""
	}

	var wheres []string

	for _, val := range whereMap {
		str := fmt.Sprintf("%s %s %s", val.Key, val.Symbol, gStr.FormatAll(val.Val, val.ValStrQuote))
		wheres = append(wheres, str)
	}

	return strings.Join(wheres, " and ")
}

func fieldName(asName string, fieldName string) string {
	if asName == "" {
		return fieldName
	} else {
		if strings.Contains(fieldName, "ifnull(") {
			repStr := fmt.Sprintf("ifnull(%s.", asName)
			str := strings.Replace(fieldName, "ifnull(", repStr, -1)
			return str
		} else {
			return gStr.Merge(asName, ".", fieldName)
		}
	}
}
