package bootstrap

import (
	"clitemplate/app/service/env"
	"fmt"
	"os"

	"go.uber.org/fx"
)

var Init = fx.Options(
	fx.Invoke(InitWorkingDirs),
	fx.Invoke(ValidateEnv),
)

func InitWorkingDirs(env *env.Env) error {
	dirs := []string{
		env.GetLogsDir(),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("cannot create dir %s", dir)
		}
	}

	return nil
}

func ValidateEnv(env *env.Env) error {
	return env.Validate()
}
