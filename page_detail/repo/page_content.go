package repo

import "context"

type (
	//PageContent contains all repositories for page content
	PageContent interface {
		GetContentFromURL(ctx context.Context, url string) (data []byte, err error)
	}
	pageContent struct {
	}
)

//NewPageContentRepo init all repositories in page content
func NewPageContentRepo() PageContent {
	return &pageContent{}
}

func (p *pageContent) GetContentFromURL(ctx context.Context, url string) (data []byte, err error) {

	return
}
