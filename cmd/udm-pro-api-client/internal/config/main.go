package config

import (
	"errors"
	"fmt"
	"github.com/creasty/defaults"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type HostAlias struct {
	Name      string `mapstructure:"name"`
	IpAddress string `mapstructure:"ip_address"`
}

type Configuration struct {
	*viper.Viper `mapstructure:"-"`

	Address  string `mapstructure:"address"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Site     string `mapstructure:"site"`

	FixedOnly          bool `mapstructure:"fixed_only"`
	LowercaseHostnames bool `mapstructure:"lowercase_hostnames"`

	HostAliases []HostAlias `mapstructure:"host_aliases"`
}

func New() *Configuration {
	return &Configuration{
		Viper: viper.New(),
	}
}

func (c *Configuration) InitConfig(configFilePath string) error {
	c.SetConfigName("udm-pro-api-client")
	c.SetConfigType("yaml")
	c.AddConfigPath(".")
	c.SetEnvPrefix("API_CLIENT")
	c.AutomaticEnv()
	c.SetDefault("site", "default")
	c.SetDefault("fixed_only", true)
	c.SetDefault("lowercase_hostnames", true)

	envFile := c.GetString("config_file")
	if configFilePath != "" {
		// --conf-file flag has been provided. Prefer its value.
		c.SetConfigFile(configFilePath)
	} else if envFile != "" {
		// Fallback to the API_CLIENT_CONFIG_FILE value.
		c.SetConfigFile(envFile)
	}

	err := c.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil
		}
		return fmt.Errorf("unable to read configuration file: %w", err)
	}

	err = c.Unmarshal(c)
	if err != nil {
		return fmt.Errorf("unable to unmarshal configuration: %w", err)
	}

	return nil
}

func (c *Configuration) GenerateCurrentYaml() (string, error) {
	encodedData := make(map[string]any)
	err := mapstructure.Decode(c, &encodedData)
	if err != nil {
		return "", fmt.Errorf("unable to encode configuration: %w", err)
	}

	yamlEncoded, err := yaml.Marshal(encodedData)
	if err != nil {
		return "", fmt.Errorf("unable to marshal configuration to yaml: %w", err)
	}

	return string(yamlEncoded), nil
}

func (c *Configuration) GenerateDefaultYaml() (string, error) {
	defaultConfig := Configuration{}
	err := defaults.Set(&defaultConfig)
	if err != nil {
		return "", fmt.Errorf("unable to generate default config: %w", err)
	}

	// We decode to a generic interface in order to rename the struct fields
	// according to the `mapstructure` tag.
	var encoded map[string]any
	err = mapstructure.Decode(defaultConfig, &encoded)
	if err != nil {
		return "", fmt.Errorf("unable to decode configuration: %w", err)
	}

	yamlEncoded, err := yaml.Marshal(encoded)
	if err != nil {
		return "", fmt.Errorf("unable to encode configuration to yaml: %w", err)
	}

	return string(yamlEncoded), nil
}
