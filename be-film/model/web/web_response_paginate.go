package web

type WebResponsePaginate struct {
	Code        int         `json:"code"`
	Status      string      `json:"status"`
	Data        interface{} `json:"data"`
	TotalPage   int         `json:"total_page"`
	CurrentPage int         `json:"current_page"`
}
