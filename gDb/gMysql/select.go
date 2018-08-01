package gMysql

import (
	"strings"
	"reflect"
	"github.com/lucky-lee/gutil/gLog"
	"fmt"
	"github.com/lucky-lee/gutil/gFmt"
	"strconv"
	"database/sql"
	"github.com/lucky-lee/gutil/gStr"
)

type DbSelect struct {
	DbBase
	field    string
	join     string
	group    string
	having   string
	order    string
	limit    string
	tags     []string
	bean     interface{}
	NameMap  map[string]interface{}
	Values   []interface{}
	orderMap map[string]string
}

func NewSelect() *DbSelect {
	var s DbSelect

	s.db = defDb

	return &s
}

//set database
func (s *DbSelect) Db(db *sql.DB) *DbSelect {
	s.db = db
	return s
}

//set table
func (s *DbSelect) Table(t string) *DbSelect {
	s.table = t
	return s
}

//set fields
func (s *DbSelect) Fields(f string) *DbSelect {
	s.field = f
	return s
}

//set bean
func (s *DbSelect) Bean(b interface{}) *DbSelect {
	s.bean = b
	return s
}

func (s *DbSelect) Where(key string, val interface{}) *DbSelect {
	return s.WhereSymbol(key, "=", val)
}

func (s *DbSelect) WhereSymbol(key string, symbol string, val interface{}) *DbSelect {
	if s.whereMap == nil {
		s.whereMap = make(map[string]Where)
	}

	s.whereMap[key] = NewWhere(key, symbol, val)

	return s
}

//order by asc
func (s *DbSelect) OrderAsc(fields ...string) {
	if s.orderMap == nil {
		s.orderMap = make(map[string]string)
	}

	for _, val := range fields {
		s.orderMap[val] = "asc"
	}
}

//order by desc
func (s *DbSelect) OrderDesc(fields ...string) {
	if s.orderMap == nil {
		s.orderMap = make(map[string]string)
	}

	for _, val := range fields {
		s.orderMap[val] = "desc"
	}
}

func (s *DbSelect) LimitPage(page int, pageSize int) *DbSelect {
	if page > 0 {
		pageIndex := (page - 1) * pageSize
		s.LimitIndex(pageIndex, pageSize)
	}

	return s
}

func (s *DbSelect) LimitIndex(pageIndex int, pageSize int) *DbSelect {
	s.limit = fmt.Sprintf("limit %d,%d", pageIndex, pageSize)

	return s
}

func (s *DbSelect) LimitPageDef(page int) *DbSelect {
	if page > 0 {
		pageIndex := (page - 1) * PAGE_SIZE
		s.LimitIndexDef(pageIndex)
	}

	return s
}

func (s *DbSelect) LimitIndexDef(pageIndex int) *DbSelect {
	s.LimitIndex(pageIndex, PAGE_SIZE)

	return s
}

func (s *DbSelect) do() {
	sqlStr := s.ToSql()

	rows, err := s.db.Query(sqlStr)

	if err != nil {
		gLog.E("select.err", err)
		gLog.E("select.sql", sqlStr)
		return
	}

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	//t := reflect.TypeOf(s.bean)
	//tVal := reflect.ValueOf(s.bean)

	//var beans []interface{}

	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)

		v := reflect.New(reflect.TypeOf(s.bean)).Elem().Elem()
		record := make(map[string]string)

		for i, col := range values {
			strVal := string(col.([]byte))
			f := v.FieldByName(columns[i])

			if f.Kind() == reflect.Int64 {
				intVal, _ := strconv.ParseInt(strVal, 10, 64)
				f.SetInt(intVal)
				//	gFmt.Println(columns[i])
				//	gFmt.Println(tVal.Field(i))
				//	gFmt.Println(t.Field(i).Name)
				//	gFmt.Println(t.Field(i).Tag.Get("name"))
			}
			//if tVal.Field(i).Kind() == reflect.Int64 {
			//	intVal, _ := strconv.ParseInt(strVal, 10, 64)
			//	tVal.Field(i).SetInt(intVal)
			//	gFmt.Println(columns[i])
			//	gFmt.Println(tVal.Field(i))
			//	gFmt.Println(t.Field(i).Name)
			//	gFmt.Println(t.Field(i).Tag.Get("name"))
			//}

			gFmt.Println(s.bean)
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
}

//create sql string
func (s *DbSelect) ToSql() string {
	s.toFieldStr()
	s.toOrderStr()

	var sb strings.Builder

	sb.WriteString("select ")
	sb.WriteString(s.field)
	sb.WriteString(" from ")
	sb.WriteString(s.table)

	//where
	if len(s.whereMap) > 0 {
		sb.WriteString(" where ")
		sb.WriteString(pubWhereStr(s.whereMap))
	}

	//order by
	if s.order != "" {
		sb.WriteString(s.order)
	}

	//limit
	if s.limit != "" {
		sb.WriteString(" ")
		sb.WriteString(s.limit)
	}

	return sb.String()
}

//生成field
func (s *DbSelect) toFieldStr() {
	if s.NameMap == nil {
		s.NameMap = make(map[string]interface{})
	}

	t := reflect.TypeOf(s.bean)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag := field.Tag.Get("name")

		if tag == "" {
			continue
		}

		if tag == "-" {
			continue
		}

		s.tags = append(s.tags, tag)
		s.NameMap[tag] = field.Name
	}

	if s.field == "" {
		s.field = strings.Join(s.tags, ",")
	}
}

//create order by string
func (s *DbSelect) toOrderStr() {
	if len(s.orderMap) == 0 {
		return
	}

	var strs []string

	for key, val := range s.orderMap {
		str := gStr.Merge(key, " ", val)
		strs = append(strs, str)
	}

	joinStr := strings.Join(strs, ",")

	if joinStr != "" {
		s.order = " order by " + joinStr
	}
}
