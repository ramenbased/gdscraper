package main

// Temp
type TempWeeks struct {
	w []TempWeek
}

type TempWeek struct {
	WeekType string
	Link     string
}

// Main

type Results struct {
	Main []_Main
}

type _Main struct {
	ID           string //TODO: data type int
	Environment  string
	WateringType string
	URL          string
}

// unsure
type Breeder struct {
	Name   string
	Strain string
}

type Soil struct {
	ID         int
	Type       string
	Percentage int
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
