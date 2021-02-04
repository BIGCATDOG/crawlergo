package analyzer

import (
	"Crawler/src/Resource"
	"Crawler/src/Storage"
	"encoding/json"
	"net/http"
)
type TranslationPair struct {
	Key string `json:"k"`
	Value string `json:"v"`
}
type TranslationResult struct {
	Errno int `json:"errno"`
	Data  []TranslationPair `json:"data"`
}
type BaiduTranslatorAnalyzer struct {
     storage Storage.Storage
}

func (b BaiduTranslatorAnalyzer) buildRequest(response *http.Response) *http.Request {
	if response !=nil{
		var bodyBytes = make([]byte,response.ContentLength)
		if rdLen,_:= response.Body.Read(bodyBytes);int64(rdLen)!=response.ContentLength{

		}
		var resultData TranslationResult
		json.Unmarshal(bodyBytes,resultData)
		for _,v:= range resultData.Data{
			b.storage.Write(&Resource.ResourceItem{SourceString: v.Key,TranslatedString: v.Value,TranslatedLanguage: "zh-cn"})
		}
		//b.storage.Write(&Resource.ResourceItem{})
		//Todo Read Translation Result and store into storage
	}
	return nil
}
