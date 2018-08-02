package gMysql

import (
	"testing"
	"github.com/lucky-lee/gutil/gFmt"
	"time"
)

type BeanExample struct {
	UserId     int64  `json:"user_id" field:"user_id"`
	Title      string `json:"title" field:"title"`
	Pic        string `json:"pic" field:"pic"`
	CreateTime int64  `json:"create_time" field:"create_time"`
}

func TestUpdate(t *testing.T) {
	gFmt.Println("update")

	sqlStr := NewUpdate().Table("test").Set("test", 1).Set("title", "title").Where("id", 1).ToSql()
	gFmt.Println(sqlStr)

	t.Log("update success")
}

func TestInsert(t *testing.T) {
	beans := make([]interface{}, 0)
	for i := 1; i < 3; i++ {
		var bean BeanExample

		bean.UserId = int64(i)
		bean.Title = "test insert"
		bean.Pic = "http://www.dripcar.com/test.png"
		bean.CreateTime = time.Now().Unix()

		beans = append(beans, bean)
	}

	NewInsert().Table("sd_news").Beans(beans).ToSql()

	t.Log("insert success")
}

func TestDelete(t *testing.T) {
	gFmt.Println("delete")
	t.Log("delete success")
}

func TestSelect(t *testing.T) {
	gFmt.Println("select")

	sqlStr := NewSelect().Table("test").ToSql()
	gFmt.Println(sqlStr)

	sqlStr1 := NewSelect().Table("test").Select("title", "name", "phone", "gender").WhereSymbol("id", ">", 1).LimitPage(1, 20).ToSql()
	gFmt.Println(sqlStr1)

	sqlStr2 := NewSelect().Table("sd_news").
		Select("title", "pic", "nid").
		Join("user", "user_id", "=", "user_id", "nickname", "phone_num", "photo").
		Where("user_id", 1).
		LimitPage(1, 5).ToSql()
	gFmt.Println(sqlStr2)

	t.Log("select success")
}
