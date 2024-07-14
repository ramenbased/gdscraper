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
	TblHarvest    []Harvest
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
	ID         string
	Type       string
	Percentage string
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
	ID        string
	Week      string
	WType     string
	Height    string
	TempDay   string
	TempNight string
	Humidity  string
	PotSize   string
	Water     string
	PH        string
	LightS    string
	TDS       string
}

func (w *Week) AddWeek(
	id string,
	week string,
	wType string,
	height string,
	tempDay string,
	tempNight string,
	humid string,
	potsize string,
	water string,
	ph string,
	lights string,
	tds string,
	tbl *Tables) {

	w.ID = id
	w.Week = week
	w.Height = height
	w.WType = wType
	w.TempDay = tempDay
	w.TempNight = tempNight
	w.Humidity = humid
	w.PotSize = potsize
	w.Water = water
	w.PH = ph
	w.LightS = lights
	w.TDS = tds
	tbl.TblWeek = append(tbl.TblWeek, *w)
	fmt.Printf("addWeek --> w.ID: %v w.WType: %v w.Week: %v w.Height: %v w.TempDay: %v w.TempNight: %v w.Humidity: %v w.PotSize: %v w.Water: %v w.PH: %v w.LightS: %v w.TDS: %v\n", w.ID, w.WType, w.Week, w.Height, w.TempDay, w.TempNight, w.Humidity, w.PotSize, w.Water, w.PH, w.LightS, w.TDS)
}

type Fertilizer struct {
	ID     string
	WeekID string
	Name   string
	Amount string //TODO: xx.x ml/L and why gallons after login scrape?? float??
}

func (f *Fertilizer) AddFert(id string, wID string, name string, amount string, tbl *Tables) {
	f.ID = id
	f.WeekID = wID
	f.Name = name
	f.Amount = amount
	tbl.TblFertilizer = append(tbl.TblFertilizer, *f)
	fmt.Printf("addFert --> f.ID: %v f.weekID: %v f.Name: %v f.Amount: %v\n", f.ID, f.WeekID, f.Name, f.Amount)
}

type Nutrients struct {
	Name string
	N    float64
	P    float64
	K    float64
	Cal  float64
	Mag  float64
}

type Harvest struct {
	ID           string
	WeekID       string
	WetWeight    string
	DryWeight    string
	AmountPlants string
	GrowRoomSize string
}

func (h *Harvest) AddHarvest(id string, wID string, wetWeight string, dryWeight string, amountPlants string, growRoomSize string, tbl *Tables) {
	h.ID = id
	h.WeekID = wID
	h.WetWeight = wetWeight
	h.DryWeight = dryWeight
	h.AmountPlants = amountPlants
	h.GrowRoomSize = growRoomSize
	tbl.TblHarvest = append(tbl.TblHarvest, *h)
	fmt.Printf("addHarvest --> h.ID: %v h.WeekID: %v h.WetWeight: %v h.DryWeight: %v h.AmountPlants: %v h.GrowRoomSize: %v\n", h.ID, h.WeekID, h.WetWeight, h.DryWeight, h.AmountPlants, h.GrowRoomSize)
}
