package rpc

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kardianos/service"
)

type Daemon struct{}

func (p Daemon) Start(s service.Service) error {
	rpc := &Rpc{}

	go func() {
		rpc.serve(Config.addr)
	}()

	return nil
}

func (p Daemon) Stop(s service.Service) error {
	return nil
}

func runService() {
	if Config.isInstalling {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Application hook URL (required): ")
		appUrl, _ := reader.ReadString('\n')
		appUrl = strings.Replace(appUrl, "\n", "", -1)

		fmt.Print("Application secret: ")
		appSecret, _ := reader.ReadString('\n')
		appSecret = strings.Replace(appSecret, "\n", "", -1)

		srv := getService([]string{
			"-srv",
			fmt.Sprintf("-app.url=%v", appUrl),
			fmt.Sprintf("-app.secret=%v", appSecret),
		})

		err := srv.Install()

		if err != nil {
			Logger.Fatal().Msgf("Cannot install service: %v\n", err)
		}

		Logger.Info().Msgf("Service installed")
		os.Exit(0)
	}

	if Config.isUninstalling {
		srv := getService([]string{})
		err := srv.Uninstall()

		if err != nil {
			Logger.Fatal().Msgf("Cannot uninstall service: %v\n", err)
		}

		Logger.Info().Msgf("Service uninstalled")
		os.Exit(0)
	}

	srv := getService([]string{})
	err := srv.Run()

	if err != nil {
		Logger.Fatal().Msgf("Cannot start the service: %v\n", err)
	}

	os.Exit(0)
}

func getService(args []string) service.Service {

	serviceConfig := &service.Config{
		Name:        Config.Service.Name,
		DisplayName: Config.Service.DisplayName,
		Description: Config.Service.Description,
		Arguments:   args,
	}

	prg := &Daemon{}

	srv, err := service.New(prg, serviceConfig)

	if err != nil {
		Logger.Fatal().Msgf("Cannot create the service: %v\n", err)
	}

	return srv
}
