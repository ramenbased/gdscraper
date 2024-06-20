package main

import (
	"context"
	"fmt"
	"log"
	"time"

	cu "github.com/Davincible/chromedp-undetected"
	"github.com/chromedp/chromedp"
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

func test(ctx context.Context) {
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://growdiaries.com/diaries/54096-royal-queen-seeds-northern-light-grow-journal-by-w0oxtard/week/321888"),
		chromedp.Sleep(10*time.Second),
	); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var url = "https://growdiaries.com/auth/signin"
	var outer string

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

	login(ctx, url)
	test(ctx)

	fmt.Println(outer)

	/*
		sr := strings.NewReader(outer)
		doc, err := html.Parse(sr)
		Er(err)

		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Type == html.TextNode {
				fmt.Println(n.Data)
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
		f(doc)
	*/
}
