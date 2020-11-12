package main

import (
	"fmt"
	"strings"
)

func firstNonRepeating(str string) string {
	minusculas := strings.ToLower(str)
	for i := range minusculas {
		if strings.Count(minusculas,minusculas[i:i+1]) == 1 {
			return str[i:i+1]
		}
	}
	return ""
}

func main() {
	fmt.Println(firstNonRepeating("sTreSS"))
	fmt.Println(firstNonRepeating("Go hang a salami, I'm a lasagna hog!"))
}