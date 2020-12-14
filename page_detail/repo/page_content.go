package repo

import (
	"context"
	"net/http"

	panics "github.com/tpranoto/gochallenge/common/panic_handler"

	"github.com/pkg/errors"
)

type (
	//PageContent contains all repositories for page content
	PageContent interface {
		GetContentFromURL(ctx context.Context, url string) (data *http.Response, err error)
	}
	pageContent struct{}
)

//NewPageContentRepo init all repositories in page content
func NewPageContentRepo() PageContent {
	return &pageContent{}
}

func (p *pageContent) GetContentFromURL(ctx context.Context, url string) (resp *http.Response, err error) {
	defer panics.HandlePanic("[GetContentFromURL]")

	resp, err = http.Get(url)
	if err != nil {
		err = errors.Wrapf(err, "[GetContentFromURL] failed to get from %s", url)
		return
	}

	return
}
