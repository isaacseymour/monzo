package parser

import (
	h "golang.org/x/net/html"
	"io"
)

func FindLinks(reader io.Reader) ([]string, error) {
	t := h.NewTokenizer(reader)
	links := make([]string, 0)

	for {
		switch t.Next() {
		case h.ErrorToken:
			err := t.Err()
			if err == io.EOF {
				return links, nil
			} else {
				return nil, err
			}

		case h.StartTagToken:
			link := extractLink(t)
			if link != "" {
				links = append(links, link)
			}
		}
	}
}

func extractLink(t *h.Tokenizer) string {
	tagName, hasAttrs := t.TagName()
	if string(tagName) != "a" {
		return ""
	}

	var name, val []byte

	for hasAttrs {
		name, val, hasAttrs = t.TagAttr()

		if string(name) == "href" {
			return string(val[:])
		}
	}

	return ""
}
