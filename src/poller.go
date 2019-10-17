package crawler

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type crawler struct {
	client  *http.Client
	sitemap *sitemap
}

func (c *crawler) Poll(q *filterQueue, url string) {
	fmt.Print(".")
	defer q.Done(url)

	resp, err := c.client.Get(url)
	if err != nil {
		fmt.Errorf("Got err %s\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return
	}

	links, err := FindLinks(resp.Body)
	if err != nil {
		return
	}

	c.sitemap.Add(url, links)
	for _, link := range links {
		q.Add(link)
	}
}

func (c *crawler) Run(baseUrl string) (*sitemap, error) {
	url, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	queue := NewFilterQueue(url, c.Poll)
	queue.Add(baseUrl)
	err = queue.WaitForEmpty(120 * time.Second)
	return c.sitemap, err
}
