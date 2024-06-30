package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func params(doc *html.Node) {
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
func ferts(doc *html.Node) {
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "fert_item" {
					fmt.Printf("%v ---> %v \n",
						n.FirstChild.NextSibling.FirstChild.FirstChild.Data,
						n.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.Data)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}

func compileWeek(weekHTML string, w TempWeek) {
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
		params(doc)
		ferts(doc)
	case "1":
		fmt.Println("Bloom..")
		params(doc)
		ferts(doc)
	case "2":
		fmt.Println("Harvest")
	}

}
