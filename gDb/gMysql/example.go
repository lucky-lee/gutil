package gMysql

import "database/sql"

//example update
func ExampleUpdate() {

	//create sql string
	NewUpdate(&sql.DB{}).Table("test").Set("test", 1).Where("id", 1).ToSql()

	//update test set test = "1" where id=1
	NewUpdate(&sql.DB{}).Table("test").Set("test", "1").Where("id", 1).Do()

	//update test set test="1" where id>=1
	NewUpdate(&sql.DB{}).Table("test").Set("test", "1").WhereSymbol("id", ">=", 1).Do()

	//update test set num increase one
	NewUpdate(&sql.DB{}).Table("test").Set("num", 1).Where("id", 1).Incr()

	//update test set num decrease one
	NewUpdate(&sql.DB{}).Table("test").Set("num", 1).Where("id", 1).Decr()

	//new function
	//before set default database
	db1 := &sql.DB{}
	db2 := &sql.DB{}

	SetDefDb(db1)
	//exec
	NewUpdates().Table("test").Set("num", 1).Where("id", 1).Do()
	NewUpdates().Table("test").Set("num", 1).Where("id", 1).Incr()
	NewUpdates().Table("test").Set("num", 1).Where("id", 1).Decr()
	NewUpdates().Table("test").Set("num", 1).Where("id", 1).ToSql()
	//exchange database
	NewUpdates().Db(db2).Table("test1").Set("num", 1).Where("id", 1).Do()
	NewUpdates().Db(db2).Table("test1").Set("num", 1).Where("id", 1).Incr()
	NewUpdates().Db(db2).Table("test1").Set("num", 1).Where("id", 1).Decr()
	NewUpdates().Db(db2).Table("test1").Set("num", 1).Where("id", 1).ToSql()

}

//example delete
func ExampleDelete() {

	//create sql string
	NewDelete(&sql.DB{}).Table("test").Where("id", 1).ToSql()

	//delete from test where id=1
	NewDelete(&sql.DB{}).Table("test").Where("id", 1).Do()

	//delete from test where id>=1
	NewDelete(&sql.DB{}).Table("test").WhereSymbol("id", ">=", 1).Do()

	//new function
	//before set default database
	db1 := &sql.DB{}
	db2 := &sql.DB{}

	SetDefDb(db1)
	//exec
	NewDeletes().Table("test").Where("id", 1).Do()
	NewDeletes().Table("test").Where("id", 1).ToSql()
	NewDeletes().Table("test").WhereSymbol("id", ">=", 1).ToSql()
	//exchange database
	NewDeletes().Db(db2).Table("test").Where("id", 1).Do()
	NewDeletes().Db(db2).Table("test").Where("id", 1).ToSql()
	NewDeletes().Db(db2).Table("test").WhereSymbol("id", ">=", 1).ToSql()
}

//example insert
func ExampleInsert() {

	//create sql string
	NewInsert(&sql.DB{}).Table("test").Value("test", 1).Value("test2", 2).ToSql()

	//insert into test (test,test2) values(1,2) and return bool
	NewInsert(&sql.DB{}).Table("test").Value("test", 1).Value("test2", 2).Do()

	//insert into test (test,test2) values(1,2) and return last id
	NewInsert(&sql.DB{}).Table("test").Value("test", 1).Value("test2", 2).DoLastId()

	//new function
	//before set default database
	db1 := &sql.DB{}
	db2 := &sql.DB{}

	SetDefDb(db1)
	//exec
	NewInserts().Table("test").Value("title", "test").Value("title2", "test").Do()
	NewInserts().Table("test").Value("title", "test").Value("title2", "test").DoLastId()
	NewInserts().Table("test").Value("title", "test").Value("title2", "test").ToSql()
	//exchange database
	NewInserts().Db(db2).Table("test").Value("title", "test").Value("title2", "test").Do()
	NewInserts().Db(db2).Table("test").Value("title", "test").Value("title2", "test").DoLastId()
	NewInserts().Db(db2).Table("test").Value("title", "test").Value("title2", "test").ToSql()

}
