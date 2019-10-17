package main

import (
	"net/http"
	"net/url"
	"time"
)

type crawler struct {
	client  *http.Client
	sitemap *sitemap
}

func (c *crawler) Poll(q *filterQueue, url string) {
	defer q.Done(url)

	resp, err := c.client.Get(url)
	if err != nil {
		return
	}

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

	queue := NewFilterQueue(url.Hostname(), c.Poll)
	queue.Add(baseUrl)
	err = queue.WaitForEmpty(30 * time.Second)
	return c.sitemap, err
}
