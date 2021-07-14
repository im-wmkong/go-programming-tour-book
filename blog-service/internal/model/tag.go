package model

import (
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"gorm.io/gorm"
)

type Tag struct {
	Model `gorm:"embedded"`
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

// func (t Tag) TableName() string {
// 	return "blog_tag"
// }

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	if err := db.Model(&t).Where("state = ?", t.State).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	if err = db.Where("state = ?", t.State).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	global.Logger.Debug(nil, db)
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB) error {
	return db.Save(&t).Error
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Delete(&t).Error
}
