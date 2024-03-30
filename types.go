package udm

import "github.com/go-resty/resty/v2"

// HostInfo describes the details needed to connect to the desired UDM.
type HostInfo struct {
	// Address is the IP address or hostname, e.g. "192.168.1.1" or
	// "router.internal".
	Address string

	// Username is the user to connect as.
	Username string

	// Password is the password for the provided user.
	Password string

	// Site is an internal identifier used by the UDM to target devices.
	// It is most likely that you should be using "default" for the value unless
	// have specific knowledge otherwise.
	Site string
}

type Client struct {
	hostInfo HostInfo
	resty    resty.Client
}

// ResponseMeta represents the `meta` property in a UDM API response.
type ResponseMeta struct {
	// Code is the response code returned by operations. Typically, "ok".
	Code    string `json:"rc"`
	Message string `json:"msg"`

	// TODO: I believe a `count` can also be returned
	// to indicate pagination. But I don't have enough
	// data to work with to determine how pagination
	// would work.
}
