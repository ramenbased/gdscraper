package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	cu "github.com/Davincible/chromedp-undetected"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"golang.org/x/net/html"
)

func Er(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

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

func getUserReportsListHTML(ctx context.Context) string {
	var rv string
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://growdiaries.com/seedbank/royal-queen-seeds/northern-light/diaries"),
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

func compileUserReportsList(outer string) {
	sr := strings.NewReader(outer)
	doc, err := html.Parse(sr)
	Er(err)

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, i := range n.Attr {
				//fmt.Printf("KEY: %v, VAL: %v", n.Key, n.Val)
				if i.Key == "href" {
					for _, y := range n.Attr {
						if y.Key == "class" && y.Val == "name" {
							fmt.Println(i.Val)
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
	log.Println("Starting Chrome")

	//login

	//login(ctx, url)
	userReportsList := getUserReportsListHTML(ctx)
	compileUserReportsList(userReportsList)
}
