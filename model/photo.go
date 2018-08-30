package model

type Photo struct {
	Id          string `json:"id"`
	PhotoName   string `json:"photo_name"`
	CreateTime  string `json:"create_time"`
	Creater     string `json:"creater"`
	Pictures    []Picture `gorm:"ForeignKey:PhotoId"`
}

type Picture struct {
	Id       string `json:"id"`
	PicName  string `json:"pic_name"`
	PicSize  int64  `json:"pic_size"`
	PicUrl   string `json:"pic_url"`
	PhotoId  string `json:"photo_id"`
}
