package cmd

import (
	"log"
	"net"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sendCmd)
}

var sendCmd = &cobra.Command{
	Use:   "send [mac address]",
	Short: "Send a magic packet to specified mac address",
	Long:  "Send a magic packet to wake up a device on the network using the specified mac address",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		mac, err := net.ParseMAC(args[0])
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

func sendMagicPacket(mac net.HardwareAddr) error {
	// Build magic packet
	// Create a buffer for the magic packet
	packet := make([]byte, 102)
	// Set the synchronization stream (first 6 bytes are 0xFF)
	for i := 0; i < 6; i++ {
		packet[i] = 0xFF
	}
	// Copy the MAC address 16 times into the packet
	for i := 1; i <= 16; i++ {
		copy(packet[i*6:], mac)
	}

	// Broadcast magic packet
	addr := &net.UDPAddr{
		IP:   net.IPv4bcast,
		Port: 9,
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(packet)
	return err
}
