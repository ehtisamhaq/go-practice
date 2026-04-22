package user

import "gorm.io/gorm"

type Repository interface {
	Create(user *User) error
	GetAll() ([]User, error)
	GetByID(id uint) (*User, error)
	Update(user *User) error
	Delete(id uint) error
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

func (r *repo) GetAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repo) GetByID(id uint) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *repo) Update(user *User) error {
	return r.db.Save(user).Error
}

func (r *repo) Delete(id uint) error {
	return r.db.Delete(&User{}, id).Error
}
