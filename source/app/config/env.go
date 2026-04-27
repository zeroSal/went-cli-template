package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Env    string
	VarDir string
}

func LoadEnv() *Env {
	_ = godotenv.Load()

	env := &Env{
		Env:    "dev",
		VarDir: "var",
	}

	if environment := os.Getenv("ENV"); environment != "" {
		env.Env = environment
	}

	if varDir := os.Getenv("VAR_DIR"); varDir != "" {
		env.VarDir = varDir
	}

	return env
}

func (e *Env) Validate() error {
	if (e.Env != "dev" && e.Env != "prod") {
		return errors.New("invalid env provided (it must be 'dev' or 'prod')")
	}

	return nil
}

func (e *Env) GetLogsDir() string {
	return fmt.Sprintf("%s/logs", e.VarDir)
}
