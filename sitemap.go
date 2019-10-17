package main

import (
	"bytes"
	"sync"
)

type sitemapEntry struct {
	pageUrl string
	links   []string
}
type sitemap struct {
	mutex   *sync.Mutex
	entries []sitemapEntry
}

func NewSitemap() *sitemap {
	return &sitemap{
		mutex:   &sync.Mutex{},
		entries: make([]sitemapEntry, 0),
	}
}

func (s *sitemap) Add(pageUrl string, links []string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.entries = append(s.entries, sitemapEntry{pageUrl: pageUrl, links: links})
}

func (s *sitemap) String() string {
	result := bytes.Buffer{}

	for _, entry := range s.entries {
		result.WriteString("## ")
		result.WriteString(entry.pageUrl)
		result.WriteString("\n")
		for _, link := range entry.links {
			result.WriteString("- ")
			result.WriteString(link)
			result.WriteString("\n")
		}
		result.WriteString("\n")
	}

	return result.String()
}
