package utils

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"siteweb/structure"
	"strings"
	"time"
)

func RenderTemplate(w http.ResponseWriter, filename string, data interface{}) {
	tmpl := template.Must(template.ParseFiles("template/" + filename))
	tmpl.Execute(w, data)
}

func GetToken() {
	Client_ID := "fd1e0a45f53640d6961f96714278c8d9"
	Client_SECRET := "25efd36f7192478c81d8087f7386ece2"

	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", Client_ID)
	data.Set("client_secret", Client_SECRET)

	req, errReq := http.NewRequest(http.MethodPost, "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if errReq != nil {
		fmt.Println("une erreur est survenue :", errReq.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, errResp := httpClient.Do(req)

	if errResp != nil {
		fmt.Println("errreur ", errResp)
		return
	}
	defer res.Body.Close()

	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("erreur ", errBody)
		return
	}

	var decodeToken structure.Token

	json.Unmarshal(body, &decodeToken)

	fmt.Println(decodeToken.Token)

	structure.TOKEN = decodeToken.Token
}

func GetArtistPicture() string {
	artistURL := "https://api.spotify.com/v1/artists/3IW7ScrzXmPvZhB27hmfgy"

	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, errReq := http.NewRequest(http.MethodGet, artistURL, nil)
	if errReq != nil {
		fmt.Println("une erreur est survenue :", errReq.Error())
	}

	req.Header.Add("User-Agent", "Ynov campus B1")
	req.Header.Set("Authorization", "Bearer "+structure.TOKEN)

	res, errResp := httpClient.Do(req)

	if errResp != nil {
		fmt.Println("une erreur est survenue : ", errResp.Error())
		return ""
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, errBody := io.ReadAll(res.Body)

	if errBody != nil {
		fmt.Println("une erreur est survenue,", errBody.Error())
	}

	var decodeData structure.Format

	json.Unmarshal(body, &decodeData)

	return decodeData.Images[0].Url
}

func GetAlbums() structure.AlbumsInfos {
	url := "https://api.spotify.com/v1/artists/2UwqpfQtNuhBwviIC0f2ie/albums"

	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, errReq := http.NewRequest(http.MethodGet, url, nil)
	if errReq != nil {
		fmt.Println("une erreur est survenue :", errReq.Error())
	}

	req.Header.Add("User-Agent", "Ynov campus B1")
	req.Header.Set("Authorization", "Bearer "+structure.TOKEN)

	res, errResp := httpClient.Do(req)

	if errResp != nil {
		fmt.Println("une erreur est survenue : ", errResp.Error())
		return structure.AlbumsInfos{}
	}

	defer res.Body.Close()

	body, errBody := io.ReadAll(res.Body)

	if errBody != nil {
		fmt.Println("une erreur est survenue,", errBody.Error())
		return structure.AlbumsInfos{}
	}

	var decodeData structure.AlbumsInfos

	json.Unmarshal(body, &decodeData)
	//Exemple
	for i := range decodeData.Albums {
		fmt.Println("Nom :", decodeData.Albums[i].Name)
		fmt.Println("Release date :", decodeData.Albums[i].Releasedate)
		fmt.Println("Total tracks :", decodeData.Albums[i].Totaltrack)
	}

	for i := range decodeData.Albums {
		if len(decodeData.Albums[i].Images) > 0 {
			decodeData.Albums[i].ImageURL = decodeData.Albums[i].Images[0].Url
		} else {
			decodeData.Albums[i].ImageURL = "/static/img/damso.jpeg"
		}
	}

	return decodeData
}
