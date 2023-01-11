package structs

type user struct {
	Name string `json:"screen_name"`
}

type ResultInfo struct {
	Id string `json:"idstr"`
	User user `json:"user"`
	Rootid string `json:"rootidstr"`
	TotalNumber int `json:"total_number"`
	Textraw string `json:"text_raw"`
}

type Response struct {
	Data     []ResultInfo `json:"data"`
	MaxId int    `json:"max_id"`
}
