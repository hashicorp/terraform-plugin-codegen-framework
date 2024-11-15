package util

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type MethodSpecParams struct {
	Method   string `yaml:"method"`
	Path     string `yaml:"path"`
	Endpoint string `yaml:"endpoint"`
	Wait     bool   `yaml:"wait"`
}

type MethodSpec struct {
	Method   string `yaml:"method"`
	Path     string `yaml:"path"`
	Endpoint string `yaml:"endpoint"`
	Wait     bool   `yaml:"wait,omitempty"`
	Response string `yaml:"response,omitempty"`
}

type APIConfig struct {
	DtoName string              `yaml:"dto_name"`
	Create  *MethodSpec         `yaml:"create"`
	Read    *MethodSpec         `yaml:"read"`
	Update  []*MethodSpecParams `yaml:"update"`
	Delete  *MethodSpec         `yaml:"delete"`
}

type Config struct {
	Provider struct {
		Name     string `yaml:"name"`
		Env      string `yaml:"env"`
		Endpoint string `yaml:"endpoint"`
	} `yaml:"provider"`
	Resources map[string]APIConfig `yaml:"resources"`
}

func ExtractConfig(file string, resourceName string) (APIConfig, string, string, error) {
	f, err := os.Open(file)
	if err != nil {
		return APIConfig{}, "", "", fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	var config Config
	var env string
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&config); err != nil {
		return APIConfig{}, "", "", fmt.Errorf("failed to decode YAML: %w", err)
	}

	resource, ok := config.Resources[resourceName]
	if !ok {
		return APIConfig{}, "", "", fmt.Errorf("resource not found in config")
	}

	if config.Provider.Env == "" {
		env = "vpc"
	} else if !(config.Provider.Env == "vpc" || config.Provider.Env == "classic") {
		return APIConfig{}, "", "", fmt.Errorf("provider.env in config.yml must be set (vpc or classic)")
	} else {
		env = config.Provider.Env
	}

	return resource, env, config.Provider.Endpoint, nil
}
