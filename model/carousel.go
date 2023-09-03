package model

import "gorm.io/gorm"

// 轮播图模型
type Carousel struct {
	gorm.Model
	ImgPath   string
	ProductID uint `gorm:"not null"`
}
