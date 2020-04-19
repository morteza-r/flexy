package flexy

type RestResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  bool        `json:"status"`
}
