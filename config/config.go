package config

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	URL    string `yaml:"url"`
	APIKey string `yaml:"apiKey"`
}

func Parse(r io.Reader) (Config, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return Config{}, fmt.Errorf("could not read config: %w", err)
	}
	var c Config
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return Config{}, fmt.Errorf("could not unmarshal data: %w", err)
	}
	return c, nil
}

func ParseFromFile(filename string) (Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return Config{}, fmt.Errorf("could not open file %s: %w", f, err)
	}
	defer f.Close() // nolint: errcheck
	return Parse(f)
}
