package utils

import "fmt"

func GenerateTitle(startTag string, endTag string) string {
	if startTag == "" && endTag == "" {
		return "[MR Tracker] All Changes"
	} else if startTag == "" {
		return fmt.Sprintf("[MR Tracker] Changes before %s", endTag)
	} else if endTag == "" {
		return fmt.Sprintf("[MR Tracker] Changes after %s", startTag)
	} else {
		return fmt.Sprintf("[MR Tracker] Changes between %s - %s", startTag, endTag)
	}
}
