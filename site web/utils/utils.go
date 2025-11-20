package utils

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"siteweb/structure"
	"time"
)

func RenderTemplate(w http.ResponseWriter, filename string, data interface{}) {
	tmpl := template.Must(template.ParseFiles("template/" + filename))
	tmpl.Execute(w, data)
}

func GetArtistPicture() string {
	token := "BQD431hkWP8DaibYvlivX_9lrm_UXOb4Thv_heiNRNn6hjJS8CjDGKxHZp3jQisNral4TPU9SWMuqqm1W9aUychquvbiHlLL0la9ZygEZzKLC3-AWsDpor20zNsmC-pHwHm3zSJqt94"
	artistURL := "https://api.spotify.com/v1/artists/3IW7ScrzXmPvZhB27hmfgy"

	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, errReq := http.NewRequest(http.MethodGet, artistURL, nil)
	if errReq != nil {
		fmt.Println("une erreur est survenue :", errReq.Error())
	}

	req.Header.Add("User-Agent", "Ynov campus B1")
	req.Header.Set("Authorization", "Bearer "+token)

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

	fmt.Println("reach")

	return decodeData.Images[0].Url
}
