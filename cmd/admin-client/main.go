package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type GlobalFlags struct {
	Address    string
	ApiVersion string
}

var globalFlags = GlobalFlags{}

func initFlags(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().StringVarP(&globalFlags.Address, "address", "a", "http://localhost:8080", "address of the http server")
	rootCmd.PersistentFlags().StringVarP(&globalFlags.ApiVersion, "version", "v", "v1", "version of the admin request api")
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "adminctl",
		Short: "adminCtl is a command line tool for administering the HStreamDB cluster",
	}
	initFlags(rootCmd)
	rootCmd.AddCommand(
		newStatsCmd(),
		newStatusCmd(),
		newStreamCmd(),
		newSubscriptionCmd(),
	)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
