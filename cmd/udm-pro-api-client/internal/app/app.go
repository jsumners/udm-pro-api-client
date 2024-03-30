package app

import (
	"github.com/jsumners/udm-pro-api-client/cmd/udm-pro-api-client/internal/config"
	"github.com/jsumners/udm-pro-api-client/pkg/udm"
)

type CliApp struct {
	Config *config.Configuration
	Client *udm.UdmClient
}
