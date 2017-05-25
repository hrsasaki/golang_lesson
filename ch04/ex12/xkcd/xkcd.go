package xkcd

import (
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func CreateOfflineIndex(index int) {
	url := "https://xkcd.com/" + strconv.Itoa(index) + "/info.0.json"
	println(url)
	resp, err := http.Get(url)
	if err != nil {
		resp.Body.Close()
		println(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		println(err)
		return
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		println(err)
		return
	}
	resp.Body.Close()
	filename := "xkcd" + strconv.Itoa(index) + ".json"
	println(filename)
	err = ioutil.WriteFile(filename, responseData, os.ModePerm)
	if err != nil {
		println(err)
		return
	}
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
