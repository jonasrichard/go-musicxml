package gomusicxml

type Measure struct {
	Number     int        `xml:"number,attr"`
	Attributes Attributes `xml:"attributes"`
	Note       []Note     `xml:"note"`
}

type Attributes struct {
	Divisions    int            `xml:"divisions"`
	Key          Key            `xml:"key"`
	Time         Time           `xml:"time"`
	Clef         []Clef         `xml:"clef"`
	StaffDetails []StaffDetails `xml:"staff-details"`
	Transpose    []Transpose    `xml:"transpose"`
}

type Key struct {
	Fifths int    `xml:"fifths"`
	Mode   string `xml:"mode"`
}

type Time struct {
	Beats    int `xml:"beats"`
	BeatType int `xml:"beat-type"`
}

type Clef struct {
	Sign string `xml:"sign"`
	Line int    `xml:"line"`
}

type StaffDetails struct {
	StaffLines  int           `xml:"staff-lines"`
	StaffTuning []StaffTuning `xml:"staff-tuning"`
}

type StaffTuning struct {
	Line         int    `xml:"line,attr"`
	TuningStep   string `xml:"tuning-step"`
	TuningOctave int    `xml:"tuning-octave"`
}

type Transpose struct {
	Number       int `xml:"number,attr"`
	Diatonic     int `xml:"diatonic"`
	Chromatic    int `xml:"chromatic"`
	OctaveChange int `xml:"octave-change"`
}

func NewMeasure(number int) Measure {
	return Measure{
		Number: number,
	}
}

func (m *Measure) WithAttributes(attributes Attributes) {
	m.Attributes = attributes
}

func NewAttributes(divisions int) Attributes {
	return Attributes{
		Divisions: divisions,
	}
}

func (a *Attributes) WithKey(fifths int, mode string) {
	a.Key = Key{
		Fifths: fifths,
		Mode:   mode,
	}
}

func (a *Attributes) WithTime(beats, beatType int) {
	a.Time = Time{
		Beats:    beats,
		BeatType: beatType,
	}
}

func (a *Attributes) AddClef(sign string, line int) {
	a.Clef = append(a.Clef, Clef{
		Sign: sign,
		Line: line,
	})
}

func (a *Attributes) AddStaffDetails(staffLines int, tunings []StaffTuning) {
	a.StaffDetails = append(a.StaffDetails, StaffDetails{
		StaffLines:  staffLines,
		StaffTuning: tunings,
	})
}

func (a *Attributes) AddTranspose(number, diatonic, chromatic, octaveChange int) {
	a.Transpose = append(a.Transpose, Transpose{
		Number:       number,
		Diatonic:     diatonic,
		Chromatic:    chromatic,
		OctaveChange: octaveChange,
	})
}
