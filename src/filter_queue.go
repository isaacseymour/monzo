package crawler

import (
	"errors"
	"fmt"
	"net/url"
	"sync"
	"time"
)

type void struct{}

var voidM void

type filterQueue struct {
	mutex          *sync.Mutex
	executionFn    func(*filterQueue, string)
	baseUrl        *url.URL
	inProgressUrls map[string]struct{}
	seenUrls       map[string]struct{}
}

func NewFilterQueue(baseUrl *url.URL, executionFn func(*filterQueue, string)) *filterQueue {
	return &filterQueue{
		mutex:          &sync.Mutex{},
		executionFn:    executionFn,
		baseUrl:        baseUrl,
		inProgressUrls: make(map[string]struct{}),
		seenUrls:       make(map[string]struct{}),
	}
}

func (q *filterQueue) Add(urlStr string) {
	url, err := url.Parse(urlStr)
	if err != nil {
		return
	}
	url = q.baseUrl.ResolveReference(url)

	if url.Hostname() != q.baseUrl.Hostname() {
		return
	}

	url.Fragment = ""
	urlStr = url.String()

	q.mutex.Lock()
	defer q.mutex.Unlock()

	_, seen := q.seenUrls[urlStr]
	if seen {
		return
	}

	q.inProgressUrls[urlStr] = voidM
	q.seenUrls[urlStr] = voidM

	go func() {
		q.executionFn(q, urlStr)
	}()
}

func (q *filterQueue) Done(urlStr string) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	delete(q.inProgressUrls, urlStr)
}

func (q *filterQueue) Len() int {
	return len(q.inProgressUrls)
}

func (q *filterQueue) WaitForEmpty(timeout time.Duration) error {
	channel := make(chan error, 1)
	go func() {
		time.Sleep(timeout)
		channel <- errors.New("timeout")
	}()
	go func() {
		for q.Len() > 0 {
			time.Sleep(time.Second)
			fmt.Printf("\ndone %d, waiting on %d\n", len(q.seenUrls), q.Len())
		}
		channel <- nil
	}()

	return <-channel
}
