package usecase

import (
	"context"
	"testing"

	repoMock "github.com/tpranoto/gochallenge/mock/page_detail/repo"
	"github.com/tpranoto/gochallenge/page_detail/model"
	"github.com/tpranoto/gochallenge/page_detail/repo"

	"github.com/golang/mock/gomock"
)

func Test_pageDetail_GetPageDetail(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repoMock.NewMockPageContent(ctrl)

	type fields struct {
		pageRepo repo.PageContent
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
