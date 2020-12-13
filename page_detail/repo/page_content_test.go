package repo

import (
	"context"
	"net/http"
	"testing"
)

func Test_pageContent_GetContentFromURL(t *testing.T) {
	type args struct {
		ctx context.Context
		url string
	}
	tests := []struct {
		name     string
		p        *pageContent
		args     args
		wantResp *http.Response
		wantErr  bool
	}{
		{
			name: "fail to get from url",
			args: args{
				ctx: context.Background(),
				url: "test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pageContent{}
			p.GetContentFromURL(tt.args.ctx, tt.args.url)
		})
	}
}
