package main

import (
	"net/url"
	"sync"
)

type void struct{}

var voidM void

type callback func()

type filterQueue struct {
	mutex          *sync.Mutex
	executionFn    func(string, callback)
	hostname       string
	inProgressUrls map[string]struct{}
	seenUrls       map[string]struct{}
}

func NewFilterQueue(hostname string, executionFn func(string, callback)) *filterQueue {
	return &filterQueue{
		mutex:          &sync.Mutex{},
		executionFn:    executionFn,
		hostname:       hostname,
		inProgressUrls: make(map[string]struct{}),
		seenUrls:       make(map[string]struct{}),
	}
}

func (q *filterQueue) Add(urlStr string) {
	url, err := url.Parse(urlStr)
	if err != nil {
		return
	}

	if url.Hostname() != q.hostname {
		return
	}

	q.mutex.Lock()
	defer q.mutex.Unlock()

	_, seen := q.seenUrls[urlStr]
	if seen {
		return
	}

	q.inProgressUrls[urlStr] = voidM
	q.seenUrls[urlStr] = voidM

	go func() {
		q.executionFn(urlStr, func() { q.done(urlStr) })
	}()
}

func (q *filterQueue) done(urlStr string) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	delete(q.inProgressUrls, urlStr)
}

func (q *filterQueue) Len() int {
	return len(q.inProgressUrls)
}
