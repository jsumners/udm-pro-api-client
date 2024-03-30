package root

import (
	"github.com/spf13/cobra"
)

type InitFn func() error

func New(configFile *string, initFns ...InitFn) *cobra.Command {
	cmd := &cobra.Command{
		Short: "udm-pro-api-client",
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
		configFile,
		"conf-file",
		"c",
		"",
		"Set the file from which configuration will be loaded.",
	)

	return cmd
}
