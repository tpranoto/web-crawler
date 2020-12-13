package usecase

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"

	"github.com/tpranoto/gochallenge/page_detail/model"
	"github.com/tpranoto/gochallenge/page_detail/repo"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

const (
	html5 = "<!doctype html>"
)

type (
	//PageDetail contains all usecases for page details
	PageDetail interface {
		GetPageDetail(ctx context.Context, url string) (result model.PageDetailsData, err error)
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

func (p *pageDetail) GetPageDetail(ctx context.Context, input string) (result model.PageDetailsData, err error) {
	doc, err := goquery.NewDocument(input)
	if err != nil {
		err = errors.Wrap(err, "[GetPageDetail]")
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		html, _ := doc.Html()
		result.Version = findHTMLVersion(html)

		//get title of the page
		doc.Find("title").Each(func(idx int, selection *goquery.Selection) {
			result.Title = selection.Text()
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		//get h1 heading counts
		doc.Find("h1").Each(func(idx int, selection *goquery.Selection) {
			result.HeadingLevels.H1++
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		//get h2 heading counts
		doc.Find("h2").Each(func(idx int, selection *goquery.Selection) {
			result.HeadingLevels.H2++
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		//get h3 heading counts
		doc.Find("h3").Each(func(idx int, selection *goquery.Selection) {
			result.HeadingLevels.H3++
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		//get h4 heading counts
		doc.Find("h4").Each(func(idx int, selection *goquery.Selection) {
			result.HeadingLevels.H4++
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		//get h5 heading counts
		doc.Find("h5").Each(func(idx int, selection *goquery.Selection) {
			result.HeadingLevels.H5++
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		//get h6 heading counts
		doc.Find("h6").Each(func(idx int, selection *goquery.Selection) {
			result.HeadingLevels.H6++
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		//save same links into a map acting as a cache
		cacheLink := map[string]model.CacheValue{}

		//get all links in the page
		doc.Find("a").Each(func(idx int, selection *goquery.Selection) {
			href, ok := selection.Attr("href")
			if !ok {
				return
			}

			//get from cache if available
			if cache, ok := cacheLink[href]; ok {
				if cache.Reachable {
					if cache.Internal {
						result.Links.Internal.Data = append(result.Links.Internal.Data, href)
					} else {
						result.Links.External.Data = append(result.Links.External.Data, href)
					}
					return
				}
				result.Links.Inaccessible.Data = append(result.Links.Inaccessible.Data, href)
				return
			}

			//separate internal and external links
			if strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://") {
				accessible := p.checkExternalLinks(ctx, &result, href)
				cacheLink[href] = model.CacheValue{
					Internal:  false,
					Reachable: accessible,
				}
			} else if strings.HasPrefix(href, "/") {
				accessible := p.checkInternalLinks(ctx, &result, input, href)
				cacheLink[href] = model.CacheValue{
					Internal:  true,
					Reachable: accessible,
				}
			}
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		doc.Find("input").Each(func(idx int, selection *goquery.Selection) {
			attr, ok := selection.Attr("type")
			if !ok {
				return
			}

			//if attr password exists counts login form available
			if attr == "password" {
				result.HasLoginForm = true
			}
		})
	}()

	wg.Wait()

	//count links
	result.Links.Internal.Count = len(result.Links.Internal.Data)
	result.Links.External.Count = len(result.Links.External.Data)
	result.Links.Inaccessible.Count = len(result.Links.Inaccessible.Data)

	return
}

func (p *pageDetail) checkExternalLinks(ctx context.Context, result *model.PageDetailsData, href string) bool {
	linkRes, err := p.pageRepo.GetContentFromURL(ctx, href)
	//if error happen counts as inaccessible
	if err != nil {
		result.Links.Inaccessible.Data = append(result.Links.Inaccessible.Data, href)
		result.Links.Inaccessible.Errors = append(result.Links.Inaccessible.Errors, fmt.Sprintf("failed to get %s, %s", href, err.Error()))
		log.Printf("[GetPageDetail]%s", err.Error())
		return false
	}
	defer linkRes.Body.Close()

	//if response status != 200 && 999(crawled) -> counts as inaccessible
	if linkRes.StatusCode != http.StatusOK && linkRes.StatusCode != 999 {
		result.Links.Inaccessible.Data = append(result.Links.Inaccessible.Data, href)
		result.Links.Inaccessible.Errors = append(result.Links.Inaccessible.Errors, fmt.Sprintf("status %d for %s", linkRes.StatusCode, href))
		return false
	}

	result.Links.External.Data = append(result.Links.External.Data, href)
	return true
}

func (p *pageDetail) checkInternalLinks(ctx context.Context, result *model.PageDetailsData, input, href string) bool {
	parsedURL, err := url.Parse(input)
	if err != nil {
		result.Links.Inaccessible.Data = append(result.Links.Inaccessible.Data, href)
		result.Links.Inaccessible.Errors = append(result.Links.Inaccessible.Errors, fmt.Sprintf("failed to parse %s, %s", input, err.Error()))
		log.Printf("[GetPageDetail]%s", err.Error())
		return false
	}

	//get baseURL + href for internal
	linkRes, err := p.pageRepo.GetContentFromURL(ctx, fmt.Sprintf("%s://%s%s", parsedURL.Scheme, parsedURL.Host, href))
	//if error happen counts as inaccessible
	if err != nil {
		result.Links.Inaccessible.Data = append(result.Links.Inaccessible.Data, href)
		result.Links.Inaccessible.Errors = append(result.Links.Inaccessible.Errors, fmt.Sprintf("failed to get %s, %s", href, err.Error()))
		log.Printf("[GetPageDetail]%s", err.Error())
		return false
	}
	defer linkRes.Body.Close()

	//if response status != 200 && 999(crawled) -> counts as inaccessible
	if linkRes.StatusCode != http.StatusOK && linkRes.StatusCode != 999 {
		result.Links.Inaccessible.Data = append(result.Links.Inaccessible.Data, href)
		result.Links.Inaccessible.Errors = append(result.Links.Inaccessible.Errors, fmt.Sprintf("status %d for %s", linkRes.StatusCode, href))
		return false
	}

	result.Links.Internal.Data = append(result.Links.Internal.Data, href)
	return true
}

func findHTMLVersion(source string) string {
	reg := regexp.MustCompile(`<(?i)!doctype (.*?)>`)
	return reg.FindString(source)
}
