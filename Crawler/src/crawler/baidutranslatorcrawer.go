package crawler

import (
	"compress/flate"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func NewBaiduTranslatorCrawler(crawlerPeriodTimeMS uint32) *BaiduTranslatorCrawler {
	return &BaiduTranslatorCrawler{crawlerPeriodTimeMS: crawlerPeriodTimeMS};
}
type BaiduTranslatorCrawler struct {
	crawlerPeriodTimeMS uint32
}

func (b BaiduTranslatorCrawler) Start() {
      b.GenerateWords()
}

func (b BaiduTranslatorCrawler) Stop() {
	panic("implement me")
}

func (b BaiduTranslatorCrawler) GenerateWords() {
	var charByte []byte = []byte{'a','b','c','d','e'}
    var genWords[]byte = make([]byte,5)
	for i:=0;i<5;i++{
		for j:=0;j<5;j++{
			for k:=0;k<5;k++{
				for l:=0;l<5;l++{
					for m:=0;m<5;m++{
                       genWords[0]=charByte[i]
						genWords[1]=charByte[j]
						genWords[2]=charByte[k]
						genWords[3]=charByte[l]
						genWords[4]=charByte[m]
						b.BuildRequest(string(genWords))
					}
				}
			}
		}
	}
}

func (b BaiduTranslatorCrawler) BuildRequest(word string) {
	client := http.Client{}
	bodyJson,_:= json.Marshal( map[string]string{"kw":word})
	rep ,_:= http.NewRequest(http.MethodPost,"https://fanyi.baidu.com/sug",strings.NewReader(string(bodyJson)))

	rep.Header = map[string][]string{"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.104 Safari/537.36"},
		"Connection":{"keep-alive"},"Content-Type":{"application/x-www-form-urlencoded; charset=UTF-8"},"Accept":{"application/json, text/javascript, */*; q=0.01"},"Accept-Encoding":{"gzip, deflate, br"}}
	resp ,_:=client.Do(rep)
	reader,_ :=switchContentEncoding(resp)

	data := make([]byte,resp.ContentLength)

	reader.Read(data)
	defer resp.Body.Close()
	fmt.Printf("response is :%s",string(data))
}

func switchContentEncoding(res *http.Response) (bodyReader io.Reader, err error) {
	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		bodyReader, err = gzip.NewReader(res.Body)
	case "deflate":
		bodyReader = flate.NewReader(res.Body)
	default:
		bodyReader = res.Body
	}
	return
}