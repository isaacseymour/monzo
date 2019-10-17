package main

import (
	"errors"
	"net/url"
	"sync"
	"time"
)

type void struct{}

var voidM void

type filterQueue struct {
	mutex          *sync.Mutex
	executionFn    func(*filterQueue, string)
	hostname       string
	inProgressUrls map[string]struct{}
	seenUrls       map[string]struct{}
}

func NewFilterQueue(hostname string, executionFn func(*filterQueue, string)) *filterQueue {
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
		}
		channel <- nil
	}()

	return <-channel
}
