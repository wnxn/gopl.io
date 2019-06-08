package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type item struct{
	Title string
	Poster string
}

func GetItem(name string)(*item, error){
	request := fmt.Sprintf("http://www.omdbapi.com/?t=%s", name)
	resp, err := http.Post(request, "application/octet-stream", strings.NewReader(""))
	if err != nil{
		return nil, err
	}
	i := &item{}
	err = json.NewDecoder(resp.Body).Decode(i)
	return i, err
}

func GetImage(i item)error{
	fmt.Printf("url = %s\n",i.Poster)
	resp, err := http.Post(i.Poster, "image/jpeg",strings.NewReader(""))
	if err != nil{
		return  err
	}
	file,err:= os.Create(i.Title+".jpg")
	if err != nil{
		return err
	}
	bufio.NewWriter(file)
	io.Copy(file, resp.Body)
	return nil
}

func main(){
	filmName:="titanic"
	i, err :=GetItem(filmName)
	if err !=nil{
		fmt.Printf(err.Error())
	}
	fmt.Printf("%+v\n",i)
	err = GetImage(*i)
	if err !=nil{
		fmt.Printf(err.Error())
	}
}