package model

type Diary struct {
	Id          string `json:"id"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	CreateTime  string `json:"create_time"`
}
