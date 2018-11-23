package main

import (
	"sort"
	"fmt"
	"strconv"
	"strings"
)

func reorderLogFiles(logs []string) []string {

	type log struct {
		identifier string
		message    string
	}
	tmpResult := make([]log, len(logs))
	fmt.Println(logs)
	for i, item := range logs {
		
		s := strings.SplitN(item, " ", 2)
		tmpResult[i] = log{
			identifier: s[0],
			message:    s[1],
		}
		print("identifier : ",tmpResult[i].identifier, " messages: ",tmpResult[i].message)
		// fmt.Println(s)

		fmt.Println("")
	}
	sort.SliceStable(tmpResult,func(i,j int )bool{
		return tmpResult[i].message >= tmpResult[j].message
	})
	for _, item := range tmpResult {
		print("Second identifier : ",item.identifier, "  messages: ",item.message)
		fmt.Println("")
	}
	return []string{"g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"}
}

func main() {

	s := "abc def"
	t := "abc bcd"
	fmt.Printf(strconv.FormatBool(s > t))
}
