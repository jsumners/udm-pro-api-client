package device

import (
	"github.com/jsumners/udm-pro-api-client/cmd/udm-pro-api-client/internal/app"
	"github.com/spf13/cobra"
)

func New(app *app.CliApp) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "device",
		Short: "Interact with devices known by the UDM.",
	}

	cmd.AddCommand(NewRestart(app))

	return cmd
}
