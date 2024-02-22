package model

type Rate struct {
	Id int
	Score float64
}

type Student struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	GroupId string `json:"group"`
}
