package crawler

import (
	"fmt"
	"net/http"
)

func Crawl(baseUrl string) (fmt.Stringer, error) {
	client := &http.Client{}
	c := &crawler{client: client, sitemap: NewSitemap()}
	sitemap, err := c.Run(baseUrl)
	return sitemap, err
}
