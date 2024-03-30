package main

import (
	"fmt"
	"github.com/jsumners/udm-pro-api-client"
	"github.com/jsumners/udm-pro-api-client/cmd/udm-pro-api-client/commands/conf"
	"github.com/jsumners/udm-pro-api-client/cmd/udm-pro-api-client/commands/gethosts"
	"github.com/jsumners/udm-pro-api-client/cmd/udm-pro-api-client/commands/root"
	"github.com/jsumners/udm-pro-api-client/cmd/udm-pro-api-client/internal/app"
	"github.com/jsumners/udm-pro-api-client/cmd/udm-pro-api-client/internal/config"
	"os"
)

var cliApp *app.CliApp
var configFilePath string

func main() {
	cliApp = &app.CliApp{
		Config: config.New(),
	}

	cmd := root.New(&configFilePath, initConfig, initClient)
	cmd.AddCommand(conf.New(cliApp))
	cmd.AddCommand(gethosts.New(cliApp))

	err := cmd.Execute()
	if err != nil {
		fmt.Printf("app error: %v\n", err)
		os.Exit(1)
	}
}

func initConfig() error {
	return cliApp.Config.InitConfig(configFilePath)
}

func initClient() error {
	cfg := cliApp.Config
	client, err := udm.NewWithLogin(udm.HostInfo{
		Address:  cfg.Address,
		Username: cfg.Username,
		Password: cfg.Password,
		Site:     cfg.Site,
	})

	if err != nil {
		return err
	}

	cliApp.Client = client
	return nil
}
