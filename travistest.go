package travistest

import (
	"io/ioutil"
	"os"
)

func CreateFile(filename string) {
	err := ioutil.WriteFile(filename, []byte("Hello"), 0644)
	if err != nil {
		panic(err)
	}
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
