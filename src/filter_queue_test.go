package crawler

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"sync"
	"testing"
	"time"
)

var baseUrl, _ = url.Parse("https://gocardless.com/")
var urls = []string{
	"https://gocardless.com/",
	"https://gocardless.com/thing1",
	"https://gocardless.com/thing2",
}

func TestIgnoreOutsideUrls(t *testing.T) {

	mutex := &sync.Mutex{}
	calledWith := make([]string, 0)

	executionFn := func(q *filterQueue, url string) {
		mutex.Lock()
		calledWith = append(calledWith, url)
		mutex.Unlock()
		q.Done(url)
	}

	q := NewFilterQueue(baseUrl, executionFn)

	for _, url := range urls {
		q.Add(url)
		q.Add(url) // shouldn't execute twice
	}
	// should be ignored
	q.Add("https://www.gocardless.com/")
	q.Add("https://other-domain.com/path")
	q.Add("https://support.gocardless.com/thing")

	// should be absolute-ified
	q.Add("/thing3")

	assert.Nil(t, q.WaitForEmpty(2*time.Second))

	assert.ElementsMatch(t, append(urls, "https://gocardless.com/thing3"), calledWith)

	assert.Equal(t, q.Len(), 0)
}
