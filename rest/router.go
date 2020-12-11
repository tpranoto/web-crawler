package rest

import "net/http"

type (
	//RouterPresenter holds all Presenter for all modules
	RouterPresenter struct {
	}
)

//AssignPaths to put all paths in rest apis
func (r RouterPresenter) AssignPaths(router *http.ServeMux) {

}
