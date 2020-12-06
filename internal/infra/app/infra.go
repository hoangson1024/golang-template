package app

import (
	"github.com/hoangson1024/golang/internal/infra"

	"github.com/google/wire"
)

var ApplicationSet = wire.NewSet(
	infra.ProvideConfig,
	infra.ProvidePostgreSQL,
	infra.ProvideDBTestRepo,
	infra.ProvideDBTestService,
)
