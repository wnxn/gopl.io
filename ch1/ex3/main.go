package main

import (
	"os"
	"strings"
)

func echo1(args []string) string {
	var s, sep string
	for i := 1; i < len(args[1:]); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func echo2(args []string) string {
	var s, sep string
	for _, value := range args[1:] {
		s += sep + value
		sep = " "
	}
	return s
}

func echo3(args []string) string {
	return strings.Join(args, " ")
}
