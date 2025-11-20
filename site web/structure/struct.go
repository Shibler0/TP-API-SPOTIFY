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
