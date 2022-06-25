package entity

import (
	"time"
)

type Post struct {
	ID       int       `json:"id" gorm:"primaryKey"`
	AuthorID int       `json:"authorId" gorm:"column:authorId"`
	Author   *User     `json:"author,omitempty" gorm:"foreignKey:ID;references:AuthorID"`
	Title    string    `json:"title"`
	Posted   time.Time `json:"posted"`
}

func (p Post) TableName() string {
	return "Post"
}
