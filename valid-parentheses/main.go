package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	var temp string
	open := "([{"
	closeBrackets := ")]}"
	for i := range s {
		res2 := strings.Contains(closeBrackets, string(s[i]))
		if res2 && temp == "" {
			return false
		}
		res1 := strings.Contains(open, string(s[i]))
		if res1 {
			temp += string(s[i])
		} else {
			last := temp[len(temp)-1:]
			if last == "(" && string(s[i]) == ")" {
				temp = temp[:len(temp)-1]
			} else if last == "[" && string(s[i]) == "]" {
				temp = temp[:len(temp)-1]
			} else if last == "{" && string(s[i]) == "}" {
				temp = temp[:len(temp)-1]
			} else {
				return false
			}
		}
	}
	if temp == "" {
		return true
	} else {
		return false
	}
}
func main() {
	strs := ")(){}"   // false
	strs1 := "()[]{}" // true
	strs2 := "(]"     // false
	strs3 := "{[]}"   // true
	strs4 := "([)]"   // false

	fmt.Println(isValid(strs))
	fmt.Println(isValid(strs1))
	fmt.Println(isValid(strs2))
	fmt.Println(isValid(strs3))
	fmt.Println(isValid(strs4))

}
