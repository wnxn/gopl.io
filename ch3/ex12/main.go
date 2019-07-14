package main

import "fmt"

func main() {
	s1 := "wangxin"
	s2 := "xwangni"
	s3 := "wangxn"
	s4 := "Hello,你好"
	s5 := "你Hell,好o"
	fmt.Println(sameBytes(s1, s2))
	fmt.Println(sameBytes(s1, s3))
	fmt.Println(sameBytes(s4, s5))
}

func sameBytes(s1, s2 string) bool {
	m1, m2 := getMap(s1), getMap(s2)
	return isEqual(m1, m2)
}

func getMap(s string) map[rune]int {
	res := map[rune]int{}
	for _, v := range s {
		res[v]++
	}
	return res
}

func isEqual(m1, m2 map[rune]int) bool {
	for k, _ := range m1 {
		if m2[k] != m1[k] {
			return false
		}
	}
	for k, _ := range m2 {
		if m2[k] != m1[k] {
			return false
		}
	}
	return true
}

func printMap(m map[rune]int) {
	fmt.Printf("%v", m)
}
