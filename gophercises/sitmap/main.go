package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	Link "github.com/shafey01/Go-Apps-Books/gophercises/sitmap/link"
)

func main() {

	pageURL := flag.String("url", "", "URL of the page")
	depth := flag.Int("depth", 1, "Depth of siteMap")
	xmlFileName := flag.String("xml", "sitemap.xml", "file name to store sitemap")
	flag.Parse()
	if *pageURL == "" {
		log.Fatal("Missing -url flag")
	}

	sitemap, err := buildSiteMap(*pageURL, *depth)

	if err != nil {
		log.Fatal(err)

	}
	// fmt.Println("Length of links", len(sitemap))
	// for _, l := range sitemap {
	// 	fmt.Println(l)
	// }

	if err := createSiteMap(sitemap, *xmlFileName); err != nil {
		log.Fatalf("Error in creating siteMap in %s: %v ", *xmlFileName, err)
	}
	log.Printf("Generated sitemap with %d link(s) for %s in %s",
		len(sitemap), *pageURL, *xmlFileName)
}

func getURLS(pageURL string) ([]string, error) {

	pageURL = strings.TrimSuffix(pageURL, "/")
	res, err := http.Get(pageURL)
	if err != nil {
		log.Fatal(err)

	}
	defer res.Body.Close()

	getURLS, err := Link.Parse(res.Body)
	if err != nil {
		log.Fatal(err)

	}
	var allURL []string
	for _, l := range getURLS {
		allURL = append(allURL, l.Href)
	}

	fmt.Println("Length of allURL", len(allURL))
	var urlsofDomine []string
	var counter = 0
	for _, url := range allURL {
		if i := strings.Compare(url, pageURL); i == 0 {
			counter++
			fmt.Println("Counter: ", counter)
		}
		if strings.HasPrefix(url, "http") && !strings.HasPrefix(url, pageURL) {
			continue
		}

		if strings.HasPrefix(url, pageURL) {

			urlsofDomine = append(urlsofDomine, url)
			continue
		}

		if strings.Contains(url, ":") {
			continue
		}

		if i := strings.Index(url, "#"); i != -1 {
			url = url[:i]
		}

		if url == "" || url[0] != '/' {
			url = "/" + url

		}

		url = pageURL + url
		urlsofDomine = append(urlsofDomine, url)
	}

	return urlsofDomine, nil
}

func buildSiteMap(pageURL string, depth int) ([]string, error) {
	urlsMap := map[string]bool{}
	urls := []string{pageURL}

	for i := 0; i < depth; i++ {
		var allURLs []string

		for _, url := range urls {
			suburls, err := getURLS(url)
			if err != nil {
				log.Fatal(err)

			}
			var uniURLs []string
			for _, suburl := range suburls {
				if urlsMap[suburl] {
					continue
				}
				urlsMap[suburl] = true
				uniURLs = append(uniURLs, suburl)
			}

			allURLs = append(allURLs, uniURLs...)

		}
		urls = allURLs
	}
	return urls, nil
}

type siteMapXML struct {
	XMLName xml.Name        `xml:"urlset"`
	Xmlns   string          `xml:"xmlns,attr"`
	URLS    []siteMapXMlURL `xml:"url"`
}

type siteMapXMlURL struct {
	Loc string `xml:"loc"`
}

func createSiteMap(urls []string, pathXML string) error {
	var siteMap siteMapXML
	siteMap.Xmlns = "https://www.sitemaps.org/schemas/sitemap/0.9"
	for _, url := range urls {
		siteMap.URLS = append(siteMap.URLS, siteMapXMlURL{
			Loc: url,
		})
	}
	siteMapIndent, err := xml.MarshalIndent(&siteMap, "", "\t")
	if err != nil {

		log.Fatal(err)
	}
	xmlData := []byte(xml.Header + string(siteMapIndent))
	return os.WriteFile(pathXML, xmlData, os.ModePerm)
}
