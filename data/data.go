package data

import (
	"fmt"
)

// --- Temp DB
type Tables struct {
	TblDiary      []Diary
	TblBreeder    []Breeder
	TblWeek       []Week
	TblSoil       []Soil
	TblFertilizer []Fertilizer
}

// --- Main
type Diary struct {
	ID          string //filled - TODO: not string maybe FOR ALL!!
	Environment string //filled
	URL         string //filled
}

func (d *Diary) AddDiary(id string, URL string, roomType string, tbl *Tables) {
	d.ID = id
	d.URL = URL
	d.Environment = roomType
	tbl.TblDiary = append(tbl.TblDiary, *d)
	fmt.Printf("addDiary --> d.ID: %v d.Environment: %v d.URL: %v\n", d.ID, d.Environment, d.URL)

}

type Soil struct {
	ID         string //filled
	Type       string //filled
	Percentage string //filled TODO: float?
}

func (s *Soil) AddSoil(id string, soil string, percent string, tbl *Tables) {
	s.ID = id
	s.Type = soil
	s.Percentage = percent
	tbl.TblSoil = append(tbl.TblSoil, *s)
	fmt.Printf("addSoil --> s.ID: %v s.Type: %v s.Percentage: %v\n", s.ID, s.Type, s.Percentage)
}

// TODO: Breeder/Strain how to data structure lmfao
type Breeder struct {
	Name   string
	Strain string
}

type Week struct {
	ID       string
	Week     string //gets it from params(), good idea?
	WType    string //get from list urls?!
	Height   string
	Temp     string
	Humidity string
	URL      string
}

func (w *Week) AddWeek(id string, week string, wType string, height string, tbl *Tables) {
	w.ID = id
	w.Week = week
	w.Height = height
	w.WType = wType
	tbl.TblWeek = append(tbl.TblWeek, *w)
	fmt.Printf("addWeek --> w.ID: %v w.WType: %v w.Week: %v w.Height: %v\n", w.ID, w.WType, w.Week, w.Height)
}

type Fertilizer struct {
	ID     string //filled
	WeekNo int    //TODO: not filled bro
	Name   string //filled
	Amount string //filled TODO: xx.x ml/L and why gallons after login scrape?? float??
}

func (f *Fertilizer) AddFert(id string, name string, amount string, tbl *Tables) {
	f.ID = id
	f.Name = name
	f.Amount = amount
	tbl.TblFertilizer = append(tbl.TblFertilizer, *f)
	fmt.Printf("addFert --> f.ID: %v f.Name: %v f.Amount: %v\n", f.ID, f.Name, f.Amount)
}

type Nutrients struct {
	Name string
	N    float64
	P    float64
	K    float64
	Cal  float64
	Mag  float64
}
