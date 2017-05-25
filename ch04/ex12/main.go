package main

import "./xkcd"

func main() {
	for i := 500; i < 600; i++ {
		xkcd.CreateOfflineIndex(i)
	}

}
