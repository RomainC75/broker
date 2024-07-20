package utils

import "unicode"

func ToEnvString(key string) string {
	// ThisIsVar
	res := ""
	for i, r := range key {
		if i == 0 {
			res += string(unicode.ToUpper(r))
		} else if unicode.IsLower(r) {
			res += string(unicode.ToUpper(r))
		} else if unicode.IsUpper(r) {
			res += "_"
			res += string(unicode.ToUpper(r))
		}
	}
	return res
}

func ToStructKeyString(key string) string {
	res := ""
	foundLowDash := false
	for i := 0; i < len(key); i++ {
		if i == 0 {
			res += string(unicode.ToUpper(rune(key[i])))
		} else if rune(key[i]) == '_' {
			foundLowDash = true
			continue
		} else if foundLowDash {
			res += string(unicode.ToUpper(rune(key[i])))
			foundLowDash = false
		} else if !foundLowDash {
			res += string(unicode.ToLower(rune(key[i])))
		}
	}
	return res
}
