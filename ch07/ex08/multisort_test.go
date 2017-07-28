package multisort

import (
	"sort"
	"testing"
)

func TestMultiSort(t *testing.T) {
	var tracks = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
		{"Go", "Delilah", "From the Roots Up", 1992, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}

	sort.Sort(Multi{tracks, "artist"})
	sort.Sort(Multi{tracks, "title"})
	sort.Sort(Multi{tracks, "year"})
	for _, track := range tracks {
		print(track.Year)
		print("\t")
		print(track.Title)
		print("\t")
		println(track.Artist)
	}
}

func TestMultiSort2(t *testing.T) {
	var tracks = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
		{"Go", "Delilah", "From the Roots Up", 1992, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("4m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m35s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m23s")},
	}

	sort.Sort(Multi{tracks, "length"})
	sort.Sort(Multi{tracks, "album"})
	sort.Sort(Multi{tracks, "title"})
	for _, track := range tracks {
		print(track.Year)
		print("\t")
		print(track.Title)
		print("\t")
		println(track.Artist)
	}
}
