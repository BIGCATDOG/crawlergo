package Storage

import (
	"Crawler/src/Resource"
	"encoding/json"
	"os"
)

func NewLocalStorage(storagePath string) *LocalStorage  {
	exist,_ :=isExists(storagePath)
	if !exist{
		os.Mkdir(storagePath,os.ModeDir)
	}
	return &LocalStorage{storagePath: storagePath}
}
type LocalStorage struct {
	storagePath string
}

func (l LocalStorage) read(sourceString string, translatedLanguage string) Resource.ResourceItem {
	return Resource.ResourceItem{}
}

func (l LocalStorage) write(item *Resource.ResourceItem) bool {
	 if item!=nil{
	 	file,err :=os.Create("data/"+item.SourceString)
	 	if err==nil{
	 		bytes,_ := json.Marshal(item)
	 		wrLen,err:=file.Write(bytes)
	 		if err==nil&&wrLen==len(bytes){
	 			return true
			}
		}

	 }
	 return false
}


