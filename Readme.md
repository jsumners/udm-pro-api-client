# udm-pro-api-client

An API client for communicating with the [Unifi UDM-PRO][udm-pro].

This is known to work with:

+ `UniFi OS 1.12.22` and `Network 7.1.68`
+ `UniFi OS 3.1.16` and `Network 8.0.7`

## CLI Tool

The provided CLI tool supports the following commands:

+ `get-hosts`: retrieves the connected devices and generates a hosts file.
Required permissions: "View".
+ `device restart`: restarts a device managed by the UDM, e.g. a wireless access
point. Required permissions: "Site Admin".

See the `--help` output of individual commands for full usage details, e.g.:

```sh
$ udm-pro-api-client device restart --help
```

## Configuration

The client can be configured with any sort of configuration file supported
by [Viper][viper]. All features of the client, except one, are also
configurable by environment variables. The exception is host aliases (described
below).

By default, the client will look for a file named `udm-pro-api-client` with
an appropriate extension (e.g. `.yaml`). A specific configuration file can
be specified with the environment variable `API_CLIENT_CONFIG_FILE`, e.g.:

```sh
$ API_CLIENT_CONFIG_FILE=/opt/api-client.yaml udm-pro-api-client
```
The configuration file can also be specified with a flag:

```sh
$ udm-pro-api-client --conf-file ./config.yaml
# or
$ udm-pro-api-client -c ./config.yaml
```

A full configuration in yaml is:

```yaml
# Specifies the remote IP address of the UDM-PRO.
# Env var: API_CLIENT_ADDRESS
address: 192.168.1.1

# Specifies the username to connect to the UDM-PRO with.
# Env var: API_CLIENT_USERNAME
username: api

# Specifies the password to connect to the UDM-PRO with.
# Env var: API_CLIENT_PASSWORD
password: super-secret

# Indicates if only statically assigned clients should be considered.
# These are clients that have had an IP address assigned to them through the
# Network OS interface instead of randomly assigned from the pool of available
# dynamic IP addresses.
# Env var: API_CLIENT_FIXED_ONLY
# Default: true
fixed_only: true

# Indicates if all discovered hostnames should be lowercased before writing to
# the hosts file that Dnsmasq will read.
# Env var: API_CLIENT_LOWERCASE_HOSTNAMES
# Default: true
lowercase_hostnames: true

# A set of names and IP addresses to add to the hosts file that Dnsmasq will
# read. This allows for setting of names that the UDM-PRO does not manage, For
# example, the UDM-PRO itself.
# This is not configurable via envionment variables.
# Default: an empty list.
host_aliases:
  - name: router
    ip_address: 192.168.1.1
```


[udm-pro]: https://store.ui.com/products/udm-pro
[viper]: https://github.com/spf13/viper/tree/5247643f02358b40d01385b0dbf743b659b0133f#reading-config-files
