package app

import (
	"clitemplate/app/bootstrap"
	"clitemplate/app/service/env"

	"go.uber.org/fx"
)

var Container = fx.Module(
	"container",
	bootstrap.Init,
	bootstrap.Loggers,
	fx.Provide(env.Load),
)