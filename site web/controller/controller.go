package controller

import (
	"net/http"
	"siteweb/structure"
	"siteweb/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {

	data := structure.Home{
		Damso: "../static/img/damso.jpeg",
	}

	utils.RenderTemplate(w, "home.html", data)
}

func Damso(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "damso.html", nil)
}
