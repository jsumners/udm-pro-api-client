package device

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/jsumners/udm-pro-api-client"
	"github.com/jsumners/udm-pro-api-client/cmd/udm-pro-api-client/internal/app"
	"github.com/spf13/cobra"
)

func NewRestart(app *app.CliApp) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "restart macAddress",
		Short: "Restart a device.",
		Long: heredoc.Doc(`
			Issue a restart command to a device that is managed by the UDM.
			For example, use this command to restart a wireless access point that
			is misbehaving.
		`),
		Args:    cobra.ExactArgs(1),
		Example: "device restart 9C:0E:CC:8E:4B:C7",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(app.Client, args[0])
		},
	}

	return cmd
}

func run(client *udm.Client, macAddress string) error {
	return client.RestartDevice(macAddress)
}
