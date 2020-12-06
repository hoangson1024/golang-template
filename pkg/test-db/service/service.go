package service

import (
	"context"

	repo_pkg "github.com/hoangson1024/golang/pkg/test-db/repo"
)

type wrapperService struct {
	repo repo_pkg.WrapperRepo
}

func CreateWrapperService(repo repo_pkg.WrapperRepo) *wrapperService {
	return &wrapperService{repo: repo}
}

type Wrapper interface {
	Load(ctx context.Context) ([]repo_pkg.Technology, error)
}

func (l *wrapperService) Load(ctx context.Context) ([]repo_pkg.Technology, error) {
	return l.repo.Load(ctx)
}
