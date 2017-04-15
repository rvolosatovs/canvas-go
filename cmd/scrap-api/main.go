package main

import (
	"fmt"
	"os"

	"github.com/anaskhan96/soup"
)

func main() {
	resp, err := soup.Get("https://canvas.instructure.com/doc/api/courses.html")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	endpoints := doc.FindAll("h3", "class", "endpoint")
	for _, endpoint := range endpoints {
		fmt.Print(endpoint.Text())
	}
}
