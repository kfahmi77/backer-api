package helper

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ApiResponse(message string, code int, status string, data interface{}) Response {
	m := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	r := Response{
		Meta: m,
		Data: data,
	}
	return r
}
