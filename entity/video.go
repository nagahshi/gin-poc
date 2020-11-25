package entity

import "time"

type Person struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment"`
	FirstName string    `json:"firstname" binding:"required" gorm:"type:varchar(32)"`
	LastName  string    `json:"lastname" binding:"required" gorm:"type:varchar(32)"`
	Age       int8      `json:"age" binding:"gte=1,lte=130"`
	Email     string    `json:"email" validate:"required,email" gorm:"type:varchar(32)"`
	CreatedAt time.Time `json:"-" json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"-" json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type Video struct {
	ID          uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Title       string    `json:"title" binding:"min=2,max=100" gorm:"type:varchar(100)"`
	Description string    `json:"description" binding:"max=200" gorm:"type:varchar(200)"`
	URL         string    `json:"url" binding:"required,url" gorm:"type:varchar(255);UNIQUE"`
	Author      Person    `json:"author" binding:"required" gorm:"foreignkey:PersonID"`
	PersonID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"-" json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"-" json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
