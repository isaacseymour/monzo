package crawler

import (
	"fmt"
	"net/http"
	"time"
)

func Crawl(baseUrl string) (fmt.Stringer, error) {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		},
	}
	c := &crawler{client: client, sitemap: NewSitemap()}
	sitemap, err := c.Run(baseUrl)
	return sitemap, err
}
