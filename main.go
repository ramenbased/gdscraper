package main

import (
	"context"
	"gdscraper/data"
	"io"
	"log"
	"os"
	"strings"
	"time"

	cu "github.com/Davincible/chromedp-undetected"
	"github.com/chromedp/chromedp"
	"golang.org/x/net/html"
)

func Er(err error) {
	if err != nil {
		log.Println(err)
	}
}

func replaceNilNodeData(n *html.Node) string {
	if n == nil {
		return "" //prev NULL
	} else {
		return n.Data
	}
}

func compileDiaryItems(itemsHTML string, diaryURL string, seedbank string, strain string, tbl *data.Tables) {
	var sr = strings.NewReader(itemsHTML)
	var doc, err = html.Parse(sr)
	Er(err)
	var d = new(data.Diary)

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "info" {
					switch n.LastChild.FirstChild.Data {
					case "Room Type":
						d.Environment = n.FirstChild.FirstChild.Data
					case "Grow medium":
						var soils = new(data.Soil)
						soils.AddSoil(regexGetID(diaryURL), n.FirstChild.FirstChild.Data, replaceNilNodeData(n.PrevSibling.LastChild.FirstChild), tbl)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	var rd = new(data.Diary)
	rd.AddDiary(regexGetID(diaryURL), diaryURL, d.Environment, seedbank, strain, tbl)

	/* TODO: dont scrape lights?..
	var f2 func(*html.Node)
	f2 = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "faza faza_0" {
					if n.FirshChild.Data == "VEG" {
						model := n.Parent.NextSibling.FirstChild.FirstChild.Data
						Obrand := ""
						if n.Parent.NextSibling.LastChild.FirstChild.Type == html.CommentNode {
							//known brand
							brand = n.Parent.NextSibling.LastChild.LastChild.FirstChild.Data
						}
						if n.Parent.NextSibling.LastChild.FirstChild.Type == html.TextNode {
							//custom brand
							brand = n.Parent.NextSibling.LastChild.FirstChild.Data
						}
						fmt.Println(model, "-", brand)
					}
				}
				if a.Key == "class" && a.Val == "faza faza_1" {
					if n.FirstChild.Data == "FLO" {
						model := n.Parent.NextSibling.FirstChild.FirstChild.Data
						brand := ""
						if n.Parent.NextSibling.LastChild.FirstChild.Type == html.CommentNode {
							//known brand
							brand = n.Parent.NextSibling.LastChild.LastChild.FirstChild.Data
						}
						if n.Parent.NextSibling.LastChild.FirstChild.Type == html.TextNode {
							//custom brand
							brand = n.Parent.NextSibling.LastChild.FirstChild.Data
						}
						fmt.Println(model, "-", brand)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f2(c)
		}
	}
	f2(doc)
	*/
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

func sanityWeekOverview(weeks *TempWeeks) *TempWeeks {
	veg, bloom, harvest := 0, 0, 0
	for _, w := range weeks.w {
		switch w.WeekType {
		case "0":
			veg += 1
		case "1":
			bloom += 1
		case "2":
			harvest += 1
		}
	}
	//TODO FINAL RULESET!
	if veg >= 2 && bloom >= 4 && harvest == 1 {
		log.Printf("internal.. sanity che=ck passed.. veg: %v bloom: %v harvest: %v\n", veg, bloom, harvest)
		weeks.sanity = true
	} else {
		log.Println("internal.. sanity check not passed, skip..")
	}
	return weeks
}

// Actual Weeks
func getUserDiary(ctx context.Context, URLs []string, seedbank string, strain string, tbl *data.Tables) {
	var itemsHTML string
	var weeksHTML string
	//iterate over URLs
	for _, diaryURL := range URLs {

		if err := chromedp.Run(ctx,
			chromedp.Navigate("https://growdiaries.com"+diaryURL),
			chromedp.Sleep(7*time.Second),
			chromedp.OuterHTML(".report_items.report_seeds", &itemsHTML),
			chromedp.OuterHTML(".day_items", &weeksHTML),

			chromedp.ActionFunc(func(ctx context.Context) error {

				var htmlID = regexGetID(diaryURL)
				weeks := compileWeekOverview(weeksHTML) //returns TempWeek stuct for chrome to iterate over weeks
				saneWeeks := sanityWeekOverview(weeks)

				if saneWeeks.sanity == true {
					//start data
					compileDiaryItems(itemsHTML, diaryURL, seedbank, strain, tbl)

					for _, w := range weeks.w {
						var diaryHTML string
						log.Println("internal.. ", w.Link, w.WeekType)
						if err := chromedp.Run(ctx,
							chromedp.Navigate("https://growdiaries.com"+w.Link),
							chromedp.Sleep(10*time.Second),
							chromedp.OuterHTML("#app", &diaryHTML, chromedp.ByID),
							chromedp.ActionFunc(func(ctx context.Context) error { compileDiaryWeek(diaryHTML, htmlID, w, tbl); return nil }),
						); err != nil {
							log.Fatal(err)
						}
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
	// LOG
	f, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	logstd := io.MultiWriter(os.Stdout, f)
	log.SetOutput(logstd)
	log.Printf("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@ START LOG\n")

	// CRHOME
	ctx, cancel, err := cu.New(cu.NewConfig(
		//cu.WithHeadless()
		cu.WithChromeFlags(chromedp.WindowSize(600, 800)),
	))
	if err != nil {
		panic(err)
	}
	defer cancel()

	//-------------------------------
	var tbl = new(data.Tables)

	login(ctx, "https://growdiaries.com/auth/signin")
	/*
		var seedbank = "fastbuds"
		var strain = "gorilla-cookies-auto"
	*/
	var seedbank = "royal-queen-seeds"
	var strain = "northern-light"
	userDiariesList := getUserDiariesListHTML(ctx, seedbank+"/"+strain)
	diariesListURLs := compileUserDiariesList(userDiariesList)

	//var diariesListURLs = []string{"/diaries/209445-zamnesia-seeds-x-10th-anniversary-grow-journal-by-schnabeldino"} //random test
	//var diariesListURLs = []string{"/diaries/149912-grow-journal-by-madebyfrancesco"} //multiple soils

	//var diariesListURLs = []string{"/diaries/171366-grow-journal-by-growwithflow/week/974789", "/diaries/213233-royal-queen-seeds-northern-light-grow-journal-by-eigenheit", "/diaries/209445-zamnesia-seeds-x-10th-anniversary-grow-journal-by-schnabeldino"}
	getUserDiary(ctx, diariesListURLs, seedbank, strain, tbl)
}
