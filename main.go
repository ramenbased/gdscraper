package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	cu "github.com/Davincible/chromedp-undetected"
	"github.com/chromedp/chromedp"
	"golang.org/x/net/html"
)

func Er(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func compileItems(itemsHTML string) {
	var sr = strings.NewReader(itemsHTML)
	var doc, err = html.Parse(sr)
	Er(err)

	var f func(*html.Node)
	f = func(n *html.Node) {
		//RoomType and Substrate
		if n.Type == html.ElementNode && n.Data == "div" {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "info" {
					if n.FirstChild != nil && n.FirstChild.Data == "div" {
						switch n.LastChild.FirstChild.Data {
						//TODO: fill Structs, SOIL CAN BE MULTIPLE ENTRIES
						case "Room Type":
							fmt.Printf("Room Type: %v \n", n.FirstChild.FirstChild.Data)
						case "Grow medium":
							fmt.Printf("Grow Medium: %v \n", n.FirstChild.FirstChild.Data)
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}

func compileWeekOverview(weeksHTML string) *TempWeeks {
	var sr = strings.NewReader(weeksHTML)
	var doc, err = html.Parse(sr)
	Er(err)

	var rv = new(TempWeeks)

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "data-faza" {
					for _, a2 := range n.Attr {
						if a2.Key == "href" {
							var w = new(TempWeek)
							w.WeekType = a.Val
							w.Link = a2.Val
							rv.w = append(rv.w, *w)
							//fmt.Printf("Weektype: %v %v \n", a.Val, a2.Val)
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return rv
}

func compileWeek(weekHTML string) {
	var sr = strings.NewReader(weekHTML)
	var doc, err = html.Parse(sr)
	if err != nil {
		panic(err)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				//fertilizers
				if a.Key == "class" && a.Val == "fert_item" {
					for ch := n.FirstChild; ch != nil; ch = ch.NextSibling {
						if ch.Type == html.ElementNode {
							switch ch.Attr[1].Val {
							case "fert_val":
								fmt.Println(ch.FirstChild.Data)
							case "fert_name":
								fmt.Println(ch.FirstChild.LastChild.Data)
							default:
								fmt.Println("nothing found...?")
							}
						}
					}
					fmt.Println("TODO: add to struct")
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}

func getUserDiary(ctx context.Context, URLs []string) {
	var itemsHTML string
	var weeksHTML string
	//iterate over URLs
	for _, reportURL := range URLs {
		fmt.Printf("visiting %v \n", "https://growdiaries.com"+reportURL)

		if err := chromedp.Run(ctx,
			chromedp.Navigate("https://growdiaries.com"+reportURL),
			chromedp.Sleep(3*time.Second),
			chromedp.OuterHTML(".report_items.report_seeds", &itemsHTML),
			chromedp.ActionFunc(func(ctx context.Context) error { compileItems(itemsHTML); return nil }),
			chromedp.OuterHTML(".day_items", &weeksHTML),
			chromedp.ActionFunc(func(ctx context.Context) error {
				weeks := compileWeekOverview(weeksHTML)

				//iterate over weeks
				for _, w := range weeks.w {
					var diaryHTML string
					fmt.Println(w.Link, w.WeekType)
					if err := chromedp.Run(ctx,
						chromedp.Navigate("https://growdiaries.com"+w.Link),
						chromedp.Sleep(5*time.Second),
						chromedp.OuterHTML("body", &diaryHTML),
						chromedp.ActionFunc(func(ctx context.Context) error { compileWeek(diaryHTML); return nil }),
					); err != nil {
						log.Fatal(err)
					}
				}
				return nil
			}),
		); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	//var url = "https://growdiaries.com/auth/signin"

	ctx, cancel, err := cu.New(cu.NewConfig(
		//cu.WithHeadless(),
		cu.WithChromeFlags(chromedp.WindowSize(600, 800)),
	))
	if err != nil {
		panic(err)
	}
	defer cancel()

	//login(ctx, url)
	//userDiariesList := getUserDiariesListHTML(ctx, "royal-queen-seeds/northern-light")
	//diariesListURLs := compileUserDiariesList(userDiariesList)
	var diariesListURLs = []string{"/diaries/213233-royal-queen-seeds-northern-light-grow-journal-by-eigenheit"}
	//var diariesListURLs = []string{"/diaries/162365-grow-journal-by-manilagrowop"}
	getUserDiary(ctx, diariesListURLs)
}
