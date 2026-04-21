package user

import (
	"errors"
	"fmt"
)

type Repository interface {
	Create(user User) error
	GetAll() ([]User, error)
	GetByID(id int) (User, error)
	Update(user User) error
	Delete(id int) error
}

type InMemoryRepository struct {
	users  []User
	nextID int
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		users:  []User{},
		nextID: 1,
	}
}

func (r *InMemoryRepository) Create(user User) error {
	user.ID = r.nextID
	r.nextID++
	r.users = append(r.users, user)
	return nil
}

func (r *InMemoryRepository) GetAll() ([]User, error) {
	return r.users, nil
}

func (r *InMemoryRepository) GetByID(id int) (User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}

func (r *InMemoryRepository) Update(user User) error {
	for i, u := range r.users {
		if u.ID == user.ID {
			r.users[i] = user
			return nil
		}
	}
	return fmt.Errorf("user not found")
}

func (r *InMemoryRepository) Delete(id int) error {
	for i, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user not found")
}
