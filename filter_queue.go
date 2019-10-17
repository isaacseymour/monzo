package main

import "net/url"
import "sync"

type void struct{}

var voidM void

type filterQueue struct {
	mutex    *sync.Mutex
	hostname string
	queue    []string
	seenUrls map[string]struct{}
}

func NewFilterQueue(hostname string) *filterQueue {
	return &filterQueue{
		mutex:    &sync.Mutex{},
		hostname: hostname,
		queue:    make([]string, 0),
		seenUrls: make(map[string]struct{}),
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

	q.queue = append(q.queue, urlStr)
	q.seenUrls[urlStr] = voidM
}

func (q *filterQueue) Len() int {
	return len(q.queue)
}

func (q *filterQueue) Dequeue() string {
	var elem string

	q.mutex.Lock()

	if len(q.queue) < 1 {
		return elem
	}

	elem, q.queue = q.queue[0], q.queue[1:]

	q.mutex.Unlock()

	return elem
}
