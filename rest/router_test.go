package rest

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	mPresenter "github.com/tpranoto/gochallenge/mock/page_detail/presenter"
	"github.com/tpranoto/gochallenge/page_detail/presenter"
)

func TestRouterPresenter_AssignPaths(t *testing.T) {
	ctrl := gomock.NewController(t)
	pDtMock := mPresenter.NewMockRestHandler(ctrl)

	type fields struct {
		Router  *http.ServeMux
		PDetail presenter.RestHandler
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "valid",
			fields: fields{
				Router:  &http.ServeMux{},
				PDetail: pDtMock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RouterPresenter{
				Router:  tt.fields.Router,
				PDetail: tt.fields.PDetail,
			}
			r.AssignPaths()
		})
	}
}
