package models

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestGetArticles(t *testing.T) {
	// 创建 sqlmock
	datab, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer datab.Close()

	// 将 sqlmock 与 GORM 连接
	gormDB, err := gorm.Open("mysql", datab)
	assert.NoError(t, err)
	defer gormDB.Close()

	// 设置预期SQL查询和结果
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `blog_article` WHERE (`blog_article`.`state` = ?) LIMIT 10 OFFSET 0")).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "tag_id", "title", "desc", "content", "created_by", "modified_by", "state"}).
			AddRow(1, 1, "Title 1", "Desc 1", "Content 1", "User 1", "User 1", 1).
			AddRow(2, 1, "Title 2", "Desc 2", "Content 2", "User 2", "User 2", 1))

	// 调用被测试函数
	a := NewArticleRepository(gormDB)
	articles := a.GetArticles(0, 10, map[string]interface{}{"state": 1})

	// 断言结果
	assert.Len(t, articles, 2)
	assert.Equal(t, articles[0].Title, "Title 1")
	assert.Equal(t, articles[1].Title, "Title 2")
	assert.NoError(t, mock.ExpectationsWereMet())
}
