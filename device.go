package udm

import (
	"encoding/json"
	"errors"
	"fmt"
)

type RestartDevicePayload struct {
	// MacAddress is used to identify the device that will be restarted.
	MacAddress string `json:"mac"`

	// RebootType indicates how the device should be rebooted. Known values:
	// + "soft"
	RebootType string `json:"reboot_type"`

	// Command to be issued to the device. Known values:
	// + "restart"
	Command string `json:"cmd"`
}

// RestartDevice is used to trigger a device restart for devices that are
// managed by the UDM, e.g. a wireless access point. "Site Admin" level
// permissions are needed in order for this operation to work.
func (client *Client) RestartDevice(macAddress string) error {
	path := fmt.Sprintf("/proxy/network/api/s/%s/cmd/devmgr", client.hostInfo.Site)
	payload := RestartDevicePayload{
		MacAddress: macAddress,
		RebootType: "soft",
		Command:    "restart",
	}

	resp, err := client.resty.R().
		SetBody(payload).
		SetHeader("content-type", "application/json").
		Post(path)
	if err != nil {
		return fmt.Errorf("failed to restart device (%s): %w", macAddress, err)
	}

	if resp.StatusCode() < 200 || resp.StatusCode() > 299 {
		return errors.New(fmt.Sprintf("http error: %s", resp.Status()))
	}

	var parsed GenericResponse
	err = json.Unmarshal(resp.Body(), &parsed)
	if err != nil {
		return fmt.Errorf("failed to parse restart response: %w", err)
	}

	if parsed.Meta.Code == "error" {
		return errors.New(fmt.Sprintf("api error: %s", parsed.Meta.Message))
	}

	return nil
}
