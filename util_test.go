package poeapi

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

const (
	repo = "github.com/willroberts/poeapi"
)

func loadFixture(filename string) (string, error) {
	path := fmt.Sprintf("%s/src/%s/%s", os.Getenv("GOPATH"), repo, filename)

	if runtime.GOOS == "windows" {
		path = strings.ReplaceAll(path, "/", "\\")
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
