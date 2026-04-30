package bootstrap

import (
	"clitemplate/app/service/logger"

	"go.uber.org/fx"
)

var Loggers = fx.Options(
	fx.Invoke(initLoggers),
	fx.Provide(logger.NewAuditLogger),
	fx.Provide(logger.NewErrorLogger),
)

func initLoggers(
	auditLogger *logger.AuditLogger,
	errorLogger *logger.ErrorLogger,
) error {
	if err := auditLogger.Init(); err != nil {
		return err
	}
	if err := errorLogger.Init(); err != nil {
		return err
	}
	return nil
}