package app

import (
	"clitemplate/app/bootstrap/module"
	"clitemplate/app/config"

	"go.uber.org/fx"
)

var Container = fx.Module(
	"container",
	fx.Provide(module.AuditLoggerProvider),
	fx.Provide(module.ErrorLoggerProvider),
	fx.Provide(config.LoadEnv),
)