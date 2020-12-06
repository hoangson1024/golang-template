package infra

import (
	db_test_repo "github.com/hoangson1024/golang/pkg/test-db/repo"
	db_test_service "github.com/hoangson1024/golang/pkg/test-db/service"

	"github.com/go-pg/pg"
)

func ProvideDBTestRepo(db *pg.DB) db_test_repo.WrapperRepo {
	return db_test_repo.CreateWrapperRepo(db)
}

func ProvideDBTestService(repo db_test_repo.WrapperRepo) db_test_service.Wrapper {
	return db_test_service.CreateWrapperService(repo)
}
