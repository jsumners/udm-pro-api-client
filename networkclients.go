package udm

import (
	"encoding/json"
	"errors"
	"fmt"
)

type NetworkClient struct {
	Hostname       string `json:"hostname"`
	FixedIpAddress string `json:"fixed_ip"`
	IpAddress      string `json:"ip"`
	MacAddress     string `json:"mac"`
	Name           string `json:"name"`
}

// NetworkClientsResponse represents the response sent when querying the
// UDM API for a list of network clients.
type NetworkClientsResponse struct {
	Meta ResponseMeta    `json:"meta"`
	Data []NetworkClient `json:"data"`
}

// GetActiveClients retrieves the list of currently connected clients.
func (client *Client) GetActiveClients() ([]NetworkClient, error) {
	return client.getNetworkClients(
		fmt.Sprintf("/proxy/network/api/s/%s/stat/sta", client.hostInfo.Site),
	)
}

// GetConfiguredClients retrieves the list of statically configured clients.
func (client *Client) GetConfiguredClients() ([]NetworkClient, error) {
	return client.getNetworkClients(
		fmt.Sprintf("/proxy/network/api/s/%s/list/user", client.hostInfo.Site),
	)
}

func (client *Client) getNetworkClients(path string) ([]NetworkClient, error) {
	resp, err := client.resty.R().Get(path)
	if err != nil {
		return nil, fmt.Errorf("failed to query for network clients: %w", err)
	}

	if resp.StatusCode() < 200 || resp.StatusCode() > 299 {
		return nil, errors.New(fmt.Sprintf("http error: %s", resp.Status()))
	}

	parsed := NetworkClientsResponse{}
	err = json.Unmarshal(resp.Body(), &parsed)
	if err != nil {
		return nil, fmt.Errorf("failed to parse network clients response: %w", err)
	}

	if parsed.Meta.Code == "error" {
		return nil, errors.New(fmt.Sprintf("api error: %s", parsed.Meta.Message))
	}

	return parsed.Data, nil
}
