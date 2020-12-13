package response

import (
	"encoding/json"
	"net/http"
	"time"
)

type (
	//Response is a struct for rest api response
	Response struct {
		Header HeaderResponse `json:"header"`
		Data   interface{}    `json:"data,omitempty"`

		startTime time.Time
	}

	//HeaderResponse is a struct for rest header api
	HeaderResponse struct {
		ResponseTime float64     `json:"response_time"`
		ErrorMessage interface{} `json:"error,omitempty"`
	}
)

//NewRestResponse is to create new rest response
func NewRestResponse(startTime time.Time) Response {
	return Response{
		startTime: startTime,
	}
}

//WriteResponse to write rest regular response
func (res *Response) WriteResponse(w http.ResponseWriter, msg interface{}) {
	res.Header.ResponseTime = time.Since(res.startTime).Seconds()
	res.Data = msg

	encoded, _ := json.Marshal(res)

	w.WriteHeader(http.StatusOK)
	w.Write(encoded)
}

//WriteError to write rest error response
func (res *Response) WriteError(w http.ResponseWriter, code int, msg interface{}) {
	res.Header.ResponseTime = time.Since(res.startTime).Seconds()
	res.Header.ErrorMessage = msg

	encoded, _ := json.Marshal(res)

	w.WriteHeader(code)
	w.Write(encoded)
}
