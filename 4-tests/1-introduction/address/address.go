package address

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// TypeOfAddress takes an address string and returns the type of address based on the first word.
// It checks if the first word contains any of the valid address types.
func TypeOfAddress(address string) string {
	typeValids := []string{"road", "street", "avenue", "boulevard", "drive", "lane", "court", "place", "way", "circle"}

	fristWord := strings.Split(strings.ToLower(address), " ")[0]

	addressValidType := false
	for _, typeValid := range typeValids {
		if strings.Contains(fristWord, typeValid) {
			addressValidType = true
			break
		}
	}
	if addressValidType {
		return cases.Title(language.English).String(fristWord)
	}
	return "unknown"
}
