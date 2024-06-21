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
	//TODO Light tent etc..
	RoomType  string
	Substrate []string
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
