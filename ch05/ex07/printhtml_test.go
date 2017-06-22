package printhtml

import "testing"

func OutlineForTest() {
	outline("https://golang.org")
}

func TestValidHtml(t *testing.T) {
	OutlineForTest()
}
