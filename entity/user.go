package entity

type User struct {
	ID    int     `json:"id" gorm:"primaryKey"`
	Name  string  `json:"name"`
	Posts *[]Post `json:"posts,omitempty" gorm:"foreignKey:AuthorID;references:ID"`
}

func (p User) TableName() string {
	return "User"
}
