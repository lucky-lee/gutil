# gutil
golang util

//example update

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

//example delete

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


//example insert

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


