package usecase

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/tpranoto/gochallenge/common/config"
	repoMock "github.com/tpranoto/gochallenge/mock/page_detail/repo"
	"github.com/tpranoto/gochallenge/page_detail/model"
	"github.com/tpranoto/gochallenge/page_detail/repo"
)

func Test_pageDetail_GetPageDetail(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repoMock.NewMockPageContent(ctrl)

	type fields struct {
		pageRepo repo.PageContent
		cfg      config.Config
	}
	type args struct {
		ctx   context.Context
		input string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult model.PageDetailsData
		wantErr    bool
		setup      func()
	}{
		{
			name: "valid",
			fields: fields{
				pageRepo: mockRepo,
				cfg: config.Config{
					Worker: config.WorkerConfig{
						Default: 10,
					},
				},
			},
			args: args{
				ctx:   context.Background(),
				input: "https://test.com",
			},
			setup: func() {
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pageDetail{
				pageRepo: tt.fields.pageRepo,
			}
			tt.setup()
			p.GetPageDetail(tt.args.ctx, tt.args.input)
		})
	}
}

func Test_pageDetail_separateLinks(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repoMock.NewMockPageContent(ctrl)

	type fields struct {
		pageRepo repo.PageContent
		cfg      config.Config
	}
	type args struct {
		ctx      context.Context
		urlInput string
		links    []string
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantInternal     model.LinkDetail
		wantExternal     model.LinkDetail
		wantInaccessible model.LinkDetailWithError
		setup            func()
	}{
		{
			name: "valid with multiple links from internal, external, inaccessible, and using cache",
			fields: fields{
				pageRepo: mockRepo,
				cfg: config.Config{
					Worker: config.WorkerConfig{
						Default: 1,
					},
				},
			},
			args: args{
				ctx:      context.Background(),
				urlInput: "https://base-url.com",
				links: []string{
					"#",
					"#",
					"https://link2.com",
					"https://link2.com",
					"/link3.com",
					"/link3.com",
					"/fail.com",
					"/fail.com",
				},
			},
			wantInternal: model.LinkDetail{
				Count: 4,
				Data:  []string{"#", "#", "/link3.com", "/link3.com"},
			},
			wantExternal: model.LinkDetail{
				Count: 2,
				Data:  []string{"https://link2.com", "https://link2.com"},
			},
			wantInaccessible: model.LinkDetailWithError{
				Count: 2,
				Data:  []string{"/fail.com", "/fail.com"},
				Errors: []string{
					"status 404 for /fail.com",
				},
			},
			setup: func() {
				mockRepo.EXPECT().GetContentFromURL(gomock.Any(), "https://link2.com").Return(&http.Response{
					StatusCode: http.StatusOK,
				}, nil)

				mockRepo.EXPECT().GetContentFromURL(gomock.Any(), "https://base-url.com/link3.com").Return(&http.Response{
					StatusCode: http.StatusOK,
				}, nil)

				mockRepo.EXPECT().GetContentFromURL(gomock.Any(), "https://base-url.com/fail.com").Return(&http.Response{
					StatusCode: http.StatusNotFound,
				}, nil)
			},
		},
		{
			name: "error in getting content",
			fields: fields{
				pageRepo: mockRepo,
				cfg: config.Config{
					Worker: config.WorkerConfig{
						Default: 1,
					},
				},
			},
			args: args{
				ctx:      context.Background(),
				urlInput: "https://base-url.com",
				links: []string{
					"/fail.com",
				},
			},
			wantInternal: model.LinkDetail{},
			wantExternal: model.LinkDetail{},
			wantInaccessible: model.LinkDetailWithError{
				Count: 1,
				Data:  []string{"/fail.com"},
				Errors: []string{
					"failed to get /fail.com, error",
				},
			},
			setup: func() {
				mockRepo.EXPECT().GetContentFromURL(gomock.Any(), "https://base-url.com/fail.com").Return(&http.Response{}, errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pageDetail{
				pageRepo: tt.fields.pageRepo,
				cfg:      tt.fields.cfg,
			}
			if tt.setup != nil {
				tt.setup()
			}

			gotInternal, gotExternal, gotInaccessible := p.separateLinks(tt.args.ctx, tt.args.urlInput, tt.args.links)
			if !reflect.DeepEqual(gotInternal, tt.wantInternal) {
				t.Errorf("pageDetail.separateLinks() gotInternal = %v, want %v", gotInternal, tt.wantInternal)
			}
			if !reflect.DeepEqual(gotExternal, tt.wantExternal) {
				t.Errorf("pageDetail.separateLinks() gotExternal = %v, want %v", gotExternal, tt.wantExternal)
			}
			if !reflect.DeepEqual(gotInaccessible, tt.wantInaccessible) {
				t.Errorf("pageDetail.separateLinks() gotInaccessible = %v, want %v", gotInaccessible, tt.wantInaccessible)
			}
		})
	}
}
