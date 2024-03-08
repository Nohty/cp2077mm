package config

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

type Config struct {
	Version string
	GameDir string
	Mods    []ModConfig
}

type ModConfig struct {
	Name  string
	Files []string
}

type ConfigStore struct {
	ConfigPath string
}

func DefaultConfig() *Config {
	return &Config{
		Version: "1.0",
		GameDir: "",
		Mods:    []ModConfig{},
	}
}

func NewConfigStore() (*ConfigStore, error) {
	configFilePath, err := xdg.ConfigFile("cp2077mm/config.json")
	if err != nil {
		return nil, fmt.Errorf("could not resolve path for config file: %w", err)
	}

	configSore := &ConfigStore{ConfigPath: configFilePath}
	config, err := configSore.Config()
	if err != nil {
		return nil, fmt.Errorf("could not read configuration: %w", err)
	}

	if config.Version != "1.0" {
		return nil, fmt.Errorf("unsupported configuration version: %s", config.Version)
	}

	err = configSore.Save(config)
	if err != nil {
		return nil, fmt.Errorf("could not save configuration: %w", err)
	}

	return configSore, nil
}

func (c *ConfigStore) Config() (*Config, error) {
	_, err := os.Stat(c.ConfigPath)
	if os.IsNotExist(err) {
		return DefaultConfig(), nil
	}

	dir, fileName := filepath.Split(c.ConfigPath)
	if len(dir) == 0 {
		dir = "."
	}

	buf, err := fs.ReadFile(os.DirFS(dir), fileName)
	if err != nil {
		return nil, fmt.Errorf("could not read the configuration file: %w", err)
	}

	if len(buf) == 0 {
		return DefaultConfig(), nil
	}

	cfg := Config{}
	if err := json.Unmarshal(buf, &cfg); err != nil {
		return nil, fmt.Errorf("configuration file does not have a valid format: %w", err)
	}

	return &cfg, nil
}

func (c *ConfigStore) Save(cfg *Config) error {
	buf, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("could not marshal configuration to JSON: %w", err)
	}

	dir, _ := filepath.Split(c.ConfigPath)
	if len(dir) == 0 {
		dir = "."
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("could not create directory for configuration file: %w", err)
	}

	if err := os.WriteFile(c.ConfigPath, buf, 0644); err != nil {
		return fmt.Errorf("could not write configuration file: %w", err)
	}

	return nil
}

func (c *ConfigStore) Mods() ([]ModConfig, error) {
	cfg, err := c.Config()
	if err != nil {
		return nil, fmt.Errorf("could not read configuration: %w", err)
	}

	return cfg.Mods, nil
}

func (c *ConfigStore) Mod(name string) (*ModConfig, error) {
	cfg, err := c.Config()
	if err != nil {
		return nil, fmt.Errorf("could not read configuration: %w", err)
	}

	for _, m := range cfg.Mods {
		if m.Name == name {
			return &m, nil
		}
	}

	return nil, fmt.Errorf("mod with the name %s does not exist", name)
}

func (c *ConfigStore) Validate(mod ModConfig) error {
	if len(mod.Name) == 0 {
		return fmt.Errorf("mod name cannot be empty")
	}

	if len(mod.Files) == 0 {
		return fmt.Errorf("mod must have at least one file")
	}

	cfg, err := c.Config()
	if err != nil {
		return fmt.Errorf("could not read configuration: %w", err)
	}

	for _, m := range cfg.Mods {
		if m.Name == mod.Name {
			return fmt.Errorf("mod with the name %s already exists", mod.Name)
		}

		for _, f := range m.Files {
			for _, mf := range mod.Files {
				if f == mf {
					return fmt.Errorf("file %s is already used by another mod", f)
				}
			}
		}
	}

	return nil
}

func (c *ConfigStore) AddMod(mod ModConfig) error {
	cfg, err := c.Config()
	if err != nil {
		return fmt.Errorf("could not read configuration: %w", err)
	}

	if len(mod.Name) == 0 {
		return fmt.Errorf("mod name cannot be empty")
	}

	if len(mod.Files) == 0 {
		return fmt.Errorf("mod must have at least one file")
	}

	for _, m := range cfg.Mods {
		if m.Name == mod.Name {
			return fmt.Errorf("mod with the name %s already exists", mod.Name)
		}

		for _, f := range m.Files {
			for _, mf := range mod.Files {
				if f == mf {
					return fmt.Errorf("file %s is already used by another mod", f)
				}
			}
		}
	}

	cfg.Mods = append(cfg.Mods, mod)

	if err := c.Save(cfg); err != nil {
		return fmt.Errorf("could not save configuration: %w", err)
	}

	return nil
}

func (c *ConfigStore) RemoveMod(name string) error {
	cfg, err := c.Config()
	if err != nil {
		return fmt.Errorf("could not read configuration: %w", err)
	}

	for i, m := range cfg.Mods {
		if m.Name == name {
			cfg.Mods = append(cfg.Mods[:i], cfg.Mods[i+1:]...)
			break
		}
	}

	if err := c.Save(cfg); err != nil {
		return fmt.Errorf("could not save configuration: %w", err)
	}

	return nil
}
