package model

type Obj struct {
	Articles []News `json: "articles"`
}

type News struct {
	Title string `json:"title"`
	Description string `json:"description"`
}