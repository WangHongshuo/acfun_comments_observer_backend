package dao

import (
	"fmt"
	"testing"

	"github.com/WangHongshuo/acfun_comments_observer_backend/dao/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_Comment(t *testing.T) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=1 port=5432 dbname=acfun_comm sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	comment1 := &model.Comment{
		Cid:         1,
		FloorNumber: 1,
		Aid:         1,
		Comment:     "TestUser1 Content on Artice 1, Floor 1",
	}
	db.Delete(comment1)

	var result []model.Comment

	db.Where("cid = ?", "1").Find(&result)
	assert.Equal(t, 0, len(result))

	db.Create(comment1)

	db.Where("cid = ?", "1").Find(&result)
	assert.Equal(t, comment1, &result[0])

	db.Delete(comment1)

	db.Where("cid = ?", "1").Find(&result)
	assert.Equal(t, 0, len(result))
}

func Test_Article(t *testing.T) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=1 port=5432 dbname=acfun_comm sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	article1 := &model.Article{
		Aid:             1,
		LastFloorNumber: 9,
	}

	article2 := &model.Article{
		Aid:             2,
		LastFloorNumber: 10,
	}

	db.Delete(article1)
	db.Delete(article2)

	var result []model.Article

	db.Where("aid = ?", "1").Find(&result)
	assert.Equal(t, 0, len(result))

	db.Create(article1)
	db.Save([]model.Article{*article1, *article2})

	db.Where("aid = ?", "2").Find(&result)
	assert.Equal(t, article2, &result[0])

	db.Delete(article1)
	db.Where("aid = ?", "1").Find(&result)
	assert.Equal(t, 0, len(result))

	db.Delete(article2)
	db.Where("aid = ?", "12").Find(&result)
	assert.Equal(t, 0, len(result))
}
