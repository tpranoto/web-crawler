package presenter

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	ucMock "github.com/tpranoto/gochallenge/mock/page_detail/usecase"
	"github.com/tpranoto/gochallenge/page_detail/model"
	"github.com/tpranoto/gochallenge/page_detail/usecase"

	"github.com/golang/mock/gomock"
)

func Test_restHandler_HandlerGetPageSpecificDetail(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUC := ucMock.NewMockPageDetail(ctrl)

	type fields struct {
		pageDetail usecase.PageDetail
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		setup  func()
	}{
		{
			name: "valid",
			fields: fields{
				pageDetail: mockUC,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: &http.Request{
					Form: url.Values{
						"url": []string{"test.com"},
					},
				},
			},
			setup: func() {
				mockUC.EXPECT().GetPageDetail(gomock.Any(), "test.com")
			},
		},
		{
			name: "error",
			fields: fields{
				pageDetail: mockUC,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: &http.Request{
					Form: url.Values{
						"url": []string{"test.com"},
					},
				},
			},
			setup: func() {
				mockUC.EXPECT().GetPageDetail(gomock.Any(), "test.com").Return(model.PageDetailsData{}, errors.New("error"))
			},
		},
		{
			name: "no input",
			fields: fields{
				pageDetail: mockUC,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: &http.Request{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rh := &restHandler{
				pageDetail: tt.fields.pageDetail,
			}
			if tt.setup != nil {
				tt.setup()
			}
			rh.HandlerGetPageSpecificDetail(tt.args.w, tt.args.r)
		})
	}
}
