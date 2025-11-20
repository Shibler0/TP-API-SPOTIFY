package main

import (
	"fmt"
	"net/http"
	"siteweb/router"
)

func main() {
	// Charge le routeur
	r := router.New()

	//choisi la route de départ
	fmt.Println("🚀 Serveur démarré sur http://localhost:8080/album/damso")
	//crée le serveur
	http.ListenAndServe(":8080", r)

}
