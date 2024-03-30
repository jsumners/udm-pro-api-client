package conf

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/jsumners/udm-pro-api-client/cmd/udm-pro-api-client/internal/app"
	"github.com/jsumners/udm-pro-api-client/cmd/udm-pro-api-client/internal/config"
	"github.com/spf13/cobra"
)

func New(app *app.CliApp) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage configuration.",
	}

	dumpConfigCommand := &cobra.Command{
		Use:   "dump",
		Short: "Write found configuration to stdout.",
		Long: heredoc.Doc(`
			Write the configuration, as the application has read it
			from the configuration file, to stdout.
		`),
		RunE: func(*cobra.Command, []string) error {
			return dumpConfig(app.Config)
		},
	}
	cmd.AddCommand(dumpConfigCommand)

	generateConfigCommand := &cobra.Command{
		Use:   "generate",
		Short: "Write default configuration to stdout.",
		RunE: func(*cobra.Command, []string) error {
			return generateConfig(app.Config)
		},
	}
	cmd.AddCommand(generateConfigCommand)

	return cmd
}

func dumpConfig(cfg *config.Configuration) error {
	yml, err := cfg.GenerateCurrentYaml()
	if err != nil {
		return err
	}
	fmt.Println(yml)
	return nil
}

func generateConfig(cfg *config.Configuration) error {
	yml, err := cfg.GenerateDefaultYaml()
	if err != nil {
		return err
	}
	fmt.Println(yml)
	return nil
}
