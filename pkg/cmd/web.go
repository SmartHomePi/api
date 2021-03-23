package cmd

import (
	"os"
	"os/signal"

	"github.com/SmartHomePi/api/pkg/config"
	"github.com/SmartHomePi/api/pkg/initialize"
	"github.com/SmartHomePi/api/pkg/logger"
	"github.com/SmartHomePi/api/pkg/routes"
	"github.com/SmartHomePi/api/pkg/thing"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(webCmd)
}

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Starts the rest api web server",
	PreRun: func(cmd *cobra.Command, args []string) {
		initialize.FullInit()
	},
	Run: func(cmd *cobra.Command, args []string) {

		router := routes.NewGin()

		go func() {
			err := router.Run(config.ServerHttpIP.GetString() + ":" + config.ServerHttpPort.GetString())
			if err != nil {
				logger.Error(err.Error())
				os.Exit(1)
			}
		}()

		go thing.StartThingServer()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit

		logger.Info("Shutting down...")

	},
}
