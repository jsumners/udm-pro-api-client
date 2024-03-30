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
func New(hostInfo HostInfo, options ...Option) (*Client, error) {
	client := &Client{
		hostInfo: hostInfo,
	}

	client.resty = *resty.New()
	client.resty.SetBaseURL(fmt.Sprintf("https://%s", hostInfo.Address))
	client.resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	// The UDM utilizes a CSRF token for non-read-only operations. Therefore,
	// we need to pick up that token on every response.
	client.resty.OnAfterResponse(func(c *resty.Client, res *resty.Response) error {
		token := res.Header().Get("X-Csrf-Token")
		if token == "" {
			return nil
		}

		if token != client.csrfToken {
			client.csrfToken = token
		}

		return nil
	})

	// And on every outgoing request, we need to add the CSRF token.
	client.resty.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		if client.csrfToken != "" {
			req.SetHeader("X-Csrf-Token", client.csrfToken)
		}
		return nil
	})

	for _, opt := range options {
		err := opt(client)
		if err != nil {
			return nil, err
		}
	}

	return client, nil
}

// NewWithLogin instantiates a [Client] instance and returns it.
// The instance will be authenticated with the remote UDM server
// and ready to issue requests. If an issue occurs while authenticating
// an error will be returned.
func NewWithLogin(hostInfo HostInfo, options ...Option) (*Client, error) {
	client, _ := New(hostInfo, options...)

	err := client.Login()
	if err != nil {
		return nil, err
	}

	return client, nil
}

// WithHttpDebug enables verbose printing of all HTTP requests.
func WithHttpDebug() Option {
	return func(c *Client) error {
		c.resty.SetDebug(true)
		return nil
	}
}

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login authenticates against the UDM. This must be done prior to issuing
// any requests to the API.
func (client *Client) Login() error {
	payload := credentials{
		Username: client.hostInfo.Username,
		Password: client.hostInfo.Password,
	}

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
