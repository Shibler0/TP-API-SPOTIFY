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

type Token struct {
	Token string `json:"access_token"`
}
