package structure

type Home struct {
	Damso string
	Jul   string
}

type Format struct {
	Images []struct {
		Url    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"images"`
}

type AlbumsInfos struct {
	Albums []struct {
		Totaltrack int `json:"total_tracks"`
		Images     []struct {
			Url string `json:"url"`
		} `json:"images"`
		Name        string `json:"name"`
		Releasedate string `json:"release_date"`
		ImageURL    string
	} `json:"items"`
}

type Laylow struct {
	Album struct {
		Name        string `json:"name"`
		ReleaseDate string `json:"release_date"`

		Images []struct {
			URL string `json:"url"`
		} `json:"images"`

		ExternalURLs struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`

		Artists []struct {
			Name string `json:"name"`
		} `json:"artists"`

		CoverUrl   string
		ArtistName string
	} `json:"album"`

	Name string `json:"name"`
}

type Token struct {
	Token string `json:"access_token"`
}
