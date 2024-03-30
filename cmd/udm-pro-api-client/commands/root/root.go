package root

import (
	"github.com/jsumners/udm-pro-api-client/cmd/udm-pro-api-client/internal/app"
	"github.com/spf13/cobra"
)

type InitFn func() error

func New(app *app.CliApp, initFns ...InitFn) *cobra.Command {
	cmd := &cobra.Command{
		Short: "udm-pro-api-client",
		// We don't need Cobra to print out the errors for us.
		// Our app does that by bubbling up errors to the main function.
		SilenceErrors: true,
		PersistentPreRunE: func(*cobra.Command, []string) error {
			for _, initFn := range initFns {
				err := initFn()
				if err != nil {
					return err
				}
			}
			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(
		&app.ConfigFilePath,
		"conf-file",
		"c",
		"",
		"Set the file from which configuration will be loaded.",
	)

	cmd.PersistentFlags().BoolVarP(
		&app.HttpDebugEnabled,
		"http-debug",
		"H",
		false,
		"enable verbose printing of HTTP requests",
	)

	return cmd
}
