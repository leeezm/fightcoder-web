package baseController

type HttpResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type PagingResponse struct {
	RequestPage int         `json:"requestPage"`
	TotalPages  int         `json:"totalPages"`
	Data        interface{} `json:"data"`
}

func (this *Base) Success(data ...interface{}) *HttpResponse {
	resp := &HttpResponse{Code: 0}
	if len(data) > 0 {
		resp.Data = data[0]
	}
	return resp
}

func (this *Base) Fail(msg ...string) *HttpResponse {
	resp := &HttpResponse{Code: 1}
	if len(msg) > 0 {
		resp.Msg = msg[0]
	}
	return resp
}

// old
func (this *Base) MakeResponseSuccess(data ...interface{}) map[string]interface{} {
	if len(data) < 1 {
		return this.MakeResponse("success", "")
	}
	return this.MakeResponse("success", data[0])
}

func (this *Base) MakeResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{"code": 0, "msg": msg, "data": data}
}
