// internal/user/service.go
package user

type UserService struct {
	repo Repository
}

func NewUserService(r Repository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Register(user *User) error {
	// Add business logic here (e.g., hashing passwords)
	return s.repo.Create(user)
}

func (s *UserService) GetAll() ([]User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetByID(id uint) (*User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) Update(user *User) error {
	return s.repo.Update(user)
}

func (s *UserService) Delete(id uint) error {
	return s.repo.Delete(id)
}
