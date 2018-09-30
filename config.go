package kaf

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	yaml "gopkg.in/yaml.v2"
)

type SASL struct {
	Mechanism string
	Username  string
	Password  string
}

type Cluster struct {
	Name             string
	Brokers          []string `yaml:"brokers"`
	SASL             *SASL    `yaml:"SASL"`
	SecurityProtocol string   `yaml:"security-protocol"`
}

type Config struct {
	CurrentCluster string     `yaml:"current-cluster"`
	Clusters       []*Cluster `yaml:"clusters"`
}

func (c *Config) ActiveCluster() *Cluster {
	return c.Clusters[0]
}

func (c *Config) Write() error {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
	}

	configPath := filepath.Join(home, ".kaf", "config")
	file, err := os.OpenFile(configPath, os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	encoder := yaml.NewEncoder(file)
	return encoder.Encode(&c)

}

func ReadConfig() (c *Config, err error) {
	file, err := os.OpenFile(getDefaultConfigPath(), os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func getDefaultConfigPath() string {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	return filepath.Join(home, ".kaf", "config")
}
