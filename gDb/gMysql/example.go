package gMysql

import "database/sql"

//example update
func ExampleUpdate() {
	//new function
	//before set default database
	db1 := &sql.DB{}
	db2 := &sql.DB{}

	SetDefDb(db1)
	//exec
	NewUpdate().Table("test").Set("num", 1).Where("id", 1).Do()
	NewUpdate().Table("test").Set("num", 1).Where("id", 1).Incr()
	NewUpdate().Table("test").Set("num", 1).Where("id", 1).Decr()
	NewUpdate().Table("test").Set("num", 1).Where("id", 1).ToSql()
	//exchange database
	NewUpdate().Db(db2).Table("test1").Set("num", 1).Where("id", 1).Do()
	NewUpdate().Db(db2).Table("test1").Set("num", 1).Where("id", 1).Incr()
	NewUpdate().Db(db2).Table("test1").Set("num", 1).Where("id", 1).Decr()
	NewUpdate().Db(db2).Table("test1").Set("num", 1).Where("id", 1).ToSql()

}

//example delete
func ExampleDelete() {
	//new function
	//before set default database
	db1 := &sql.DB{}
	db2 := &sql.DB{}

	SetDefDb(db1)
	//exec
	NewDelete().Table("test").Where("id", 1).Do()
	NewDelete().Table("test").Where("id", 1).ToSql()
	NewDelete().Table("test").WhereSymbol("id", ">=", 1).ToSql()
	//exchange database
	NewDelete().Db(db2).Table("test").Where("id", 1).Do()
	NewDelete().Db(db2).Table("test").Where("id", 1).ToSql()
	NewDelete().Db(db2).Table("test").WhereSymbol("id", ">=", 1).ToSql()
}

//example insert
func ExampleInsert() {
	//new function
	//before set default database
	db1 := &sql.DB{}
	db2 := &sql.DB{}

	SetDefDb(db1)
	//exec
	NewInsert().Table("test").Value("title", "test").Value("title2", "test").Do()
	NewInsert().Table("test").Value("title", "test").Value("title2", "test").DoLastId()
	NewInsert().Table("test").Value("title", "test").Value("title2", "test").ToSql()
	//exchange database
	NewInsert().Db(db2).Table("test").Value("title", "test").Value("title2", "test").Do()
	NewInsert().Db(db2).Table("test").Value("title", "test").Value("title2", "test").DoLastId()
	NewInsert().Db(db2).Table("test").Value("title", "test").Value("title2", "test").ToSql()

}
