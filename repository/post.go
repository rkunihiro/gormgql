package repository

import (
	"time"

	"gorm.io/gorm"

	"github.com/rkunihiro/gormgql/entity"
)

type PostRepository interface {
	FindByID(id int) (*entity.Post, error)
	Find() ([]*entity.Post, error)
	Create(authorId int, title string, posted *time.Time) (*entity.Post, error)
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepo{db}
}

type postRepo struct {
	db *gorm.DB
}

func (p *postRepo) FindByID(id int) (u *entity.Post, err error) {
	u = &entity.Post{}
	err = p.db.Joins("Author").First(u, id).Error
	return u, err
}

func (p *postRepo) Find() (posts []*entity.Post, err error) {
	posts = make([]*entity.Post, 0)
	err = p.db.Joins("Author").Order("id DESC").Find(&posts).Error
	return posts, err
}

func (p *postRepo) Create(authorID int, title string, posted *time.Time) (*entity.Post, error) {
	if posted == nil {
		t := time.Now()
		posted = &t
	}
	post := &entity.Post{
		AuthorID: authorID,
		Title:    title,
		Posted:   *posted,
	}
	result := p.db.Create(post)
	return post, result.Error
}
