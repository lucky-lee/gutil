package gMysql

import (
	"database/sql"
	"strings"
)

type DbDelete struct {
	DbBase
}

//before you need set default database
func NewDelete() *DbDelete {
	var d DbDelete

	d.db = defDb

	return &d
}

//set database
func (d *DbDelete) Db(db *sql.DB) *DbDelete {
	d.db = db
	return d
}

//set table
func (d *DbDelete) Table(s string) *DbDelete {
	d.table = s

	return d
}

func (d *DbDelete) Where(key string, val interface{}) *DbDelete {
	return d.WhereSymbol(key, "=", val)
}

func (d *DbDelete) WhereSymbol(key string, symbol string, val interface{}) *DbDelete {
	if d.whereMap == nil {
		d.whereMap = make(map[string]Where)
	}

	d.whereMap[key] = NewWhere(key, symbol, val, true)

	return d
}

func (d *DbDelete) Do() bool {
	sqlStr := d.ToSql()

	return ExecEasy(sqlStr, d.db)
}

func (d *DbDelete) ToSql() string {
	var sb strings.Builder

	sb.WriteString("delete from ")
	sb.WriteString(d.table)
	sb.WriteString(" where ")
	sb.WriteString(pubWhereStr(d.whereMap))

	return sb.String()
}
