package cmd

import (
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mnizarzr/wol/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleApiWake(t *testing.T) {
	prev := cfg
	t.Cleanup(func() { cfg = prev })
	cfg = &config.Config{
		Broadcast: config.Broadcast{Address: "255.255.255.255", Port: 9},
	}

	t.Run("invalid mac", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/wake", strings.NewReader(`{"mac_address":"not-a-mac"}`))
		rec := httptest.NewRecorder()
		handleApiWake(rec, req)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("missing mac", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/wake", strings.NewReader(`{}`))
		rec := httptest.NewRecorder()
		handleApiWake(rec, req)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("valid mac broadcasts", func(t *testing.T) {
		// Broadcast to a loopback UDP target that DialUDP accepts; we only
		// assert the handler returns success after building a valid packet.
		cfg.Broadcast = config.Broadcast{Address: "127.0.0.1", Port: 9}
		req := httptest.NewRequest(http.MethodPost, "/wake", strings.NewReader(`{"mac_address":"01:02:03:04:05:06"}`))
		rec := httptest.NewRecorder()
		handleApiWake(rec, req)
		assert.Equal(t, http.StatusNoContent, rec.Code)
	})

	t.Run("body too large", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/wake", strings.NewReader(strings.Repeat("x", maxAPIBody+1)))
		rec := httptest.NewRecorder()
		handleApiWake(rec, req)
		assert.Equal(t, http.StatusRequestEntityTooLarge, rec.Code)
	})
}

func TestHandleSleepReversesMAC(t *testing.T) {
	const macStr = "01:02:03:04:05:06"
	cfg := &config.Config{
		Broadcast: config.Broadcast{Address: "255.255.255.255", Port: 9},
		Machines:  []config.Machine{{Name: "alpha", Mac: macStr}},
	}

	var gotMac net.HardwareAddr
	wake := func(mac net.HardwareAddr, addr string) error {
		gotMac = mac
		return nil
	}
	req := httptest.NewRequest(http.MethodPost, "/sleep", strings.NewReader("name=alpha"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	newServer(cfg, nil, wake).routes().ServeHTTP(rec, req)

	require.Equal(t, http.StatusSeeOther, rec.Code)
	want, err := reverseMac(mustParseMAC(t, macStr))
	require.NoError(t, err)
	assert.Equal(t, want, gotMac)
}

func mustParseMAC(t *testing.T, s string) net.HardwareAddr {
	t.Helper()
	mac, err := net.ParseMAC(s)
	require.NoError(t, err)
	return mac
}
