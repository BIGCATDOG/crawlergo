package consumer

import (
	"compress/flate"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
)

type BaiduConsumer struct {
	taskChan     <-chan *http.Request
	taskMaxCount int
}

func NewConsumer(taskMaxCount int,taskChan chan *http.Request) ConsumerInterface  {
	return &BaiduConsumer{taskMaxCount: taskMaxCount,taskChan: taskChan}
}
func (b *BaiduConsumer) CanDo(workType WorkType) bool {
	if workType == WorkHttpRequest {
		return true
	}
	return false
}

func (b *BaiduConsumer) Work() {
	for i := 0; i < b.taskMaxCount; i++ {
		go func() {
			for {
				req, _ := <-b.taskChan
				client := http.Client{}
				resp, _ := client.Do(req)
				reader, _ := switchContentEncoding(resp)

				data := make([]byte, resp.ContentLength)

				reader.Read(data)
				defer resp.Body.Close()

				//b.storage.Write(&Resource.ResourceItem{SourceString: word, TranslatedString: "hhh", TranslatedLanguage: "zh-cn"})

				fmt.Println("response is " + string(data) + "\r\n")
			}
		}()

	}
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
