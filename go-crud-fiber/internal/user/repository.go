package user

import "gorm.io/gorm"

type Repository interface {
	Create(user *User) error
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db}
}

func (r *repo) Create(user *User) error {
	return r.db.Create(user).Error
}
