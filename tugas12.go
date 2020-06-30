package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "jeruk pisang9 ANGGUR7"

	exec1(text)
	exec2(text)
}

func exec1(str string) {
	regex, err := regexp.Compile(`(?i)[a-z]+\d`)
	if err != nil {
		fmt.Println(err.Error())
	}

	hasil1 := regex.FindAllString(str, -1)
	fmt.Println(hasil1)
}

func exec2(str string) {
	regex, err := regexp.Compile(`(?i)[a-z]+\d`)
	if err != nil {
		fmt.Println(err.Error())
	}

	hasil2 := regex.FindStringIndex(str)
	fmt.Println(hasil2)
}
