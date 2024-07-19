package main

import (
	"fmt"
	"gdscraper/data"
	"strings"

	"golang.org/x/net/html"
)

func week(doc *html.Node, id string, tbl *data.Tables) *data.Week {
	var w = new(data.Week)
	//w fills up in loop then gives to rw.AddWeek()
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "statistic_box active" {
					key := n.FirstChild.NextSibling.NextSibling.FirstChild.Data
					val := n.FirstChild.NextSibling.FirstChild.Data
					switch key {
					case "Height":
						w.Height = val
					case "Vegetation":
						w.WType = key
						w.Week = val
					case "Flowering":
						w.WType = key
						w.Week = val
					case "Day Air Temperature":
						w.TempDay = val
					case "Night Air Temperature":
						w.TempNight = val
					case "Air Humidity":
						w.Humidity = val
					case "Pot Size":
						w.PotSize = val
					case "Watering Volume Per Plant Per 24h":
						w.Water = val
					case "pH":
						w.PH = val
					case "Light Schedule":
						w.LightS = val
					case "TDS":
						w.TDS = val
					}
				}
				if a.Key == "class" && a.Val == "method" {
					val := n.LastChild.FirstChild.LastChild.Data
					switch val {
					case "LST":
						w.LST = true
					case "HST":
						w.HST = true
					case "SoG":
						w.SoG = true
					case "ScrOG":
						w.ScrOG = true
					case "Topping":
						w.Topping = true
					case "FIMing":
						w.FIMing = true
					case "Main-Lining":
						w.MainLining = true
					case "Defoliation":
						w.Defoliation = true
					case "12-12 From Seed":
						w.FromSeed1212 = true
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	var rw = new(data.Week)
	rw.AddWeek(id, w.Week, w.WType, w.Height, w.TempDay, w.TempNight, w.Humidity, w.PotSize, w.Water, w.PH, w.LightS, w.TDS, w.LST, w.HST, w.SoG, w.ScrOG, w.Topping, w.FIMing, w.MainLining, w.Defoliation, w.FromSeed1212, tbl)
	return rw
}

func ferts(doc *html.Node, id string, tbl *data.Tables, weekID string) {
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "fert_item" {
					var f = new(data.Fertilizer)
					name := n.FirstChild.NextSibling.FirstChild.FirstChild.Data
					amount := n.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.Data
					f.AddFert(id, weekID, name, amount, tbl)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}

func harvest(doc *html.Node, id string, tbl *data.Tables) *data.Harvest {
	var h = new(data.Harvest)
	//w fills up in loop then gives to rw.AddWeek()
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "statistic_box active" {

					key := n.FirstChild.NextSibling.NextSibling.FirstChild.Data
					val := n.FirstChild.NextSibling.FirstChild.Data

					switch key {
					case "Harvest":
						h.WeekID = val
					case "Bud wet weight":
						h.WetWeight = val
					case "Bud dry weight":
						h.DryWeight = val
					case "Number of plants harvested":
						h.AmountPlants = val
					case "Grow Room size":
						h.GrowRoomSize = val
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	var rh = new(data.Harvest)
	rh.AddHarvest(id, h.WeekID, h.WetWeight, h.DryWeight, h.AmountPlants, h.GrowRoomSize, tbl)
	return rh
}

func compileDiaryWeek(weekHTML string, id string, w TempWeek, tbl *data.Tables) {
	sr := strings.NewReader(weekHTML)
	doc, err := html.Parse(sr)
	if err != nil {
		panic(err)
	}
	switch w.WeekType {
	case "-1":
		fmt.Println("Germination..")
	case "0":
		fmt.Println("Veg..")
		//dislike return
		wID := week(doc, id, tbl)
		ferts(doc, id, tbl, wID.Week)
	case "1":
		fmt.Println("Bloom..")
		wID := week(doc, id, tbl)
		ferts(doc, id, tbl, wID.Week)
	case "2":
		fmt.Println("Harvest..")
		harvest(doc, id, tbl)
	}

}
