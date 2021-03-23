package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:    "SmartHomePi",
	Short:  "SmartHomePi is Home automation suite",
	Long:   `SmartHomePi is self hosted Home automation suite.`,
	PreRun: webCmd.PreRun,
	Run:    webCmd.Run,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
