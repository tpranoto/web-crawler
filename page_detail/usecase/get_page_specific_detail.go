package usecase

import (
	"context"

	"github.com/tpranoto/gochallenge/page_detail/repo"
)

type (
	//PageDetail contains all usecases for page details
	PageDetail interface {
		GetPageDetail(ctx context.Context, url string)
	}
	pageDetail struct {
		pageRepo repo.PageContent
	}
)

//NewPageDetailUsecase init all deps in page detail
func NewPageDetailUsecase(pageContentRepo repo.PageContent) PageDetail {
	return &pageDetail{
		pageRepo: pageContentRepo,
	}
}

func (p *pageDetail) GetPageDetail(ctx context.Context, url string) {
	return
}
