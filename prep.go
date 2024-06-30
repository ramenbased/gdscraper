package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
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
	//takes full html of one strain and gets URLs to report
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

func getUserDiariesListHTML(ctx context.Context, strain string) string {
	//scrolls down one strain of breeder to get full loaded html
	var rv string
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://growdiaries.com/seedbank/"+strain+"/diaries"),
		chromedp.Sleep(3*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			//how many times scroll to load
			for i := 0; i <= 1; i++ {
				chromedp.KeyEvent(kb.End).Do(ctx)
				chromedp.Sleep(2 * time.Second).Do(ctx)
			}
			return nil
		}),
		chromedp.OuterHTML(".report_boxs.reports_grid", &rv),
	); err != nil {
		log.Fatal(err)
	}
	return rv
}
