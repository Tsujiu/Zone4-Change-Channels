package main

import (
	"log"
	"net/http"

	"mu-server-manager/config"
	"mu-server-manager/controllers"
)

func main() {
	err := config.LoadChannels("config/channels.json")
	if err != nil {
		log.Fatalf("Failed to load channels: %v", err)
	}

	// Rotas da API (ordem importante!)
	http.HandleFunc("/start/", controllers.StartChannelHandler)
	http.HandleFunc("/stop/", controllers.StopChannelHandler)
	http.HandleFunc("/status", controllers.StatusHandler)
	http.HandleFunc("/start-all", controllers.StartAllHandler)
	http.HandleFunc("/stop-all", controllers.StopAllHandler)

	// Servir logs dos canais (acesso: /logs/channel1.log etc)
	http.Handle("/logs/", http.StripPrefix("/logs/", http.FileServer(http.Dir("./logs"))))

	// Servir painel web estático (web/index.html)
	http.Handle("/", http.FileServer(http.Dir("./web")))

	log.Println("Channel Manager running on :8080")
	http.ListenAndServe(":8080", nil)
}
