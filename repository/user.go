package repository

import (
	"gorm.io/gorm"

	"github.com/rkunihiro/gormgql/entity"
)

type UserRepository interface {
	FindByID(id int) (*entity.User, error)
	Find() ([]*entity.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

type userRepo struct {
	db *gorm.DB
}

func (p *userRepo) FindByID(id int) (user *entity.User, err error) {
	user = &entity.User{}
	err = p.db.First(user, id).Error
	return user, err
}

func (p *userRepo) Find() (users []*entity.User, err error) {
	users = make([]*entity.User, 0)
	err = p.db.Find(&users).Order("id ASC").Error
	return users, err
}
