package presenter

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/tpranoto/gochallenge/common/response"
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
	ctx := context.Background()
	writer := response.NewRestResponse(time.Now())
	//take url from query param
	url := r.FormValue("url")

	res, err := rh.pageDetail.GetPageDetail(ctx, url)
	if err != nil {
		writer.WriteError(w, http.StatusInternalServerError, err.Error())
		log.Printf("[HandlerGetPageSpecificDetail]%s", err.Error())
		return
	}

	writer.WriteResponse(w, res)
}
