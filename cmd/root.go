package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/trugamr/wol/config"
)

var cfg = config.NewConfig()

var rootCmd = &cobra.Command{
	Use:   "wol",
	Short: "Discover and wake up devices on the network",
	Long:  "Discover devices on the network and wake them by sending magic Wake-On-LAN packets",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return cfg.Load()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
