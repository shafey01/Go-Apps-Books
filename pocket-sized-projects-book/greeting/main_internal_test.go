package main

import (
	"testing"
)

func TestGreet(t *testing.T) {

	type testCases struct {
		want string
		lang language
	}

	var tests = map[string]testCases{

		"English": {
			lang: "en",
			want: "Hello world",
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde",
		},
		"Akkadian, not supported": {
			lang: "akk",
			want: `unsupported language: "akk"`,
		},
		"Greek": {
			lang: "el",
			want: "Χαίρετε Κόσμε",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			got := greet(tc.lang)
			if got != tc.want {
				t.Errorf("excepted %q got %q", tc.want, got)
			}

		})
	}

}
