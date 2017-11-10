package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// Config holds the settings for the Kube Bot application
type Config struct {
	// SlackAPIToken is the token to use for Slack API authentication
	SlackAPIToken string

	// KubeConfigPath is the path to a Kubernetes auth configuration file.
	// Which is normally created by the kubectl.
	// If empty the value will default to "$HOME/.kube/config".
	KubeConfigPath string

	// Bot holds bot specified configuration
	Bot BotConfig
}

// BotConfig holds Bot specified settings. Which are embedded in the Config
// struct in the bot field.
type BotConfig struct {
	// Name is the Slack username the Bot will respond to
	Name string
}

// PathEnvKey is the name of the environment variable which can be used to
// modify which file LoadConfig loads.
const PathEnvKey string = "APP_CONFIG_PATH"

// LoadConfig reads a toml configuration file and loads the values in a Config
// struct, which is then returned. If the process is successful a nil error
// will be returned. The error otherwise.
//
// By default it will load from the "config.toml" file in the working
// directory. This can be modified by setting the APP_CONFIG_PATH environment
// variable.
func LoadConfig() (Config, error) {
	var config Config
	var filePath string

	// Default file location
	if cwd, err := os.Getwd(); err == nil {
		filePath = filepath.Join(cwd, "config.toml")
	} else {
		return config, fmt.Errorf("error getting working directory: %s", err.Error())
	}

	// Check if PathEnvKey is set
	if val, set := os.LookupEnv(PathEnvKey); set {
		// Override filePath with custom user provided value
		filePath = val
	}

	// Check if filePath exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return config, fmt.Errorf("config path does not exist: %s", filePath)
	}

	// Read file contents
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, fmt.Errorf("error reading config file \"%s\": %s", filePath, err.Error())
	}

	// Load toml
	if _, err := toml.Decode(string(data), &config); err != nil {
		return config, fmt.Errorf("failed to parse toml config file \"%s\": %s", filePath, err.Error())
	}

	// All done, return
	return config, nil
}
