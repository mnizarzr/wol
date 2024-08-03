package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sendCmd)
}

var sendCmd = &cobra.Command{
	Use:   "send [mac address]",
	Short: "Send a magic packet to wake up a device",
	Long:  "Send a magic packet to wake up a device on the network",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		mac, err := validateMacAddress(args[0])
		if err != nil {
			cobra.CheckErr(err)
		}

		log.Printf("Sending magic packet to %s", mac)
		err = sendMagicPacket(mac)
		if err != nil {
			cobra.CheckErr(err)
		}

		log.Printf("Magic packet sent")
	},
}

func validateMacAddress(mac string) (string, error) {
	// TODO: Validate MAC address format
	return mac, nil
}

func sendMagicPacket(mac string) error {
	// TODO: Send magic packet to the device
	return nil
}
