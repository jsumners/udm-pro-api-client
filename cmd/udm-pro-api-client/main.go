package main

import (
	"fmt"
	"github.com/jsumners/udm-pro-api-client"
	"github.com/jsumners/udm-pro-api-client/cmd/udm-pro-api-client/commands/conf"
	"github.com/jsumners/udm-pro-api-client/cmd/udm-pro-api-client/commands/device"
	"github.com/jsumners/udm-pro-api-client/cmd/udm-pro-api-client/commands/gethosts"
	"github.com/jsumners/udm-pro-api-client/cmd/udm-pro-api-client/commands/root"
	"github.com/jsumners/udm-pro-api-client/cmd/udm-pro-api-client/internal/app"
	"github.com/jsumners/udm-pro-api-client/cmd/udm-pro-api-client/internal/config"
	"os"
)

var cliApp *app.CliApp

func main() {
	cliApp = &app.CliApp{
		Config: config.New(),
	}

	cmd := root.New(cliApp, initConfig, initClient)
	cmd.AddCommand(conf.New(cliApp))
	cmd.AddCommand(gethosts.New(cliApp))
	cmd.AddCommand(device.New(cliApp))

	err := cmd.Execute()
	if err != nil {
		fmt.Printf("app error: %v\n", err)
		os.Exit(1)
	}
}

func initConfig() error {
	return cliApp.Config.InitConfig(cliApp.ConfigFilePath)
}

func initClient() error {
	cfg := cliApp.Config
	options := make([]udm.Option, 0)

	if cliApp.HttpDebugEnabled == true {
		options = append(options, udm.WithHttpDebug())
	}

	hostInfo := udm.HostInfo{
		Address:  cfg.Address,
		Username: cfg.Username,
		Password: cfg.Password,
		Site:     cfg.Site,
	}
	client, err := udm.NewWithLogin(hostInfo, options...)

	if err != nil {
		return err
	}

	cliApp.Client = client
	return nil
}
