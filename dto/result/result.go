package result_dto

type SuccessResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
