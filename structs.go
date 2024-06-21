package main

type Strain struct {
	Diaries []Diary
}

type Diary struct {
	Items          ReportItems
	GerminationDay Germination
	HarvestDay     Harvest
	Weeks          []Week
}

type Germination struct {
	//fill
}

type Harvest struct {
	//fill
}

type ReportItems struct {
	Breed      string
	Strain     string
	LightVeg   string
	LightBloom string
	RoomType   string
	//TODO: Training
}

type Week struct {
	WeekNo        int
	Type          string //Veg or Bloom
	Height        int
	LightSchedule int
	Fertilizers   []Fert
}

type Fert struct {
	Name   string
	Amount float64
}
