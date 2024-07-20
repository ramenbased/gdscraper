package data

import (
	"log"
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
	log.Printf("addDiary --> d.ID: %v d.Environment: %v d.URL: %v\n", c_NoSpace(d.ID), c_NoSpace(d.Environment), c_NoSpace(d.URL))
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
	log.Printf("addSoil --> s.ID: %v s.Type: %v s.Percentage: %v\n", c_NoSpace(s.ID), c_NoSpace(s.Type), c_WeekInt(s.Percentage))
}

// TODO: Breeder/Strain how to data structure lmfao
type Breeder struct {
	Name   string
	Strain string
}

type Week struct {
	ID        string
	Week      string //int
	WType     string //string
	Height    string //float64
	TempDay   string //float64
	TempNight string //float64
	Humidity  string //float64
	PotSize   string //float64
	Water     string //float64
	PH        string //float64
	LightS    string //int
	TDS       string //float64

	//methods		 //bool
	LST          bool
	HST          bool
	SoG          bool
	ScrOG        bool
	Topping      bool
	FIMing       bool
	MainLining   bool
	Defoliation  bool
	FromSeed1212 bool
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
	//methods
	lst bool,
	hst bool,
	sog bool,
	scrog bool,
	topping bool,
	fiming bool,
	mainlining bool,
	defoliation bool,
	fromseed1212 bool,
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
	//methods
	w.LST = lst
	w.HST = hst
	w.SoG = sog
	w.ScrOG = scrog
	w.Topping = topping
	w.FIMing = fiming
	w.MainLining = mainlining
	w.Defoliation = defoliation
	w.FromSeed1212 = fromseed1212
	tbl.TblWeek = append(tbl.TblWeek, *w)
	log.Printf("addWeek --> w.ID: %v w.WType: %v w.Week: %v w.Height: %v w.TempDay: %v w.TempNight: %v w.Humidity: %v w.PotSize: %v w.Water: %v w.PH: %v w.LightS: %v w.TDS: %v\n",
		w.ID,
		c_NoSpace(w.WType),
		c_WeekInt(w.Week),
		c_StringFloat(w.Height),
		c_StringFloat(w.TempDay),
		c_StringFloat(w.TempNight),
		c_WeekInt(w.Humidity),
		c_StringFloat(w.PotSize),
		c_StringFloat(w.Water),
		c_StringFloat(w.PH),
		c_WeekInt(w.LightS),
		c_StringFloat(w.TDS))
	log.Printf("addWeek Methods --> w.LST: %v w.HST: %v w.SoG: %v w.ScrOG: %v w.Topping: %v w.FIMing: %v w.MainLining: %v w.Defoliation: %v w.FromSeed1212: %v\n", w.LST, w.HST, w.SoG, w.ScrOG, w.Topping, w.FIMing, w.MainLining, w.Defoliation, w.FromSeed1212)
}

type Fertilizer struct {
	ID     string
	WeekID string
	Name   string
	Amount string //TODO: xx.x ml/L and why gallons after login scrape?? float??
	Href   string
}

func (f *Fertilizer) AddFert(id string, wID string, name string, amount string, href string, tbl *Tables) {
	f.ID = id
	f.WeekID = wID
	f.Name = name
	f.Amount = amount
	f.Href = href
	tbl.TblFertilizer = append(tbl.TblFertilizer, *f)
	log.Printf("addFert --> f.ID: %v f.weekID: %v f.Name: %v f.Amount: %v f.Href: %v\n", c_NoSpace(f.ID), c_WeekInt(f.WeekID), f.Name, c_FertAmount(f.Amount), f.Href)
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
	log.Printf("addHarvest --> h.ID: %v h.WeekID: %v h.WetWeight: %v h.DryWeight: %v h.AmountPlants: %v h.GrowRoomSize: %v\n", c_NoSpace(h.ID), c_WeekInt(h.WeekID), c_StringFloat(h.WetWeight), c_StringFloat(h.DryWeight), c_WeekInt(h.AmountPlants), c_StringFloat(h.GrowRoomSize))
}
