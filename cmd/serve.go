package cmd

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

//go:embed templates/*
var templates embed.FS

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve a web interface to wake up machines",
	Long:  "Serve a web interface that lists all the configured machines and allows you to wake them up",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		mux := http.NewServeMux()

		mux.HandleFunc("GET /{$}", handleIndex)
		mux.HandleFunc("POST /wake", handleWake)

		log.Printf("Listening on %s", cfg.Server.Listen)
		err := http.ListenAndServe(cfg.Server.Listen, mux)
		if err != nil {
			cobra.CheckErr(err)
		}
	},
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// Parse the template
	index, err := template.ParseFS(templates, "templates/index.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the template
	err = index.Execute(w, map[string]interface{}{
		"Machines": cfg.Machines,
	})
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handleWake(w http.ResponseWriter, r *http.Request) {
	// If we were to get MAC address, validate if it exists in the config
	// If it does, send the magic packet

	mac, err := getMacByName(r.FormValue("name"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Sending magic packet to %s", mac)
	err = sendMagicPacket(mac)
	if err != nil {
		log.Printf("Error sending magic packet: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
