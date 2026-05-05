package main

import (
	"encoding/xml"
	"fmt"
	"os"

	musicxml "github.com/jonasrichard/go-musicxml"
)

func main2() {
	buf, err := os.ReadFile("/Users/richardjonas/projects/music/_resources/A and A7 arpeggio forms.xml")
	if err != nil {
		panic(err)
	}

	var score musicxml.ScorePartwise
	if err := xml.Unmarshal(buf, &score); err != nil {
		panic(err)
	}

	part := score.Part[0]
	measure := part.Measure[0]

	for _, note := range measure.Note {
		fmt.Printf("%s%d\n", note.Pitch.Step, note.Pitch.Octave)

		for _, notation := range note.Notations.Items {
			fmt.Printf("notation: %T\n", notation)
			switch n := notation.(type) {
			case musicxml.Technical:
				fmt.Printf("technical: fret=%d string=%d\n", *n.Fret, *n.String)
				if n.Extra != "" {
					fmt.Printf("guitar pro data: %s\n", n.Extra)
				}
			default:
				fmt.Printf("unknown notation: %T\n", notation)
			}
		}
	}
}

func main() {
	var note musicxml.Note
	note.Pitch.Step = "A"
	note.Pitch.Octave = 4
	note.Duration = 1
	note.Type = "quarter"
	note.Notations = musicxml.Notations{
		Items: []musicxml.Notation{
			musicxml.Technical{
				Fret:   new(5),
				String: new(4),
				HammerOn: &musicxml.HammerOn{
					Number: 1,
					Text:   "H",
					Type:   "start",
				},
				Extra: `<?GP <root><string>6</string><fret>12</fret></root> ?>`,
			},
		},
	}

	buf, err := xml.MarshalIndent(note, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}
