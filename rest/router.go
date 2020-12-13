package rest

import (
	"net/http"

	"github.com/tpranoto/gochallenge/page_detail/presenter"
)

type (
	//RouterPresenter holds all presenters for all modules
	RouterPresenter struct {
		Router *http.ServeMux

		//presenters
		PDetail presenter.RestHandler
	}
)

//AssignPaths to put all paths in rest apis
func (r RouterPresenter) AssignPaths() {
	r.Router.HandleFunc("/pagedetail", r.PDetail.HandlerGetPageSpecificDetail)
}
