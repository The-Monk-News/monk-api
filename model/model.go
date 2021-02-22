package model

//main
type Obj struct {
	Articles []News `json: "articles"`
}

//news struct
type News struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json: "url"`
	UrlImage    string `json "urlToImage"`
	Author      string `json "author"`
	PublishedAt string `json "publishedAt"`
}
