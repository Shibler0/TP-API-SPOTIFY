package controller

import (
	"net/http"
	"siteweb/structure"
	"siteweb/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {

	if structure.TOKEN == "" {
		utils.GetToken()
		utils.GetAlbums()
	}

	data := structure.Home{
		Damso: "../static/img/damso.jpeg",
		Jul:   utils.GetArtistPicture(),
	}

	utils.RenderTemplate(w, "home.html", data)
}

func Damso(w http.ResponseWriter, r *http.Request) {

	utils.GetToken()

	instance := utils.GetAlbums()

	data := structure.AlbumsInfos{
		Albums: instance.Albums,
	}

	utils.RenderTemplate(w, "damso.html", data)
}
