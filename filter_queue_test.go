package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIgnoreOutsideUrls(t *testing.T) {
	q := NewFilterQueue("gocardless.com")
	q.Add("https://gocardless.com/")
	assert.Equal(t, q.Len(), 1)

	q.Add("https://www.gocardless.com/")
	assert.Equal(t, q.Len(), 1)

	q.Add("https://gocardless.com/thing1")
	assert.Equal(t, q.Len(), 2)

	q.Add("https://gocardless.com/thing2")
	assert.Equal(t, q.Len(), 3)

	q.Add("https://gocardless.com/thing2")
	assert.Equal(t, q.Len(), 3)

	assert.Equal(t, q.Dequeue(), "https://gocardless.com/")
	assert.Equal(t, q.Dequeue(), "https://gocardless.com/thing1")
	assert.Equal(t, q.Dequeue(), "https://gocardless.com/thing2")
	assert.Equal(t, q.Dequeue(), "")

	assert.Equal(t, q.Len(), 0)
}
