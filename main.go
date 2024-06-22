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
	sr := strings.NewReader(itemsHTML)
	doc, err := html.Parse(sr)
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

func compileWeekOverview(weeksHTML string) {
	sr := strings.NewReader(weeksHTML)
	doc, err := html.Parse(sr)
	Er(err)

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "data-faza" {
					for _, a2 := range n.Attr {
						if a2.Key == "href" {
							fmt.Printf("Weektype: %v %v \n", a.Val, a2.Val)
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

func getUserDiary(ctx context.Context, URLs []string) {
	var itemsHTML string
	var weeksHTML string
	for _, reportURL := range URLs {
		fmt.Printf("visiting %v \n", "https://growdiaries.com"+reportURL)
		if err := chromedp.Run(ctx,
			chromedp.Navigate("https://growdiaries.com"+reportURL),
			chromedp.Sleep(3*time.Second),
			chromedp.OuterHTML(".report_items.report_seeds", &itemsHTML),
			chromedp.ActionFunc(func(ctx context.Context) error {
				//get and add items
				compileItems(itemsHTML)
				return nil
			}),
			chromedp.OuterHTML(".day_items", &weeksHTML),
			chromedp.ActionFunc(func(ctx context.Context) error {
				//get and add items
				compileWeekOverview(weeksHTML)
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
	userDiariesList := getUserDiariesListHTML(ctx, "royal-queen-seeds/northern-light")
	diariesListURLs := compileUserDiariesList(userDiariesList)
	//var diariesListURLs = []string{"/diaries/213233-royal-queen-seeds-northern-light-grow-journal-by-eigenheit"}
	getUserDiary(ctx, diariesListURLs)
}
