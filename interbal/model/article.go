package model

import "github.com/jinzhu/gorm"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	CoverImageUrl string `json:"cover_image_url"`
	Content       string `json:"content"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Get(db *gorm.DB) (Article, error) {
	if err := db.First(&a).Error; err != nil {
		return Article{}, err
	}
	return a, nil
}

func (a Article) Count(db *gorm.DB) (int, error) {
	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}
	db = db.Where("state = ?", a.State)

	var count int
	if err := db.Model(&a).Where("id_del = ?", 0).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func (a Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var articles []*Article
	var err error
	if pageOffset > 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}
	db = db.Where("state = ?", a.State)
	err = db.Model(&a).Where("id_del = ?", 0).Find(&articles).Error
	return articles, err
}

func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Article) Update(db *gorm.DB, values any) error {
	return db.Model(&Article{}).Where("id = ? and is_del = ?", a.Model.ID, 0).Updates(values).Error
}

func (a Article) Delete(db *gorm.DB) error {
	return db.Where("id = ? and is_del = ?", a.Model.ID, 0).Delete(&a).Error
}
