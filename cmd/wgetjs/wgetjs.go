package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

var (
	targetUrl    = flag.String("target-url", "", "")
	readyTimeout = flag.Uint("ready-timeout", 2500, "")
	outputFile   = flag.String("output", "", "")
)

func init() {
	flag.Parse()
}

func main() {
	if nil == targetUrl || len(*targetUrl) == 0 {
		fmt.Println("Bad option: target-url can not be empty!")
	}

	if nil == readyTimeout || *readyTimeout == 0 {
		fmt.Println("Bad option: ready-timeout must be great than zero!")
	}

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var htmlData string
	err := chromedp.Run(ctx, chromedp.Tasks{
		network.Enable(),
		network.SetExtraHTTPHeaders(network.Headers(map[string]interface{}{
			"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3",
			"Accept-Language": "ru,en-US;q=0.9,en;q=0.8,uk;q=0.7",
			"User-Agent":      "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36",
		})),

		chromedp.Navigate(*targetUrl),
		chromedp.Sleep(time.Millisecond * time.Duration(*readyTimeout)),

		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			htmlData, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			return err
		}),
	})
	if err != nil {
		log.Fatal(err)
	}

	if nil != outputFile && len(*outputFile) > 0 {
		if err = ioutil.WriteFile(*outputFile, []byte(htmlData), 0644); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println(htmlData)
	}
}
