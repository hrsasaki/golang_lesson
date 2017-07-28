package multisort

import "time"

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type Multi struct {
	tracks []*Track
	key    string
}

func (m Multi) Len() int {
	return len(m.tracks)
}

var history []string

func (m Multi) Less(i, j int) bool {
	history = append([]string{m.key}, history...)
	for _, k := range history {
		switch k {
		case "title":
			if m.tracks[i].Title == m.tracks[j].Title {
				continue
			}
			return m.tracks[i].Title < m.tracks[j].Title
		case "artist":
			if m.tracks[i].Artist == m.tracks[j].Artist {
				continue
			}
			return m.tracks[i].Artist < m.tracks[j].Artist
		case "album":
			if m.tracks[i].Album == m.tracks[j].Album {
				continue
			}
			return m.tracks[i].Album < m.tracks[j].Album
		case "year":
			if m.tracks[i].Year == m.tracks[j].Year {
				continue
			}
			return m.tracks[i].Year < m.tracks[j].Year
		case "length":
			if m.tracks[i].Length == m.tracks[j].Length {
				continue
			}
			return m.tracks[i].Length < m.tracks[j].Length
		}
	}
	return false
}

func (m Multi) Swap(i, j int) {
	m.tracks[i], m.tracks[j] = m.tracks[j], m.tracks[i]
}
