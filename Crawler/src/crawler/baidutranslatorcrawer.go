package crawler

import (
	"Crawler/src/Storage"
	"Crawler/src/TaskQueue"
	"Crawler/src/consumer"
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"
)

func NewBaiduTranslatorCrawler(crawlerPeriodTimeMS uint32) *BaiduTranslatorCrawler {
	return &BaiduTranslatorCrawler{crawlerPeriodTimeMS: crawlerPeriodTimeMS, storage: Storage.NewStorage(Storage.LocalStorageType, ""), threadPool: TaskQueue.NewThreadTool(6),taskChan: make(chan *http.Request,6)}
}

type BaiduTranslatorCrawler struct {
	crawlerPeriodTimeMS uint32
	storage             Storage.Storage
	threadPool          TaskQueue.ThreadPoolInterface
	taskChan chan *http.Request
	consumer consumer.ConsumerInterface
}

func (b *BaiduTranslatorCrawler) Start() {
	b.consumer = consumer.NewConsumer(6,b.taskChan)
	b.consumer.Work()
	b.GenerateWords()

}

func (b *BaiduTranslatorCrawler) Stop() {
	panic("implement me")
}

func (b *BaiduTranslatorCrawler) GenerateWords() {
	var charByte []byte = []byte{'a', 'b', 'c', 'd', 'e'}
	var genWords []byte = make([]byte, 5)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 5; k++ {
				for l := 0; l < 5; l++ {
					for m := 0; m < 5; m++ {
						genWords[0] = charByte[i]
						genWords[1] = charByte[j]
						genWords[2] = charByte[k]
						genWords[3] = charByte[l]
						genWords[4] = charByte[m]
						b.BuildRequest(string(genWords))


					}
				}
			}
		}
	}
}

func (b *BaiduTranslatorCrawler) BuildRequest(word string) {
	bodyJson, _ := json.Marshal(map[string]string{"kw": word})
	req, _ := http.NewRequest(http.MethodPost, "https://fanyi.baidu.com/sug", strings.NewReader(string(bodyJson)))
	req.Header = map[string][]string{"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.104 Safari/537.36"},
		"Connection": {"keep-alive"}, "Content-Type": {"application/json; charset=UTF-8"}, "Accept": {"application/json, text/javascript, */*; q=0.01"}, "Accept-Encoding": {"gzip, deflate, br"}}
	b.taskChan<-req
}



func u2s(form string) (to string, err error) {
	bs, err := hex.DecodeString(strings.Replace(form, `\u`, ``, -1))
	if err != nil {
		return
	}
	for i, bl, br, r := 0, len(bs), bytes.NewReader(bs), uint16(0); i < bl; i += 2 {
		binary.Read(br, binary.BigEndian, &r)
		to += string(r)
	}
	return
}
