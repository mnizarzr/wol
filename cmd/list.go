package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List machines from config file",
	Long:  "Show a list of all the configured machines",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(cfg.Machines) == 0 {
			fmt.Println("No machines configured")
			return
		}

		// Render the list of machines
		fmt.Println("Name\tMAC")
		for _, machine := range cfg.Machines {
			fmt.Printf("%s\t%s\n", machine.Name, machine.Mac)
		}
	},
}
