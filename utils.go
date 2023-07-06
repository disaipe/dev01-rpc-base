package rpc

import (
	"flag"
)

func parseFlags() {
	var serviceFlag = false
	var helpFlag = false

	flag.BoolVar(&Config.serve, "serve", false, "Start HTTP server")
	flag.BoolVar(&serviceFlag, "srv", false, "Start as Windows service")
	flag.BoolVar(&Config.isInstalling, "srv.install", false, "Install Windows service")
	flag.BoolVar(&Config.isUninstalling, "srv.uninstall", false, "Uninstall Windows service")
	flag.BoolVar(&helpFlag, "help", false, "Usage help")

	flag.StringVar(&Config.addr, "http.addr", ":8090", "Listening network address")
	flag.StringVar(&Config.appUrl, "app.url", "", "Application hook URL")
	flag.StringVar(&Config.appSecret, "app.secret", "", "Application secret")
}
