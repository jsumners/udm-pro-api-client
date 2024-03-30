package app

import (
	"github.com/jsumners/udm-pro-api-client"
	"github.com/jsumners/udm-pro-api-client/cmd/udm-pro-api-client/internal/config"
)

type CliApp struct {
	ConfigFilePath   string
	Config           *config.Configuration
	Client           *udm.Client
	HttpDebugEnabled bool
}
