package Storage

import (
	"Crawler/src/Resource"
	"os"
)

func NewStorage(storageType StorageType,storagePath string) Storage  {
	switch storageType {
	case LocalStorageType:
		return NewLocalStorage(storagePath)
	}


	return nil
}
type Storage interface {
	Read(sourceString string,translatedLanguage string)Resource.ResourceItem
	Write(item *Resource.ResourceItem)bool
}

func isExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}