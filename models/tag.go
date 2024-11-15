package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// 创建标签之前自动调用
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

// 修改标签之前自动调用
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

// 获取标签列表
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

// 获取标签总数
func GetTagsTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

// 确认标签是否存在
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	return tag.ID > 0
}
func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	return tag.ID > 0
}

// 新建标签
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}

// 修改标签
func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Update(data)

	return true
}

// 删除标签
func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}
