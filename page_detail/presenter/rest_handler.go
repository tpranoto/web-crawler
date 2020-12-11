package presenter

import (
	"net/http"

	"github.com/tpranoto/gochallenge/page_detail/usecase"
)

type (
	//RestHandler contains all handler for rest
	RestHandler interface {
		HandlerGetPageSpecificDetail(w http.ResponseWriter, r *http.Request)
	}
	restHandler struct {
		pageDetail usecase.PageDetail
	}
)

//NewRestHandler init all usecases to be presented
func NewRestHandler(pDetail usecase.PageDetail) RestHandler {
	return &restHandler{
		pageDetail: pDetail,
	}
}

func (rh *restHandler) HandlerGetPageSpecificDetail(w http.ResponseWriter, r *http.Request) {

}
