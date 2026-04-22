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
