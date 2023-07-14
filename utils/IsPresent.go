package utils

import "strings"

func IsPresent(tab []string, s string) bool {
	for _, val := range tab {
		if strings.EqualFold(val, s) {
			return true
		}
	}
	return false
}
