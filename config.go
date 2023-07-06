package rpc

type Config struct {
	addr      string
	appUrl    string
	appSecret string

	serve          bool
	isService      bool
	isInstalling   bool
	isUninstalling bool
}

func (cfg *Config) Serving() bool {
	return cfg.serve || isService()
}

func (cfg *Config) GetAppUrl() string {
	return cfg.appUrl
}
