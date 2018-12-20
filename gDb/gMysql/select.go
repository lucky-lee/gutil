package gMysql

import (
	"database/sql"
	"fmt"
	"github.com/lucky-lee/gutil/gFmt"
	"github.com/lucky-lee/gutil/gLog"
	"github.com/lucky-lee/gutil/gStr"
	"reflect"
	"strconv"
	"strings"
)

type DbSelect struct {
	DbBase
	field      string
	join       string
	group      string
	having     string
	order      string
	limit      string
	tableIndex uint8
	bean       interface{}
	fieldArr   []string
	orderMap   map[string]string
	joinMap    map[string]string
}

type BeanSqlJoin struct {
	Table    string
	TableAs  string
	FieldArr []string
}

//type JoinFunc func() (b BeanSqlJoin)

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
	s.tableAs = "t0"
	s.tableIndex = 1
	return s
}

////set table and as name
//func (s *DbSelect) TableAs(t string, a string) *DbSelect {
//	s.table = t
//	s.tableAs = a
//
//	return s
//}

//
func (s *DbSelect) Select(p ...string) *DbSelect {
	if s.fieldArr == nil {
		s.fieldArr = make([]string, 0)
	}

	for _, val := range p {
		s.fieldArr = append(s.fieldArr, fieldName(s.tableAs, val))
	}

	return s
}

//set bean
func (s *DbSelect) Bean(b interface{}) *DbSelect {
	s.bean = b

	t := reflect.TypeOf(s.bean)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag := field.Tag.Get(DB_TAG)
		tagIfnull := field.Tag.Get(DB_TAG_IFNULL)

		if tag == "" {
			continue
		}

		if tag == "-" {
			continue
		}

		if tagIfnull == "true" { //if equal true use mysql ifnull
			tagStr := fmt.Sprintf("ifnull(%s,'') as %s", tag, tag)
			s.fieldArr = append(s.fieldArr, fieldName(s.tableAs, tagStr))
		} else {
			s.fieldArr = append(s.fieldArr, fieldName(s.tableAs, tag))
		}
	}

	return s
}

//join table
func (s *DbSelect) Join(joinTable string, selfOn string, symbol string, otherOn string, joinFields ...string) *DbSelect {

	if len(s.fieldArr) == 0 {
		s.fieldArr = append(s.fieldArr, s.tableAs+".*")
	}

	tAs := fmt.Sprintf("t%d", s.tableIndex)
	s.tableIndex += 1

	for _, val := range joinFields {
		s.fieldArr = append(s.fieldArr, fieldName(tAs, val))
	}

	s.join += gStr.Merge(" join ", joinTable, " as ", tAs, " on ", tAs, ".", selfOn, symbol, s.tableAs, ".", otherOn, " ")
	return s
}

func (s *DbSelect) Where(key string, val interface{}) *DbSelect {
	return s.WhereSymbolQuote(key, "=", val, true)
}

func (s *DbSelect) WhereSymbol(key string, symbol string, val interface{}) *DbSelect {
	s.WhereSymbolQuote(key, symbol, val, true)
	return s
}

func (s *DbSelect) WhereSymbolQuote(key string, symbol string, val interface{}, valQuote bool) *DbSelect {
	if s.whereMap == nil {
		s.whereMap = make(map[string]Where)
	}

	wKey := fieldName(s.tableAs, key)
	s.whereMap[wKey] = NewWhere(wKey, symbol, val, valQuote)

	return s
}
func (s *DbSelect) WhereIn(key string, ins string) *DbSelect {
	insStr := gStr.Merge("(", ins, ")")
	s.WhereSymbolQuote(key, "in", insStr, false)

	return s
}

func (s *DbSelect) GroupBy(column string) *DbSelect {
	s.group = " group by " + column
	return s
}

//order by asc
func (s *DbSelect) OrderAsc(fields ...string) *DbSelect {
	s.initOrderMap()

	for _, val := range fields {
		s.orderMap[val] = "asc"
	}

	return s
}

//order by desc
func (s *DbSelect) OrderDesc(fields ...string) *DbSelect {
	s.initOrderMap()

	for _, val := range fields {
		s.orderMap[val] = "desc"
	}

	return s
}

func (s *DbSelect) LimitPage(page int, pageSize int) *DbSelect {
	if page > 1 {
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
	if s.field == "" {
		if len(s.fieldArr) > 0 {
			s.toFieldArrStr()
		}
	}

	s.toOrderStr()

	if s.field == "" { //if equal empty and set it all
		s.field = "*"
	}

	var sb strings.Builder

	sb.WriteString("select ")
	sb.WriteString(s.field)
	sb.WriteString(" from ")
	sb.WriteString(s.table)

	if s.tableAs != "" {
		sb.WriteString(" as ")
		sb.WriteString(s.tableAs)
	}

	if s.join != "" {
		sb.WriteString(s.join)
	}

	//where
	if len(s.whereMap) > 0 {
		sb.WriteString(" where ")
		sb.WriteString(pubWhereStr(s.whereMap))
	}

	if s.group != "" {
		sb.WriteString(s.group)
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

	gLog.Sql("select", sb.String())

	return sb.String()
}

func (s *DbSelect) ToSqlOne() string {
	s.LimitIndex(0, 1)
	return s.ToSql()
}
func (s *DbSelect) toFieldArrStr() {
	if len(s.fieldArr) == 0 {
		return
	}

	s.field = strings.Join(s.fieldArr, ",")
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

//init function

func (s *DbSelect) initOrderMap() {
	if s.orderMap == nil {
		s.orderMap = make(map[string]string)
	}
}
