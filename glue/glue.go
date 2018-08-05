package glue

import (
	"github.com/pankona/di-sandbox/repository"
	"github.com/pankona/di-sandbox/service"
)

type toServiceRepoer struct {
	repo *repository.Repo
}

func (s *toServiceRepoer) CreateNewUser() error {
	return s.repo.Create()
}

func ServiceRepoer(r *repository.Repo) service.Repoer {
	return &toServiceRepoer{repo: r}
}
