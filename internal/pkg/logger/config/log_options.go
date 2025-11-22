package config

import (
	"github.com/reoden/go-echo-template/internal/pkg/config"
	"github.com/reoden/go-echo-template/internal/pkg/config/environment"
	"github.com/reoden/go-echo-template/internal/pkg/logger/models"
	typeMapper "github.com/reoden/go-echo-template/internal/pkg/reflection/typemapper"

	"github.com/iancoleman/strcase"
)

var optionName = strcase.ToLowerCamel(typeMapper.GetGenericTypeNameByT[LogOptions]())

type LogOptions struct {
	LogLevel      string         `mapstructure:"level"`
	LogType       models.LogType `mapstructure:"logType"`
	CallerEnabled bool           `mapstructure:"callerEnabled"`
	EnableTracing bool           `mapstructure:"enableTracing" default:"true"`
}

func ProvideLogConfig(env environment.Environment) (*LogOptions, error) {
	return config.BindConfigKey[*LogOptions](optionName, env)
}
