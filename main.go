package main

import (
	"context"
	"fmt"
	"gdscraper/data"
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

// Week overview
func compileItems(itemsHTML string) (roomType string, growMedium []string) {
	var sr = strings.NewReader(itemsHTML)
	var doc, err = html.Parse(sr)
	Er(err)

	var f func(*html.Node)
	f = func(n *html.Node) {
		//TODO: fill structs
		//RoomType and Substrate
		if n.Type == html.ElementNode && n.Data == "div" {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "info" {
					if n.FirstChild != nil && n.FirstChild.Data == "div" {
						switch n.LastChild.FirstChild.Data {
						case "Room Type":
							roomType = n.FirstChild.FirstChild.Data
						case "Grow medium":
							growMedium = append(growMedium, n.FirstChild.FirstChild.Data)
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
	return roomType, growMedium
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

// Actual Weeks
func getUserDiary(ctx context.Context, URLs []string, tbl *data.Tables) {
	var itemsHTML string
	var weeksHTML string
	//iterate over URLs
	for _, diaryURL := range URLs {

		if err := chromedp.Run(ctx,
			chromedp.Navigate("https://growdiaries.com"+diaryURL),
			chromedp.Sleep(3*time.Second),
			chromedp.OuterHTML(".report_items.report_seeds", &itemsHTML),
			chromedp.OuterHTML(".day_items", &weeksHTML),

			chromedp.ActionFunc(func(ctx context.Context) error {
				//TODO: Sanity check here
				roomType, soils := compileItems(itemsHTML) //TODO: return soils with %

				var diary = new(data.Diary)
				var soil = new(data.Soil)
				diary.AddDiary(regexGetID(diaryURL), diaryURL, roomType, tbl)
				soil.AddSoil(diary.ID, soils, tbl)

				//TODO: add soils with % and ID from diary.ID)

				weeks := compileWeekOverview(weeksHTML) //returns TempWeek stuct for chrome to iterate over weeks
				for _, w := range weeks.w {
					var diaryHTML string
					fmt.Println("internal.. ", w.Link, w.WeekType)
					if err := chromedp.Run(ctx,
						chromedp.Navigate("https://growdiaries.com"+w.Link),
						chromedp.Sleep(10*time.Second),
						chromedp.OuterHTML("#app", &diaryHTML, chromedp.ByID),

						chromedp.ActionFunc(func(ctx context.Context) error { compileWeek(diaryHTML, w); return nil }),
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
	ctx, cancel, err := cu.New(cu.NewConfig(
		//cu.WithHeadless(),
		cu.WithChromeFlags(chromedp.WindowSize(600, 800)),
	))
	if err != nil {
		panic(err)
	}
	defer cancel()

	var tbl = new(data.Tables)

	login(ctx, "https://growdiaries.com/auth/signin")
	//userDiariesList := getUserDiariesListHTML(ctx, "royal-queen-seeds/northern-light")
	//diariesListURLs := compileUserDiariesList(userDiariesList)
	//var diariesListURLs = []string{"/diaries/197545-royal-queen-seeds-northern-light-grow-journal-by-nugcaleb"} //random test
	var diariesListURLs = []string{"/diaries/149912-grow-journal-by-madebyfrancesco"} //multiple soils

	//var diariesListURLs = []string{"/diaries/213233-royal-queen-seeds-northern-light-grow-journal-by-eigenheit"}
	getUserDiary(ctx, diariesListURLs, tbl)
}
