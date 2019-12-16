package models

import (
	"github.com/jinzhu/gorm"
)

type Location struct {
	gorm.Model
	Slug      string  `gorm:"unique_index;not null"`
	Title     string  `gorm:"not null"`
	Body      string  `gorm:"size:2048"`
	Lat       float64 `gorm:"type:decimal(10,8)"`
	Lng       float64 `gorm:"type:decimal(11,8)"`
	Comments  []Comment
	Favorites []User `gorm:"many2many:favorites;"`
	Tags      []Tag  `gorm:"many2many:location_tags;association_autocreate:false"`
}

type Comment struct {
	gorm.Model
	Location   Location
	LocationID uint
	User       User
	UserID     uint
	Body       string
}

type Tag struct {
	gorm.Model
	Tag       string     `gorm:"unique_index"`
	Locations []Location `gorm:"many2many:location_tags;"`
}
