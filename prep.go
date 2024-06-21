package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"golang.org/x/net/html"
)

func login(ctx context.Context, url string) {
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(20*time.Second),
		chromedp.WaitVisible("#email"),
		chromedp.SendKeys("#email", "totallyhuman@cock.li"),
		chromedp.SendKeys("#password", "iamhuman"),
		chromedp.MouseClickXY(165, 480),
		chromedp.Sleep(5*time.Second),
		chromedp.EvaluateAsDevTools("document.querySelector('.btn.primary').click()", nil),
		chromedp.Sleep(15*time.Second),
	); err != nil {
		log.Fatal(err)
	}
}

func compileUserDiariesList(outer string) []string {
	var rv []string

	sr := strings.NewReader(outer)
	doc, err := html.Parse(sr)
	Er(err)

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, i := range n.Attr {
				//TODO: does this need two iterations?!
				if i.Key == "href" {
					for _, y := range n.Attr {
						if y.Key == "class" && y.Val == "name" {
							rv = append(rv, i.Val)
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
	fmt.Printf("Found %v User Diaries \n", len(rv))
	return rv
}

func getUserDiary(ctx context.Context, URLs []string) {
	var items string
	for _, reportURL := range URLs {
		fmt.Printf("visiting %v \n", "https://growdiaries.com"+reportURL)
		if err := chromedp.Run(ctx,
			chromedp.Navigate("https://growdiaries.com"+reportURL),
			chromedp.Sleep(5*time.Second),
			chromedp.InnerHTML(".report_items.report_seeds", &items),
			chromedp.ActionFunc(func(ctx context.Context) error {
				//get and add items
				fmt.Println(items)
				return nil
			}),
		); err != nil {
			panic(err)
		}
	}
}
