package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

var urls = []string{
	"https://gocardless.com/",
	"https://www.gocardless.com/",
	"https://gocardless.com/thing1",
	"https://gocardless.com/thing2",
}

func TestIgnoreOutsideUrls(t *testing.T) {
	fmt.Println("hello!")

	mutex := &sync.Mutex{}
	calledWith := make([]string, 0)

	executionFn := func(url string, cb callback) {
		fmt.Println("executed")
		mutex.Lock()
		calledWith = append(calledWith, url)
		mutex.Unlock()
		cb()
	}

	q := NewFilterQueue("gocardless.com", executionFn)
	fmt.Println("built")

	for _, url := range urls {
		fmt.Println("adding")
		q.Add(url)
		q.Add(url)
	}
	q.Add("https://other-domain.com/path")
	q.Add("https://support.gocardless.com/thing")

	fmt.Println("sleeping")
	time.Sleep(100)

	assert.Equal(t, calledWith, urls)

	for k, _ := range q.inProgressUrls {
		fmt.Printf("waiting on %s\n", k)
	}

	assert.Equal(t, q.Len(), 0)
}
