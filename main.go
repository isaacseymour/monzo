package main

import (
	"./src"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please pass the base url as the first arg")
		os.Exit(1)
	}
	baseUrl := os.Args[1]

	if baseUrl == "" {
		fmt.Println("Please pass the base url as the first arg")
		os.Exit(1)
	}

	sitemap, err := crawler.Crawl(baseUrl)

	if err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println(sitemap)
}
