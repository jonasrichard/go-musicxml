package gomusicxml

import "encoding/xml"

type ScorePartwise struct {
	XMLName  xml.Name `xml:"score-partwise"`
	Version  string   `xml:"version,attr"`
	PartList PartList `xml:"part-list"`
	Part     []Part   `xml:"part"`
}

type PartList struct {
	ScorePart []ScorePart `xml:"score-part"`
}

type ScorePart struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"part-name"`
}

type Part struct {
	ID      string    `xml:"id,attr"`
	Measure []Measure `xml:"measure"`
}

// NewScorePartwise creates a new ScorePartwise with default values.
func NewScorePartwise() ScorePartwise {
	return ScorePartwise{
		Version: "4.0",
	}
}

// AddPart adds a new part to the score with the given ID and name.
func (s *ScorePartwise) AddPart(id, name string) *Part {
	s.PartList.ScorePart = append(s.PartList.ScorePart, ScorePart{
		ID:   id,
		Name: name,
	})
	s.Part = append(s.Part, Part{ID: id})

	return &s.Part[len(s.Part)-1]
}

func (p *Part) AddMeasure(measure *Measure) {
	p.Measure = append(p.Measure, *measure)
}
