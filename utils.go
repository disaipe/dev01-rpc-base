package rpc

import (
	"flag"
	"strings"
)

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func isService() bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "srv" || strings.HasPrefix(f.Name, "srv.") {
			found = true
		}
	})
	return found
}

func parseFlags() {
	var serviceFlag = false
	var helpFlag = false

	AppConfig = &Config{}

	flag.BoolVar(&AppConfig.serve, "serve", false, "Start HTTP server")
	flag.BoolVar(&serviceFlag, "srv", false, "Start as Windows service")
	flag.BoolVar(&AppConfig.isInstalling, "srv.install", false, "Install Windows service")
	flag.BoolVar(&AppConfig.isUninstalling, "srv.uninstall", false, "Uninstall Windows service")
	flag.BoolVar(&helpFlag, "help", false, "Usage help")

	flag.StringVar(&AppConfig.addr, "http.addr", ":8090", "Listening network address")
	flag.StringVar(&AppConfig.appUrl, "app.url", "", "Application hook URL")
	flag.StringVar(&AppConfig.appSecret, "app.secret", "", "Application secret")
}
