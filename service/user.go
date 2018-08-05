package service

type Repoer interface {
	CreateNewUser() error
}

type User struct {
	repo Repoer
}

func New(r Repoer) *User {
	return &User{repo: r}
}

func (u *User) CreateNewUser() error {
	return u.repo.CreateNewUser()
}
