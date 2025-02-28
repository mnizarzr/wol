package magicpacket

import "net"

// MagicPacket represents a wake-on-LAN packet
type MagicPacket struct {
	// The MAC address of the machine to wake up
	MacAddress net.HardwareAddr
}

// NewMagicPacket creates a new MagicPacket for the given MAC address
func NewMagicPacket(macAddress net.HardwareAddr) *MagicPacket {
	return &MagicPacket{MacAddress: macAddress}
}

// Broadcast sends the magic packet to the broadcast address
func (p *MagicPacket) Broadcast() error {
	// Build the actual packet
	packet := make([]byte, 102)
	// Set the synchronization stream (first 6 bytes are 0xFF)
	for i := 0; i < 6; i++ {
		packet[i] = 0xFF
	}
	// Copy the MAC address 16 times into the packet
	for i := 1; i <= 16; i++ {
		copy(packet[i*6:], p.MacAddress)
	}

	// Broadcast the packet
	// TODO: Broadcast to more common ports and addresses?
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
