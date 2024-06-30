package data

import "fmt"

// --- Temp DB
type Tables struct {
	TblDiary   []Diary
	TblBreeder []Breeder
	TblSoil    []Soil
	TblWeek    []Week
}

// --- Main
type Diary struct {
	ID          string //filled - TODO: not string maybe FOR ALL!!
	Environment string //filled
	URL         string //filled
}

func (d *Diary) AddDiary(ID string, URL string, roomType string, tbl *Tables) {
	d.ID = ID
	d.URL = URL
	d.Environment = roomType
	tbl.TblDiary = append(tbl.TblDiary, *d)
	fmt.Printf("addDiary --> d.ID: %v d.Environment: %v d.URL: %v\n", d.ID, d.Environment, d.URL)

}

type Soil struct {
	ID         string
	Type       string
	Percentage float64
}

func (s *Soil) AddSoil(ID string, soils []string, tbl *Tables) {
	for _, soil := range soils {
		s.ID = ID
		s.Type = soil
		//s.Percentage = Percentage
		tbl.TblSoil = append(tbl.TblSoil, *s)
		fmt.Printf("addSoil --> s.ID: %v s.Type: %v\n", s.ID, s.Type)
	}
}

// TODO: Breeder/Strain how to data structure lmfao
type Breeder struct {
	Name   string
	Strain string
}

type Week struct {
	ID       int
	Week     int
	Type     string
	Height   int     //float??
	Temp     int     //float??
	Humidity int     //float??
	Water    float64 //Litres per 24h
	URL      string
}

type Fertilizer struct {
	ID     int
	WeekNo int
	Name   string
	Amount float64 //xx.x ml/L float??
}

type Nutrients struct {
	Name string
	N    float64
	P    float64
	K    float64
	Cal  float64
	Mag  float64
}
