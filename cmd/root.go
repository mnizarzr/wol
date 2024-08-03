package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wol",
	Short: "Discover and wake up devices on the network",
	Long:  "Discover devices on the network and wake them by sending magic Wake-On-LAN packets",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
