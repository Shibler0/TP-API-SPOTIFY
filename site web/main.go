package main

import (
	"TpSpotify/router"
	"fmt"
	"net/http"
)

func main() {
	// Charge le routeur
	r := router.New()

	//choisi la route de départ
	fmt.Println("🚀 Serveur démarré sur http://localhost:8080/home")
	//crée le serveur
	http.ListenAndServe(":8080", r)

}
