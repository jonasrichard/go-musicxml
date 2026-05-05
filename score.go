package gomusicxml

import "encoding/xml"

type ScorePartwise struct {
	XMLName  xml.Name `xml:"score-partwise"`
	Version  string   `xml:"version,attr"`
	PartList PartList `xml:"part-list"`
	Part     []Part   `xml:"part"`
}

type PartList struct {
	ScorePart []struct {
		ID   string `xml:"id,attr"`
		Name string `xml:"part-name"`
	} `xml:"score-part"`
}

type Part struct {
	ID      string    `xml:"id,attr"`
	Measure []Measure `xml:"measure"`
}

type Measure struct {
	Number     string     `xml:"number,attr"`
	Attributes Attributes `xml:"attributes"`
	Note       []Note     `xml:"note"`
}

type Attributes struct {
	Divisions int `xml:"divisions"`
	Key       struct {
		Fifths int `xml:"fifths"`
	} `xml:"key"`
	Time struct {
		Beats    int `xml:"beats"`
		BeatType int `xml:"beat-type"`
	} `xml:"time"`
	Clef []struct {
		Sign string `xml:"sign"`
		Line int    `xml:"line"`
	} `xml:"clef"`
}
