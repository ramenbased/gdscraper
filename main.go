package main

import (
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

func getHtml() string {
	url := "https://growdiaries.com/auth/signin"

	ctx, cancel, err := cu.New(cu.NewConfig(
		// Remove this if you want to see a browser window.
		//cu.WithHeadless(),
		// If the webelement is not found within 10 seconds, timeout.
		cu.WithChromeFlags(chromedp.WindowSize(600, 800)),
	//cu.WithTimeout(10*time.Second),
	))

	if err != nil {
		panic(err)
	}
	defer cancel()
	log.Println("Starting Chrome")

	var outer string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(20*time.Second),
		chromedp.WaitVisible("#email"),
		chromedp.SendKeys("#email", "totallyhuman@cock.li"),
		chromedp.SendKeys("#password", "iamhuman"),
		chromedp.MouseClickXY(165, 480),
		chromedp.Sleep(5*time.Second),
		chromedp.EvaluateAsDevTools("document.querySelector('.btn.primary').click()", nil),
		chromedp.Sleep(20*time.Second),
	); err != nil {
		log.Fatal(err)
	}

	return outer
}

func main() {

	outer := getHtml()
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
