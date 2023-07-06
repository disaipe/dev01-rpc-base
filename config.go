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

type config struct {
	addr           string
	appUrl         string
	appSecret      string
	serve          bool
	isInstalling   bool
	isUninstalling bool
	actions        map[string]*ActionFunction
	Service        ServiceConfig
}

func (cfg *config) Serving() bool {
	return cfg.serve || cfg.IsService()
}

func (cfg *config) IsService() bool {
	found := true

	flag.Visit(func(f *flag.Flag) {
		if f.Name == "srv" || strings.HasPrefix(f.Name, "srv.") {
			found = true
		}
	})

	return found
}

func (cfg *config) GetAppUrl() string {
	return cfg.appUrl
}

func (cfg *config) SetServiceSettings(name string, displayName string, description string) {
	cfg.Service.Name = name
	cfg.Service.DisplayName = displayName
	cfg.Service.Description = description
}

func (cfg *config) SetAction(uri string, action *ActionFunction) {
	if cfg.actions == nil {
		cfg.actions = map[string]*ActionFunction{}
	}

	cfg.actions[uri] = action
}

func (cfg *config) GetAction(uri string) *ActionFunction {
	return cfg.actions[uri]
}

func (cfg *config) GetActions() map[string]*ActionFunction {
	return cfg.actions
}
