package responses

type DatabaseResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
