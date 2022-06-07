package dto

type ResponseBody struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Error      interface{} `json:"error"`
	Data       interface{} `json:"data"`
}

func (responseBody *ResponseBody) SetStatus(status string) *ResponseBody {
	responseBody.Status = status
	return responseBody
}

func (responseBody *ResponseBody) SetStatusCode(statusCode int) *ResponseBody {
	responseBody.StatusCode = statusCode
	return responseBody
}

func (responseBody *ResponseBody) SetError(err interface{}) *ResponseBody {
	responseBody.Error = err
	return responseBody
}

func (responseBody *ResponseBody) SetData(data interface{}) *ResponseBody {
	responseBody.Data = data
	return responseBody
}

func (responseBody ResponseBody) GetResponseBody() ResponseBody {
	return responseBody
}
