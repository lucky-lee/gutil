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
	whereMap map[string]Where
}

//type PubFunc interface {
//	ToSql() string
//}

type Where struct {
	Key    string
	Symbol string
	Val    interface{}
}

func NewWhere(key string, symbol string, val interface{}) Where {
	var w Where

	w.Key = key
	w.Symbol = symbol
	w.Val = val

	return w
}

//get set string
func pubSetStr(setMap map[string]interface{}) string {
	if len(setMap) == 0 {
		return ""
	}

	var sets []string

	for key, val := range setMap {
		str := fmt.Sprintf("%s=%s", key, gStr.FormatAny(val))
		sets = append(sets, str)
	}

	return strings.Join(sets, ",")
}

//get where string
func pubWhereStr(whereMap map[string]Where) string {
	if len(whereMap) == 0 {
		return ""
	}

	var wheres []string

	for key, val := range whereMap {
		str := fmt.Sprintf("%s %s %s", key, val.Symbol, gStr.FormatAny(val.Val))
		wheres = append(wheres, str)
	}

	return strings.Join(wheres, " and ")
}