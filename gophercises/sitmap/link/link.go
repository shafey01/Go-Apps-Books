package link

import (
	"io"
	"log"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
}

func Parse(r io.Reader) ([]Link, error) {
	// file, err := os.Open("html.html")
	// if err != nil {
	// 	log.Fatal(err)

	// }
	// defer file.Close()
	// s, err := io.ReadAll(file)
	// if err != nil {

	// 	log.Fatal(err)

	// }
	linkCh := make(chan string)
	doc, err := html.Parse(r)
	if err != nil {

		log.Fatal(err)
	}
	go getHref(doc, linkCh)
	var links []Link

	for l := range linkCh {
		link := Link{
			Href: l,
		}
		links = append(links, link)
	}
	return links, nil
}

func getHref(n *html.Node, lc chan string) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				lc <- a.Val
				break
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getHref(c, lc)
	}
	if n.Parent == nil {
		close(lc)
	}
}
