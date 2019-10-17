package crawler

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"testing"
)

type RoundTripFunc func(req *http.Request) (*http.Response, error)

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func successRoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.String() != "https://gocardless.com" {
		return &http.Response{
			StatusCode: 404,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte{})),
		}, nil
	}

	f, err := os.Open("../fixtures/home.html")

	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bufio.NewReader(f)),
	}, nil
}

func TestPoll(t *testing.T) {
	crawler := &crawler{
		client:  &http.Client{Transport: RoundTripFunc(successRoundTrip)},
		sitemap: NewSitemap(),
	}
	url, _ := url.Parse("https://gocardless.com")
	queue := NewFilterQueue(url, crawler.Poll)

	crawler.Poll(queue, "https://gocardless.com")

	assert.Regexp(t, "## https://gocardless.com\n", crawler.sitemap.String())
}

func TestRun(t *testing.T) {
	crawler := &crawler{
		client:  &http.Client{Transport: RoundTripFunc(successRoundTrip)},
		sitemap: NewSitemap(),
	}

	crawler.Run("https://gocardless.com")

	assert.Regexp(t, "## https://gocardless.com\n", crawler.sitemap.String())
}
