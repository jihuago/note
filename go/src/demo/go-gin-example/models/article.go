package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

/*
	GORM本身是由回调驱动的，所以我们可以根据需要完全定制GORM，以此达到我们的目的，如下：
		* 注册一个新的回调
		* 删除现有的回调
		* 替换现有的回调
		* 注册回调的顺序
	在GORM中包含已上四类Callbacks
 */

type Article struct {
	Model
	TagId int `json:"tag_id" gorm:"index"` // gorm:index 声明这个字段为索引，如果使用自动迁移功能则有影响，不使用无影响
	Tag Tag `json:"tag"` // Tag 字段实际是一个嵌套的struct，利用TagId与Tag模型相互关联
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

func ExistArticleById(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}

	return false
}

func GetArticleTotal(maps interface{}) (count int) {
	//db.Model(&Article{}).Where(maps).Count(&count)
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)

	return
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Update(data)

	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagId: data["tag_id"].(int),
		Title: data["title"].(string),
		Desc: data["desc"].(string),
		Content: data["content"].(string),
		CreatedBy : data["created_by"].(string),
		State : data["state"].(int),
	})

	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(&Article{})

	return true
}

func CleanAllArticle() bool {
	db.Unscoped().Where("deleted != ?", 0).Delete(&Article{})

	return true
}