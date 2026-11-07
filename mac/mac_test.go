package mac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseMacAddress(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"01:02:03:04:05:06", "06:05:04:03:02:01"},
		{"aa:bb:cc:dd:ee:ff", "ff:ee:dd:cc:bb:aa"},
		{"00:00:00:00:00:00", "00:00:00:00:00:00"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, ReverseMacAddress(tt.in), "in=%q", tt.in)
	}
}
