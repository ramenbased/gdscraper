package main

import (
	"context"
	"fmt"
	"log"
	"time"

	cu "github.com/Davincible/chromedp-undetected"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
)

func Er(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func getUserDiariesListHTML(ctx context.Context, strain string) string {
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
	getUserDiary(ctx, diariesListURLs)
}
