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

	pictures := utils.GetArtistPictures()

	data := structure.Home{
		Damso: pictures[1],
		Jul:   pictures[0],
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

func Laylow(w http.ResponseWriter, r *http.Request) {

	utils.GetToken()

	instance := utils.GetLaylowTrack()

	data := structure.Laylow{
		Album: instance.Album,
	}

	utils.RenderTemplate(w, "laylow.html", data)

}
