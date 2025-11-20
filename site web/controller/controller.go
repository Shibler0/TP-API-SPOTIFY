package controller

import (
	"net/http"
	"siteweb/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "home.html", nil)
}
