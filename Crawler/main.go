package main

import crawler2 "Crawler/src/crawler"

type hhh struct {
	Name string
}
func main()  {
	crawler := crawler2.NewCrawler(crawler2.BaiduTranslator,1)
	crawler.Start()
}
