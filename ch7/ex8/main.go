package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Table struct{
	tracks []*Track
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
			if t.tracks[i].Title != t.tracks[j].Title{
				return t.tracks[i].Title < t.tracks[j].Title
			}
		case TrackArtistName:
			if t.tracks[i].Artist != t.tracks[j].Artist{
				return t.tracks[i].Artist < t.tracks[j].Artist
			}
		case TrackAlbumName:
			if t.tracks[i].Album != t.tracks[j].Album{
				return t.tracks[i].Album < t.tracks[j].Album
			}
		case TrackYearName:
			if t.tracks[i].Year != t.tracks[j].Year{
				return t.tracks[i].Year < t.tracks[j].Year
			}
		case TrackLengthName:
			if t.tracks[i].Length != t.tracks[j].Length{
				return t.tracks[i].Length < t.tracks[j].Length
			}
		}
	}
	return false
}

func (t Table)Len()int{
	return len(t.tracks)
}

func (t Table)Swap(i,j int){
	t.tracks[i], t.tracks[j]=t.tracks[j],t.tracks[i]
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
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
	t := Table{tracks:tracks,priority:priority}
	printTracks(t.tracks)
	sort.Sort(t)
	printTracks(t.tracks)

	t.Click(TrackAlbumName)
	printTracks(t.tracks)

	t.Click(TrackLengthName)
	sort.Sort(t)
	printTracks(t.tracks)
}
