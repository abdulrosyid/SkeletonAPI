package helper

const (
	ErrorDataNotFound = "data %s tidak ditemukan"
	ErrorAccountBalance = "saldo tidak mencukupi"
)

type ResponseDetail struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseDetailOutput(success bool, code int, message string, data interface{}) ResponseDetail {
	res := ResponseDetail{
		Success: success,
		Code:    code,
		Message: message,
		Data:    data,
	}
	return res
}
