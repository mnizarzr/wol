package cmd

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/mnizarzr/wol/magicpacket"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveApiCmd)
}

var serveApiCmd = &cobra.Command{
	Use:   "serveApi",
	Short: "Serve a minimal API to wake up machines",
	Long:  "Serve a minimal API that accepts MAC addresses and sends wake-up packets",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		mux := http.NewServeMux()

		mux.HandleFunc("POST /wake", handleApiWake)

		log.Printf("API listening on %s", cfg.Server.Listen)
		err := http.ListenAndServe(cfg.Server.Listen, mux)
		if err != nil {
			cobra.CheckErr(err)
		}
	},
}

func handleApiWake(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var request struct {
		MacAddress string `json:"mac_address"`
	}

	if err := json.Unmarshal(body, &request); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if request.MacAddress == "" {
		http.Error(w, "MAC address is required", http.StatusBadRequest)
		return
	}

	mac, err := net.ParseMAC(request.MacAddress)
	if err != nil {
		http.Error(w, "Invalid MAC address format", http.StatusBadRequest)
		return
	}

	log.Printf("Sending magic packet to %s", request.MacAddress)
	mp := magicpacket.NewMagicPacket(mac)
	if err := mp.Broadcast(); err != nil {
		log.Printf("Error sending magic packet: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
