package rpc

import (
	"flag"
)

func parseFlags() {
	var serviceFlag = false
	var helpFlag = false

	flag.BoolVar(&AppConfig.serve, "serve", false, "Start HTTP server")
	flag.BoolVar(&serviceFlag, "srv", false, "Start as Windows service")
	flag.BoolVar(&AppConfig.isInstalling, "srv.install", false, "Install Windows service")
	flag.BoolVar(&AppConfig.isUninstalling, "srv.uninstall", false, "Uninstall Windows service")
	flag.BoolVar(&helpFlag, "help", false, "Usage help")

	flag.StringVar(&AppConfig.addr, "http.addr", ":8090", "Listening network address")
	flag.StringVar(&AppConfig.appUrl, "app.url", "", "Application hook URL")
	flag.StringVar(&AppConfig.appSecret, "app.secret", "", "Application secret")
}
