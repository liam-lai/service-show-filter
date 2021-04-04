package util

import (
	"regexp"
)

func RemoveBracket(s string) string {
	bracketRE := regexp.MustCompile(`[\s]*\(.*\)[\s]*`)
	removedBracketStr := bracketRE.ReplaceAllString(s, " ")

	frontAndBackSpaceRE := regexp.MustCompile(`^[\s]+|[\s]+$`)
	removedSpaceStr := frontAndBackSpaceRE.ReplaceAllString(removedBracketStr, "")
	return removedSpaceStr
}
