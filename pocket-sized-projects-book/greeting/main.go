package main

import (
	"flag"
	"fmt"
)

type language string

var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",    // Greek
	"en": "Hello world",      // English
	"fr": "Bonjour le monde", // French
	"he": "שלום עולם",        // Hebrew
	"ur": "ہیلو دنیا",        // Urdu

}

func main() {

	// lang := flag.String("lang", "en", "The required language")

	var lang string
	flag.StringVar(&lang, "lang", "en", "The required language")

	flag.Parse()

	fmt.Println(greet(language(lang)))
}

func greet(l language) string {

	greeting, ok := phrasebook[l]

	if !ok {
		return fmt.Sprintf("This language %s dosen't exsit", l)

	}

	return greeting

}
