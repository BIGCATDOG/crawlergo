package Storage

import "os"

type Storage interface {
	read(sourceString string,translatedLanguage string)string
	write(sourceString string,translatedString string,translatedLanguage string)string
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