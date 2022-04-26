package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type GlobalFlags struct {
	Address    string
	PrefixPath string
}

var globalFlags = GlobalFlags{}

func initFlags(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().StringVarP(&globalFlags.Address, "address", "a", "http://localhost:8080", "address of the http server")
	rootCmd.PersistentFlags().StringVarP(&globalFlags.PrefixPath, "prefix", "p", "v1", "prefix path of the admin request URL")
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
