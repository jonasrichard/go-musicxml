package main

import (
	"encoding/xml"
	"fmt"
	gomusicxml "go-musicxml"
	"os"
)

func main() {
	buf, err := os.ReadFile("../_resources/A and A7 arpeggio forms.xml")
	if err != nil {
		panic(err)
	}

	var score gomusicxml.ScorePartwise
	if err := xml.Unmarshal(buf, &score); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", score)
}
