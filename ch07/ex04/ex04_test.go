package ex04

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"golang.org/x/net/html"
)

func TestNewReader(t *testing.T) {
	str, err := RespString("https://golang.org")
	println(str)
	_, err = html.Parse(NewReader(str))
	if err != nil {
		t.Error("")
	}
}

func RespString(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return "", fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", fmt.Errorf("parsing %s as string: %v", url, err)
	}
	return string(b), err
}
