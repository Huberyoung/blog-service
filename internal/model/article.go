package model

import (
	"github.com/jinzhu/gorm"
	"service/pkg/app"
)

type ArticleSwagger struct {
	List  []*Article
	Paper *app.Pager
}

func (a Article) Count(db *gorm.DB) (int, error) {
	var count int
	if a.Title != "" {
		db = db.Where("name = ?", a.Title)
	}
	db = db.Where("state = ?", a.State)

	if err := db.Model(&a).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (a Article) Get(db *gorm.DB, id uint32) (Article, error) {
	var article Article
	if err := db.Where("id = ?", id).Select(&article).Error; err != nil {
		return article, err
	}
	return article, nil
}

func (a Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var articles []*Article
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if a.Title != "" {
		db = db.Where("Title = ? ", a.Title)
	}
	db = db.Where("state = ?", a.State)

	if err = db.Where("is_del = ?", 0).Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Article) Update(db *gorm.DB, values any) error {
	return db.Model(&Article{}).Where("id = ? AND is_del = ?", a.ID, 0).Update(values).Error
}

func (a Article) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", a.ID, 0).Delete(&a).Error
}
