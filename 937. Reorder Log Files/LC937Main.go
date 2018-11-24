package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func reorderLogFiles(logs []string) []string {

	type log struct {
		identifier string
		isNumber   bool
		message    string
	}
	tmpResult := make([]log, len(logs))
	fmt.Println(logs)
	// expression := regexp.QuoteMeta("^\\d+")
	reg, err1 := regexp.Compile("^[0-9]+ *")

	if err1 != nil {
		fmt.Println(err1)
	}

	for i, item := range logs {

		s := strings.SplitN(item, " ", 2)

		tmpResult[i] = log{
			identifier: s[0],
			isNumber:   reg.MatchString(s[1]),
			message:    s[1],
		}
		// print("identifier : ", tmpResult[i].identifier, " messages: ", tmpResult[i].message, " isNumber: ", tmpResult[i].isNumber)
		// fmt.Println(s)

		// fmt.Println("")
	}
	sort.SliceStable(tmpResult, func(i, j int) bool {
		if tmpResult[i].isNumber && tmpResult[j].isNumber {
			return false
		}
		if btou(tmpResult[i].isNumber)+btou(tmpResult[j].isNumber) == 1 {
			return !tmpResult[i].isNumber
		}

		return tmpResult[i].message < tmpResult[j].message
	})
	result := make([]string, len(tmpResult))
	for i, item := range tmpResult {
		// print("Second identifier : ", item.identifier, "  messages: ", item.message)
		// fmt.Println("")
		result[i] = item.identifier + " " + item.message
	}
	return result
}

func main() {

	s := "abc def"
	t := "abc bcd"
	fmt.Printf(strconv.FormatBool(s > t))
}

func btou(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}
