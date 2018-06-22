package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

func log(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
}

func load(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		io.Copy(ioutil.Discard, resp.Body)
		return nil, errors.Errorf("wrong HTTP status code: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

func has(all []string, val string) bool {
	for _, v := range all {
		if v == val {
			return true
		}
	}
	return false
}
