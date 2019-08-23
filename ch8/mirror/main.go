package main

import "fmt"

func mirroredQuery()string{
	responses := make(chan string,3)
	go func() {responses <- "asia"}()
	go func() {responses <- "america"}()
	go func() {responses <- "euro"}()

	return <-responses
}

func main() {
	fmt.Println(mirroredQuery())
}
