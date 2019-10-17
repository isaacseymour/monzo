package crawler

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

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

	q := NewFilterQueue("gocardless.com", executionFn)

	for _, url := range urls {
		q.Add(url)
		q.Add(url) // shouldn't execute twice
	}
	q.Add("https://www.gocardless.com/")
	q.Add("https://other-domain.com/path")
	q.Add("https://support.gocardless.com/thing")

	assert.Nil(t, q.WaitForEmpty(2*time.Second))

	assert.ElementsMatch(t, urls, calledWith)

	assert.Equal(t, q.Len(), 0)
}
