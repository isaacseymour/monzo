# Monzo crawler

Code is split into a 4 areas:

- `parser.go` handles pulling out the links from a streamed HTML response
- `filter_queue.go` handles restricting crawling only to the base URL, and tracking work done/left
  to do
- `sitemap` is the data structure for building up the sitemap, and then printing it out
- `poller` handles the HTTP of fetching, and parsing each page, and pushing it into the queue &
  sitemap

Finally `crawler.go` presents the main interface: pass in a base URL string, and get back a sitemap.
`main.go` deals with the command-line input.

Run it with:
```bash
go run main.go $BASE_URL
```

Some weird things I think could be improved:
- While the queue handles relative paths & removing fragments, the sitemap doesn't
- I wasn't totally sure whether to include external links in the sitemap (I left them in)
- The way the polling & queue interact they kinda know too much about each other. I tried a few ways
  of cutting this but couldn't find one that felt totally right.
- There's a lotta mutexing. I think this could maybe become neater with channels?
