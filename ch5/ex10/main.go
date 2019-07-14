package main

import (
	"fmt"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, v := range toposort(prereqs) {
		fmt.Println(i, v)
	}
}

func toposort(prereqs map[string][]string) []string {
	res := []string{}
	keys := map[string]bool{}
	var visit func(map[string]bool)
	visit = func(mp map[string]bool) {
		for k, _ := range mp {
			if !keys[k] {
				keys[k] = true
				visit(stringsToMap(prereqs[k], keys))
				res = append(res, k)
			}
		}
	}

	for k := range prereqs {
		keys[k] = false
	}
	visit(keys)
	return res
}

func stringsToMap(s []string, visited map[string]bool) map[string]bool {
	res := map[string]bool{}
	for i := range s {
		res[s[i]] = visited[s[i]]
	}
	return res
}
