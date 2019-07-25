package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

var trackList = template.Must(template.New("trackList").Parse(`
<h1>tracks</h1>
<table>
<tr style='text-align: left'>
  <th>Title</th>
  <th>Artist</th>
  <th>Album</th>
  <th>Year</th>
  <th>Length</th>
</tr>
{{range .}}
<tr>
  <td>{{.Title}}</a></td>
  <td>{{.Artist}}</a></td>
  <td>{{.Album}}</a></td>
  <td>{{.Year}}</a></td>
  <td>{{.Length}}</a></td>
</tr>
{{end}}
</table>
`))

type Table struct{
	Tracks []*Track
	priority map[PriorityLevel]TrackField
}

type PriorityLevel int

type TrackField string

const(
	TrackTitleName = "Title"
	TrackArtistName = "Artist"
	TrackAlbumName = "Album"
	TrackYearName = "Year"
	TrackLengthName = "Length"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

var priority = map[PriorityLevel]TrackField{
	1:TrackTitleName,
	4:TrackArtistName,
	3:TrackAlbumName,
	2:TrackYearName,
	5:TrackLengthName}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func (t Table)Less(i,j int)bool{
	for k:=0;k<5;k++{
		switch t.priority[PriorityLevel(k)]{
		case TrackTitleName:
			if t.Tracks[i].Title != t.Tracks[j].Title{
				return t.Tracks[i].Title < t.Tracks[j].Title
			}
		case TrackArtistName:
			if t.Tracks[i].Artist != t.Tracks[j].Artist{
				return t.Tracks[i].Artist < t.Tracks[j].Artist
			}
		case TrackAlbumName:
			if t.Tracks[i].Album != t.Tracks[j].Album{
				return t.Tracks[i].Album < t.Tracks[j].Album
			}
		case TrackYearName:
			if t.Tracks[i].Year != t.Tracks[j].Year{
				return t.Tracks[i].Year < t.Tracks[j].Year
			}
		case TrackLengthName:
			if t.Tracks[i].Length != t.Tracks[j].Length{
				return t.Tracks[i].Length < t.Tracks[j].Length
			}
		}
	}
	return false
}

func (t Table)Len()int{
	return len(t.Tracks)
}

func (t Table)Swap(i,j int){
	t.Tracks[i], t.Tracks[j]=t.Tracks[j],t.Tracks[i]
}

func printTracks(tracks []*Track) {
	if err := trackList.Execute(os.Stdout,tracks); err != nil {
		log.Fatal(err)
	}
}

func (t *Table)Click(field TrackField){
	mp := t.priority
	if mp[1] == field{
		return
	}
	var fieldLevel PriorityLevel = 0
	for i := range mp{
		if mp[i] == field{
			fieldLevel = i
		}
	}
	for i:= fieldLevel; i>1;i--{
		mp[i] = mp[i-1]
	}
	mp[1]=field
}

func main() {
	t := Table{Tracks:tracks,priority:priority}
	printTracks(t.Tracks)
}
