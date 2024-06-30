package main

import (
	"fmt"
	"gdscraper/data"
	"strings"

	"golang.org/x/net/html"
)

/*
func params(doc *html.Node, id string, tbl *data.Tables) {
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "statistic_box active" {
					fmt.Printf("%v ---> %v \n",
						n.FirstChild.NextSibling.NextSibling.FirstChild.Data,
						n.FirstChild.NextSibling.FirstChild.Data)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
*/

func ferts(doc *html.Node, id string, tbl *data.Tables) {
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "fert_item" {
					var f = new(data.Fertilizer)
					name := n.FirstChild.NextSibling.FirstChild.FirstChild.Data
					amount := n.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.Data
					f.AddFert(id, name, amount, tbl)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
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
		//params(doc, id, tbl)
		ferts(doc, id, tbl)
	case "1":
		fmt.Println("Bloom..")
		//params(doc, id, tbl)
		ferts(doc, id, tbl)
	case "2":
		fmt.Println("Harvest")
	}

}
