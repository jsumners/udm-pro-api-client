package udm

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

// New instantiates a [Client] instance and returns it.
// The client will need to issue a [Client.Login] before any requests to the API
// can be issued.
func New(hostInfo HostInfo) (*Client, error) {
	client := &Client{
		hostInfo: hostInfo,
	}

	client.resty = *resty.New()
	client.resty.SetBaseURL(fmt.Sprintf("https://%s", hostInfo.Address))
	client.resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	return client, nil
}

// NewWithLogin instantiates a [Client] instance and returns it.
// The instance will be authenticated with the remote UDM server
// and ready to issue requests. If an issue occurs while authenticating
// an error will be returned.
func NewWithLogin(hostInfo HostInfo) (*Client, error) {
	client, _ := New(hostInfo)

	err := client.Login()
	if err != nil {
		return nil, err
	}

	return client, nil
}

// Login authenticates against the UDM. This must be done prior to issuing
// any requests to the API.
func (client *Client) Login() error {
	payload := fmt.Sprintf(
		`{"username":"%s","password":"%s"}`,
		client.hostInfo.Username,
		client.hostInfo.Password,
	)

	resp, err := client.resty.R().
		SetHeader("content-type", "application/json").
		SetBody(payload).
		Post("/api/auth/login")

	if err != nil {
		return fmt.Errorf("failed to login: %w", err)
	}
	if resp.StatusCode() >= 400 {
		return errors.New(fmt.Sprintf("login error: %s", resp.Status()))
	}

	return nil
}
