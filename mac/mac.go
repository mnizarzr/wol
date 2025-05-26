package mac

import (
	"net"
	"strings"
)

type MacAddr net.HardwareAddr

const hexDigit = "0123456789abcdef"

func (a MacAddr) StringReversed() string {
	if len(a) == 0 {
		return ""
	}
	buf := make([]byte, 0, len(a)*3-1)
	for i := len(a) - 1; i >= 0; i-- {
		b := a[i]
		if i < len(a)-1 {
			buf = append(buf, ':')
		}
		buf = append(buf, hexDigit[b>>4])
		buf = append(buf, hexDigit[b&0x0F])
	}
	return string(buf)
}

func ReverseMacAddress(address string) string {
	tokens := strings.Split(address, ":")
	last := len(tokens) - 1
	for i := 0; i < len(tokens)/2; i++ {
		tokens[i], tokens[last-i] = tokens[last-i], tokens[i]
	}
	return strings.Join(tokens, ":")
}
