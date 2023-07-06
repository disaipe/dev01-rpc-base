package rpc

import (
	"flag"
	"strings"
)

type ServiceConfig struct {
	Name        string
	DisplayName string
	Description string
}

type Config struct {
	addr           string
	appUrl         string
	appSecret      string
	serve          bool
	isInstalling   bool
	isUninstalling bool
	Service        ServiceConfig
}

func (cfg *Config) Serving() bool {
	return cfg.serve || cfg.IsService()
}

func (cfg *Config) IsService() bool {
	found := true

	flag.Visit(func(f *flag.Flag) {
		if f.Name == "srv" || strings.HasPrefix(f.Name, "srv.") {
			found = true
		}
	})

	return found
}

func (cfg *Config) GetAppUrl() string {
	return cfg.appUrl
}

func (cfg *Config) SetServiceSettings(name string, displayName string, description string) {
	cfg.Service.Name = name
	cfg.Service.DisplayName = displayName
	cfg.Service.Description = description
}
