package gethosts

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/jsumners/udm-pro-api-client/cmd/udm-pro-api-client/internal/app"
	"github.com/jsumners/udm-pro-api-client/cmd/udm-pro-api-client/internal/config"
	"github.com/jsumners/udm-pro-api-client/cmd/udm-pro-api-client/internal/slug"
	"github.com/jsumners/udm-pro-api-client/pkg/udm"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"text/tabwriter"
)

func New(app *app.CliApp) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-hosts",
		Short: "Generate hosts file for connected devices.",
		Long: heredoc.Doc(`
			Retrieve the list of clients connected to the UDM and generate a hosts
			file representation of them. Any hosts that match the configured host
			aliases will be added to the generated file according to the
			configuration.

			The result will be written to stdout.
		`),
		RunE: func(*cobra.Command, []string) error {
			return run(app.Client, app.Config)
		},
	}

	return cmd
}

func run(client *udm.UdmClient, config *config.Configuration) error {
	foundClients := client.GetConfiguredClients()
	if !config.FixedOnly {
		foundClients = append(foundClients, client.GetActiveClients()...)
	}

	networkClients := reduceNetworkClients(foundClients, config)

	outputHostAliases(config.HostAliases)
	outputRecords(networkClients)

	return nil
}

type hostRecord struct {
	macAddress string
	name       string
	ipAddress  string
}

// reduceNetworkClients formats a clients list into a map indexed by mac address
// and removes any clients with invalid names or ip addresses.
func reduceNetworkClients(clients []udm.NetworkClient, config *config.Configuration) map[string]hostRecord {
	networkClients := make(map[string]hostRecord)
	lowercaseHostnames := config.LowercaseHostnames

	for _, networkClient := range clients {
		name := networkClient.Name
		if name == "" {
			name = networkClient.Hostname
		}
		name = slug.Hostname(name)
		if len(name) < 1 {
			// It is possible that an entry has an empty string set for both
			// "hostname" and "name" in the UDM data. In that case, there's not
			// much we can do. We _could_ make up a name based on the MAC, or other
			// data, but at this time, we just do not care.
			continue
		}
		if lowercaseHostnames {
			name = strings.ToLower(name)
		}

		ip := networkClient.IpAddress
		if len(networkClient.FixedIpAddress) > 0 {
			ip = networkClient.FixedIpAddress
		}
		if len(ip) < 1 {
			continue
		}

		mac := networkClient.MacAddress

		networkClients[mac] = hostRecord{
			macAddress: mac,
			name:       name,
			ipAddress:  ip,
		}
	}

	return networkClients
}

func outputHostAliases(aliases []config.HostAlias) {
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	fmt.Fprint(tw, "# host aliases\n")
	for _, v := range aliases {
		fmt.Fprintf(tw, "%s\t%s", v.IpAddress, v.Name)
	}
	fmt.Fprint(tw, "\n\n")
	tw.Flush()
}

func outputRecords(records map[string]hostRecord) {
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	fmt.Fprint(tw, "# UDM records\n")
	for _, v := range records {
		fmt.Fprintf(tw, "%s\t%s\t# %s\n", v.ipAddress, v.name, v.macAddress)
	}
	tw.Flush()
}
