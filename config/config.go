package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

var k = koanf.New(".")

type Machine struct {
	Name string `koanf:"name"`
	Mac  string `koanf:"mac"`
}

type Config struct {
	Machines []Machine `koanf:"machines"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	// Order here matters as later values will override earlier ones
	paths := []string{
		filepath.Join(home, ".wol", "config.yaml"),
		filepath.Join(".", "config.yaml"),
	}

	for _, path := range paths {
		err = k.Load(file.Provider(path), yaml.Parser())

		// Ignore error if file does not exist
		if err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to load config file: %w", err)
		}
	}

	err = k.Unmarshal("", c)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return nil
}
