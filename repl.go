package main

import "strings"

func cleanInput(text string) []string {
	stringsSlc := []string{}

	for string := range strings.FieldsSeq(text) {
		stringsSlc = append(stringsSlc, strings.ToLower(string))
	}

	return stringsSlc
}
