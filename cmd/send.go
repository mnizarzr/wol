package cmd

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/mnizarzr/wol/mac"
	"github.com/mnizarzr/wol/magicpacket"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().StringP("mac", "m", "", "MAC address of the device to wake up or put to sleep.")
	sendCmd.Flags().StringP("name", "n", "", "Name of the device to wake up or put to sleep.")
}

var sendCmd = &cobra.Command{
	Use:   "send <action>",
	Short: "Send a magic packet to specified mac address",
	Long:  "Send a magic packet to wake up a device on the network using the specified mac address",
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Only one of the flags should be specified
		if cmd.Flags().Changed("mac") == cmd.Flags().Changed("name") {
			return fmt.Errorf("either --mac or --name must be specified")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		action := args[0]

		macAddr, err := getTargetMacAddress(cmd)
		if err != nil {
			cobra.CheckErr(err)
		}

		if action == "sleep" {
			macAddr, err = reverseMac(macAddr)
			if err != nil {
				cobra.CheckErr(err)
			}
		}

		log.Printf("Sending magic packet to %s (action: %s)", macAddr, action)

		mp := magicpacket.NewMagicPacket(macAddr)
		if err := mp.Broadcast(); err != nil {
			cobra.CheckErr(err)
		}

		log.Printf("Magic packet sent for %s action", action)
	},
}

// getTargetMacAddress determines the MAC address based on either --mac or --name
func getTargetMacAddress(cmd *cobra.Command) (net.HardwareAddr, error) {
	switch {
	case cmd.Flags().Changed("mac"):
		macStr, err := cmd.Flags().GetString("mac")
		if err != nil {
			return nil, err
		}
		return net.ParseMAC(macStr)

	case cmd.Flags().Changed("name"):
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return nil, err
		}
		return getMacByName(name, false)

	default:
		return nil, fmt.Errorf("either --mac or --name must be specified")
	}
}

// getMacByName returns the MAC address of the machine with the specified name
func getMacByName(name string, reverse bool) (net.HardwareAddr, error) {
	for _, machine := range cfg.Machines {
		if strings.EqualFold(machine.Name, name) {
			m := machine.Mac
			if reverse {
				m = mac.ReverseMacAddress(m)
			}
			macAddr, err := net.ParseMAC(m)
			if err != nil {
				return nil, fmt.Errorf("failed to parse MAC address: %w", err)
			}
			return macAddr, nil
		}
	}
	return nil, fmt.Errorf("machine with name %q not found", name)
}

// reverseMac reverses the string representation of the MAC address
func reverseMac(addr net.HardwareAddr) (net.HardwareAddr, error) {
	reversedMacStr := mac.ReverseMacAddress(addr.String())
	return net.ParseMAC(reversedMacStr)
}
