package main

// --- Temp
type TempWeeks struct {
	w      []TempWeek
	sanity bool //is handled in seperate func
}

type TempWeek struct {
	WeekType string
	Link     string
}
