package models

type Body struct {
	Title string `json:"title"`
	ID    int    `json:"id"`
}

type Input struct {
	Num int `json:"num"`
}

type Output struct {
	Numb int `json:"num"`
}
