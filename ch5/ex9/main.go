package main

import (
	"fmt"
	"strings"
)

// dwsa$foo dwsagpp
// dwsa$$foo dwsahqq

func main() {
	res := Expand("dwad$$foo$qwe", Trans)
	fmt.Println(res)
}

func Expand(s string, f func(string) string) string {
	res := ""
	tmp := ""
	for i := range s {
		if s[i] == '$' {
			if len(tmp) == 0 || tmp[len(tmp)-1] == '$' {
				tmp += string(s[i])
			} else {
				res += f(tmp)
				tmp = "$"
			}
		} else {
			if len(tmp) != 0 {
				tmp += string(s[i])
			} else {
				res += string(s[i])
			}
		}
	}
	if len(tmp) != 0 {
		res += f(tmp)
	}
	return res
}

func Trans(s string) string {
	if strings.ReplaceAll(s, "$", "") == "" {
		return ""
	}
	if s[0] == '$' {
		s = Trans(s[1:])
	} else {
		return s
	}
	res := []byte(s)
	for i := range res {
		res[i] += 1
	}
	return string(res)
}
