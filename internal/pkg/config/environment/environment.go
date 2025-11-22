package environment

import (
	"log"
	"syscall"

	"github.com/reoden/go-echo-template/internal/pkg/constants"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Environment string

var (
	Development = Environment(constants.Dev)
	Test        = Environment(constants.Test)
	Production  = Environment(constants.Production)
)

func ConfigEnv(environments ...Environment) Environment {
	environment := Environment("")
	if len(environments) > 0 {
		environment = environments[0]
	} else {
		environment = Development
	}

	// setup viper to read from os environment with `viper.Get`
	viper.AutomaticEnv()

	// https://articles.wesionary.team/environment-variable-configuration-in-your-golang-project-using-viper-4e8289ef664d
	// load environment variables form .env files to system environment variables
	err := godotenv.Load()
	if err != nil {
		log.Printf(".env file cannot be found, err: %v", err)
	}

	setRootWorkingDirectoryEnvironment()

	FixProjectRootWorkingDirectoryPath()

	manualEnv := viper.Get(constants.AppEnv).(string)

	if manualEnv != "" {
		environment = Environment(manualEnv)
	}

	return environment
}

func (env Environment) IsDevelopment() bool {
	return env == Development
}

func (env Environment) IsProduction() bool {
	return env == Production
}

func (env Environment) GetEnvironmentName() string {
	return string(env)
}

func EnvString(key, fallback string) string {
	if value, ok := syscall.Getenv(key); ok {
		return value
	}
	return fallback
}

func setRootWorkingDirectoryEnvironment() {
	absoluteRootWorkingDirectory := GetProjectRootWorkingDirectory()

	// when we `Set` a viper with string value, we should get it from viper with `viper.GetString`, elsewhere we get empty string
	viper.Set(constants.AppRootPath, absoluteRootWorkingDirectory)
}
