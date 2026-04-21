package user

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) Create(user User) error {
	// Business Logic: e.g., check if email is valid
	return s.repo.Create(user)
}

func (s *service) GetAll() ([]User, error) {
	return s.repo.GetAll()
}

func (s *service) GetByID(id int) (User, error) {
	return s.repo.GetByID(id)
}

func (s *service) Update(user User) error {
	return s.repo.Update(user)
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}
