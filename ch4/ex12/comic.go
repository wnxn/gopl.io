package ex12

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var buffer []Item

type ItemSearchResult struct {
	TotalCound int
	Items []*Item
}

type Item struct {
	Month string `json: "month"`
	Num int `json: "num"`
	Link string `json: "link"`
	Year string `json: "year"`
	News string `json: "news"`
	Safe_Title string `json: "safe_title"`
	Transcript string `json: "transcript"`
	Alt string `json: "alt"`
	Img string `json: "img"`
	Day string `json: "day"`
	Title string `json:"title"`
}

func IndexUrl(index int)string{
	if index <=0{
		return ""
	}
	return fmt.Sprintf("https://xkcd.com/%d/info.0.json", index)
}

func GetContent(index int)(item *Item, err error){
	resp, err := http.Get(IndexUrl(index))
	if err != nil{
		return nil, err
	}
	if resp.StatusCode != http.StatusOK{
		resp.Body.Close()
		return nil,fmt.Errorf("search query failed: %s", resp.Status)
	}
	item = &Item{}
	if err := json.NewDecoder(resp.Body).Decode(item);err !=nil{
		return nil, err
	}
	resp.Body.Close()
	return item, nil
}

func Download(){
	for i:=1;i<20;i++{
		content, err := GetContent(i)
		if err != nil{
			fmt.Printf("Get %d content failed: %s", i, err.Error())
		}
		buffer = append(buffer, *content)
	}
}

func SearchAsYear(year int)[]Item{
	res := []Item{}
	for _, v:=range buffer{
		res = append(res, v)
	}
	return res
}