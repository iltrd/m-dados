package main

import (
	"strings"
	"unicode"
)

func CleanData(data [][]string) [][]string {
	cleanedData := make([][]string, len(data))
	for i, record := range data {
		cleanedRecord := make([]string, len(record))
		for j, value := range record {
			cleanedValue := strings.Map(func(r rune) rune {
				if unicode.Is(unicode.Mn, r) {
					return -1
				}
				return r
			}, value)

			cleanedValue = strings.ToUpper(cleanedValue)

			cleanedRecord[j] = cleanedValue
		}
		cleanedData[i] = cleanedRecord
	}

	return cleanedData
}
