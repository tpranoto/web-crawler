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

	"github.com/tpranoto/gochallenge/common/config"
	"github.com/tpranoto/gochallenge/page_detail/model"
	"github.com/tpranoto/gochallenge/page_detail/repo"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

type (
	//PageDetail contains all usecases for page details
	PageDetail interface {
		GetPageDetail(ctx context.Context, url string) (result model.PageDetailsData, err error)
	}
	pageDetail struct {
		pageRepo repo.PageContent
		cfg      config.Config
	}
)

//NewPageDetailUsecase init all deps in page detail
func NewPageDetailUsecase(pageContentRepo repo.PageContent, cfg config.Config) PageDetail {
	return &pageDetail{
		pageRepo: pageContentRepo,
		cfg:      cfg,
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

		//get all links in the page
		links := []string{}
		//get all links in the page
		doc.Find("a").Each(func(idx int, selection *goquery.Selection) {
			href, ok := selection.Attr("href")
			if !ok {
				return
			}
			links = append(links, href)
		})
		//process the links
		result.Links.Internal, result.Links.External, result.Links.Inaccessible = p.separateLinks(ctx, input, links)
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

	return
}

func (p *pageDetail) separateLinks(ctx context.Context, urlInput string, links []string) (internal model.LinkDetail, external model.LinkDetail, inaccessible model.LinkDetailWithError) {
	//save same links into a map acting as a local cache
	localCache := map[string]model.CacheValue{}

	//set up waitgroup,channel, mutexes for concurrent
	var wg sync.WaitGroup
	var mutexCache, mutexInternal, mutexExternal, mutexInaccessible sync.Mutex
	channel := make(chan string)

	//limit concurrent to 10 worker
	for worker := 0; worker < p.cfg.Worker.Default; worker++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//get all links from channel
			for link := range channel {
				//get from cache if available
				mutexCache.Lock()
				cache, ok := localCache[link]
				mutexCache.Unlock()
				if ok {
					if cache.Reachable {
						if cache.Internal {
							mutexInternal.Lock()
							internal.Data = append(internal.Data, link)
							mutexInternal.Unlock()
							continue
						}
						mutexExternal.Lock()
						external.Data = append(external.Data, link)
						mutexExternal.Unlock()
						continue
					}
					mutexInaccessible.Lock()
					inaccessible.Data = append(inaccessible.Data, link)
					mutexInaccessible.Unlock()
					continue
				}

				var isInternal, accessible bool
				href := link
				//prepare internal links
				if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
					//if internal links prefix != / than count it as internal accessible link
					//ex: tel:, mailto:, #section, javascript:
					if !strings.HasPrefix(link, "/") {
						mutexInternal.Lock()
						internal.Data = append(internal.Data, link)
						mutexInternal.Unlock()
						isInternal = true
						accessible = true

						//save link details to map cache
						mutexCache.Lock()
						localCache[link] = model.CacheValue{
							Internal:  isInternal,
							Reachable: accessible,
						}
						mutexCache.Unlock()
						continue
					}
					//prepare internal links
					//internal links
					parsedURL, _ := url.Parse(urlInput)
					href = fmt.Sprintf("%s://%s%s", parsedURL.Scheme, parsedURL.Host, link)
					isInternal = true
				}

				//get status code from the links
				linkRes, err := p.pageRepo.GetContentFromURL(ctx, href)
				if err != nil {
					mutexInaccessible.Lock()
					inaccessible.Data = append(inaccessible.Data, link)
					inaccessible.Errors = append(inaccessible.Errors, fmt.Sprintf("failed to get %s, %s", link, err.Error()))
					mutexInaccessible.Unlock()
					log.Printf("[GetPageDetail]%s", err.Error())

					//save link details to map cache
					mutexCache.Lock()
					localCache[link] = model.CacheValue{
						Internal:  isInternal,
						Reachable: accessible,
					}
					mutexCache.Unlock()
					continue
				}

				//if response status != 200 && 999(crawled) -> counts as inaccessible
				if linkRes.StatusCode != http.StatusOK && linkRes.StatusCode != 999 {
					mutexInaccessible.Lock()
					inaccessible.Data = append(inaccessible.Data, link)
					inaccessible.Errors = append(inaccessible.Errors, fmt.Sprintf("status %d for %s", linkRes.StatusCode, link))
					mutexInaccessible.Unlock()

					//save link details to map cache
					mutexCache.Lock()
					localCache[link] = model.CacheValue{
						Internal:  isInternal,
						Reachable: accessible,
					}
					mutexCache.Unlock()
					continue
				}

				accessible = true
				//save to internal data / external data
				if isInternal {
					mutexInternal.Lock()
					internal.Data = append(internal.Data, link)
					mutexInternal.Unlock()
				} else {
					mutexExternal.Lock()
					external.Data = append(external.Data, link)
					mutexExternal.Unlock()
				}

				//save link details to map cache
				mutexCache.Lock()
				localCache[link] = model.CacheValue{
					Internal:  isInternal,
					Reachable: accessible,
				}
				mutexCache.Unlock()
			}
		}()

	}
	for _, link := range links {
		channel <- link
	}

	close(channel)
	wg.Wait()

	//count links
	internal.Count = len(internal.Data)
	external.Count = len(external.Data)
	inaccessible.Count = len(inaccessible.Data)

	return
}

func findHTMLVersion(source string) string {
	reg := regexp.MustCompile(`<(?i)!doctype (.*?)>`)
	return reg.FindString(source)
}
