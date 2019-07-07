package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func main() {
	resp,err:=http.Get("https://golang.org")
	if err !=nil{
		fmt.Errorf(err.Error())
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err !=nil{
		fmt.Errorf(err.Error())
	}
	res := ElementsByTagName(doc, "a")
	for _,v:=range res{
		fmt.Printf("%+v\n",v)
	}
//	forEachNode(doc, startElement, endElement)
}

func ElementsByTagName(n *html.Node, name ...string)[]*html.Node{
	res := []*html.Node{}
	if n.Type == html.ElementNode {
		if arrayContains(n.Data, name...){
			//fmt.Println(n)
			res = append(res, n)
		}
	}
	for i:=n.FirstChild; i != nil; i= i.NextSibling{
		tmp:=ElementsByTagName(i, name...)
		res = append(res,tmp...)
	}
	return res
}

func arrayContains(find string, array ...string)bool{
	for _,v:=range array{
		if v==find{
			return true
		}
	}
	return false
}