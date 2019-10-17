package crawler

import "testing"
import "github.com/stretchr/testify/assert"

func TestSitemap(t *testing.T) {
	s := NewSitemap()
	s.Add("https://gocardless.com", []string{
		"https://gocardless.com/page1",
		"https://gocardless.com/page2",
		"https://gocardless.com/page3",
		"https://gocardless.com/",
	})
	s.Add("https://gocardless.com/page1", []string{
		"https://support.gocardless.com/thing",
		"https://gocardless.com/page3",
	})
	s.Add("https://gocardless.com/page2", []string{
		"https://youtube.com/gocardless",
		"https://gocardless.com/",
	})
	s.Add("https://gocardless.com/page3", []string{
		"https://twitter.com/gocardless",
		"https://gocardless.com/page1",
	})

	assert.Equal(t, `## https://gocardless.com
- https://gocardless.com/page1
- https://gocardless.com/page2
- https://gocardless.com/page3
- https://gocardless.com/

## https://gocardless.com/page1
- https://support.gocardless.com/thing
- https://gocardless.com/page3

## https://gocardless.com/page2
- https://youtube.com/gocardless
- https://gocardless.com/

## https://gocardless.com/page3
- https://twitter.com/gocardless
- https://gocardless.com/page1

`, s.String())
}
