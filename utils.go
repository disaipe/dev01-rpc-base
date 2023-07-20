package rpc

import (
	"flag"
)

func parseFlags() {
	flag.BoolVar(&Config.serve, "serve", false, "Start HTTP server")
	flag.BoolVar(&Config.isService, "srv", false, "Start as Windows service")
	flag.BoolVar(&Config.isInstalling, "srv.install", false, "Install Windows service")
	flag.BoolVar(&Config.isUninstalling, "srv.uninstall", false, "Uninstall Windows service")

	flag.StringVar(&Config.addr, "http.addr", ":8090", "Listening network address")
	flag.StringVar(&Config.appUrl, "app.url", "", "Application hook URL")
	flag.StringVar(&Config.appSecret, "app.secret", "", "Application secret")
}
