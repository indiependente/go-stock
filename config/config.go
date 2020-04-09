package config

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Config holds the configuration information of the tool.
type Config struct {
	URL    string `yaml:"url"`
	APIKey string `yaml:"apiKey"`
}

// Parse parses the input io.Reader into a Config.
// Returns an error if any.
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

// ParseFromFile reads the content of the input filename and parses it into a Config.
// Returns an error if any.
func ParseFromFile(filename string) (Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return Config{}, fmt.Errorf("could not open file %s: %w", filename, err)
	}
	defer f.Close() // nolint: errcheck
	return Parse(f)
}

// Save persists the configuration c in the location pointed by filename.
// Returns an error if any.
func Save(c Config, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file: %w", err)
	}
	defer f.Close()
	data, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("could not marshal configuration: %w", err)
	}
	_, err = f.Write(data)
	if err != nil {
		return fmt.Errorf("could not write to file: %w", err)
	}
	return nil
}
