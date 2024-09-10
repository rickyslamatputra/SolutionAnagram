package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Hi! This service will checking anagram")
	fmt.Println("Please Input the first string")
	str1, _ := reader.ReadString('\n')
	fmt.Println("Please Input the secong string")
	str2, _ := reader.ReadString('\n')

	fmt.Printf("Check Anagram : %t \n", checkAnagramBySort(str1, str2))
	//fmt.Printf("Check Anagram : %t \n", checkAnagramByCount(str1, str2))

}

func checkAnagramBySort(str1, str2 string) bool {
	return sortString(str1) == sortString(str2)
}

func sortString(str string) string {
	r := strings.Split(str, "")
	sort.Strings(r)
	return strings.Join(r, "")
}

func checkAnagramByCount(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	count := make(map[rune]int)

	for _, char := range str1 {
		count[char]++
	}
	for _, char := range str2 {
		count[char]--
		if count[char] < 0 {
			return false
		}
	}

	return true
}
