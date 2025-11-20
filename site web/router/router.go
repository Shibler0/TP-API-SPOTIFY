package router

import (
	"net/http"
	"siteweb/controller"
)

func New() *http.ServeMux {

	mux := http.NewServeMux() // Création d'un nouveau ServeMux, qui est un routeur simple pour les requêtes HTTP

	// On associe les chemins URL à des fonctions spécifiques du controller
	mux.HandleFunc("/home", controller.Home)
	mux.HandleFunc("/album/damso", controller.Damso)

	//gere le css
	fileServer := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	return mux // On retourne le routeur configuré
}
