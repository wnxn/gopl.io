// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var flag byte
	fraction := ""
	if len(s) > 1 {
		if s[0] == '-' || s[0] == '+' {
			flag = s[0]
			s = s[1:]
		}
		index := strings.LastIndex(s, ".")
		if index >= 0 {

			s, fraction = s[:index], s[index:]
		}
	}
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		buf.WriteByte(s[i])
		if i != (len(s)-1) && (len(s)-1-i)%3 == 0 {
			buf.WriteRune(',')
		}
	}
	return string(flag) + buf.String() + fraction
}

//!-
