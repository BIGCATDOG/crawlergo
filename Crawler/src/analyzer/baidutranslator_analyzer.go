package analyzer

import (
	"net/http"
)

type BaiduTranslatorAnalyzer struct {

}

func (b BaiduTranslatorAnalyzer) buildRequest(response http.Response) *http.Request {
	panic("implement me")
}
