package crawler

import "errors"

func NewCrawler(crawlerType CrawlerType,crawlerPeriodTimeMS uint32) Crawler {
	switch crawlerType {
	case BaiduTranslator:
		;
	default:
		panic(errors.New("unhandled crawler type"))
	}
	return nil;
}
type Crawler interface {
	Start()
	Stop()
}


