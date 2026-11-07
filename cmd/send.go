package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/mnizarzr/wol/mac"
	"github.com/mnizarzr/wol/magicpacket"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().StringP("mac", "m", "", "MAC address of the device to wake up or put to sleep.")
	sendCmd.Flags().StringP("name", "n", "", "Name of the device to wake up or put to sleep.")
	sendCmd.Flags().StringP("broadcast", "b", "", "Broadcast address to send the magic packet to (overrides config)")
	sendCmd.Flags().IntP("port", "p", 0, "UDP port to send the magic packet to (overrides config)")
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
		if action != "wake" && action != "sleep" {
			cobra.CheckErr(fmt.Errorf("action must be wake or sleep, got %q", action))
		}

		var macAddr net.HardwareAddr
		// Default to the global broadcast target; the --name path may refine it
		// with the machine's per-machine override.
		broadcast := cfg.Broadcast

		switch {
		case cmd.Flags().Changed("mac"):
			value, err := cmd.Flags().GetString("mac")
			if err != nil {
				cobra.CheckErr(err)
			}

			macAddr, err = net.ParseMAC(value)
			if err != nil {
				cobra.CheckErr(err)
			}
		case cmd.Flags().Changed("name"):
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				cobra.CheckErr(err)
			}

			machine, err := cfg.MachineByName(name)
			if err != nil {
				cobra.CheckErr(err)
			}

			macAddr, err = machine.HardwareAddr()
			if err != nil {
				cobra.CheckErr(err)
			}
			broadcast = cfg.BroadcastFor(machine)
		default:
			log.Fatalf("mac address should come from either --mac or --name")
		}

		if action == "sleep" {
			var err error
			macAddr, err = reverseMac(macAddr)
			if err != nil {
				cobra.CheckErr(err)
			}
		}

		// Command-line flags take precedence over config.
		if cmd.Flags().Changed("broadcast") {
			value, err := cmd.Flags().GetString("broadcast")
			if err != nil {
				cobra.CheckErr(err)
			}
			broadcast.Address = value
		}
		if cmd.Flags().Changed("port") {
			value, err := cmd.Flags().GetInt("port")
			if err != nil {
				cobra.CheckErr(err)
			}
			broadcast.Port = value
		}

		addr := broadcast.Addr()
		log.Printf("Sending magic packet to %s via %s (action: %s)", macAddr, addr, action)
		mp := magicpacket.NewMagicPacket(macAddr)
		if err := mp.Broadcast(addr); err != nil {
			cobra.CheckErr(err)
		}

		log.Printf("Magic packet sent for %s action", action)
	},
}

// reverseMac reverses the string representation of the MAC address
func reverseMac(addr net.HardwareAddr) (net.HardwareAddr, error) {
	reversedMacStr := mac.ReverseMacAddress(addr.String())
	return net.ParseMAC(reversedMacStr)
}
