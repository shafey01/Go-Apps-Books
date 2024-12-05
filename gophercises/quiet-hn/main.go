package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/shafey01/Go-Apps-Books/gophercises/quiet-hn/hn"
)

var (
	cache = map[int]hn.Item{}
	lock  sync.RWMutex
)

func main() {
	// parse flags
	var port, numStories int
	flag.IntVar(&port, "port", 3000, "the port to start the web server on")
	flag.IntVar(&numStories, "num_stories", 30, "the number of top stories to display")
	flag.Parse()

	tpl := template.Must(template.ParseFiles("./index.gohtml"))

	http.HandleFunc("/", handler(numStories, tpl))

	// Start the server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handler(numStories int, tpl *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var client hn.Client
		ids, err := client.TopItems()
		if err != nil {
			http.Error(w, "Failed to load top stories", http.StatusInternalServerError)
			return
		}

		ids = ids[:int(float64(numStories)*1.25)]
		var stories []sortedItems
		storiesCh := make(chan sortedItems)
		var wg sync.WaitGroup

		for i, id := range ids {
			wg.Add(1)
			go func(id, i int) {
				defer wg.Done()

				// Check if item in the cache and if not writing it...
				if _, ok := cacheRead(id); !ok {
					hnItem, err := client.GetItem(id)
					if err != nil {
						return
					}
					cacheWrite(hnItem, id)
					// cache[id] = hnItem
				}

				// Reading from cache
				hnItem, _ := cacheRead(id)
				item := parseHNItem(hnItem)
				if isStoryLink(item) {
					storiesCh <- sortedItems{item: item, idx: i}
					// if len(storiesCh) >= numStories {
					// 	log.Println("Break the channal")
					// 	return
					// }
				}
			}(id, i)
		}

		go func() {
			wg.Wait()

			close(storiesCh)
		}()

		for storisfromCh := range storiesCh {
			stories = append(stories, storisfromCh)

			// log.Printf("Append %v ", storisfromCh.Item)
		}
		sort.Slice(stories, func(i, j int) bool {
			return stories[i].idx < stories[j].idx
		})

		data := templateData{
			Stories: stories[:numStories],
			Time:    time.Now().Sub(start),
		}
		err = tpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Failed to process the template", http.StatusInternalServerError)
			return
		}
	})
}

func isStoryLink(item item) bool {
	return item.Type == "story" && item.URL != ""
}

func parseHNItem(hnItem hn.Item) item {
	ret := item{Item: hnItem}
	url, err := url.Parse(ret.URL)
	if err == nil {
		ret.Host = strings.TrimPrefix(url.Hostname(), "www.")
	}
	return ret
}

// item is the same as the hn.Item, but adds the Host field
type item struct {
	hn.Item
	Host string
}

type templateData struct {
	Stories []sortedItems
	Time    time.Duration
}

type sortedItems struct {
	item
	idx int
}

func cacheRead(id int) (hn.Item, bool) {
	lock.RLock()
	defer lock.RUnlock()
	hnItem, ok := cache[id]
	return hnItem, ok
}

func cacheWrite(hn hn.Item, id int) {
	lock.Lock()
	defer lock.Unlock()
	cache[id] = hn
}
