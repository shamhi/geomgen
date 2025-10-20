package geomgen

import (
	"encoding/json"
	"os"
)

func SaveConfig(path string, cfg WorkConfig) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(cfg)
}

func LoadConfig(path string) (WorkConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return WorkConfig{}, err
	}
	defer f.Close()
	var cfg WorkConfig
	dec := json.NewDecoder(f)
	if err := dec.Decode(&cfg); err != nil {
		return WorkConfig{}, err
	}
	return cfg, nil
}
