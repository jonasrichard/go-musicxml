package gomusicxml

import "encoding/xml"

type ScorePartwise struct {
	XMLName xml.Name `xml:"score-partwise"`
	Version string   `xml:"version,attr"`
	Part    []Part   `xml:"part"`
}

type Part struct {
	ID      string    `xml:"id,attr"`
	Measure []Measure `xml:"measure"`
}

type Measure struct {
	Number string `xml:"number,attr"`
	Note   []Note `xml:"note"`
}

type Note struct {
	Pitch    Pitch     `xml:"pitch"`
	Duration int       `xml:"duration"`
	Type     string    `xml:"type"`
	Rest     *struct{} `xml:"rest,omitempty"`
}

type Pitch struct {
	Step   string `xml:"step"`
	Octave int    `xml:"octave"`
}
