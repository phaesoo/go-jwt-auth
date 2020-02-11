package account

import (
	"fmt"
	"unicode"
)

func CheckUsername(s string) error {
	// check length
	length := len(s)
	if length < 6 {
		return fmt.Errorf("username length should be at least 6")
	} else if length > 20 {
		return fmt.Errorf("Username length should be not be greater than 20")
	}

	// check first letter
	if unicode.IsOneOf([]*unicode.RangeTable{unicode.Number, unicode.Digit}, rune(s[0])) {
		return fmt.Errorf("Username should not have to start with digits")
	}

	// check by loop
	for name, classes := range map[string][]*unicode.RangeTable{
		"valid types": {unicode.Lower, unicode.Number, unicode.Digit},
	} {
		for _, r := range s {
			if !unicode.IsOneOf(classes, r) {
				return fmt.Errorf("Username should have to be composed of lower case and digits", name)
			}
		}
	}
	return nil
}

// copy from https://gist.github.com/fearblackcat/d0199d6a47d60b067a4d4be173b0ef97
func CheckPassword(s string) error {
	// check length
	if len(s) < 6 {
		return fmt.Errorf("Password length should be at least 6")
	} else if len(s) > 20 {
		return fmt.Errorf("Password length should be not be greater than 20")
	}

next:
	for name, classes := range map[string][]*unicode.RangeTable{
		"upper case": {unicode.Upper, unicode.Title},
		"lower case": {unicode.Lower},
		"numeric":    {unicode.Number, unicode.Digit},
		"special":    {unicode.Space, unicode.Symbol, unicode.Punct, unicode.Mark},
	} {
		for _, r := range s {
			if unicode.IsOneOf(classes, r) {
				continue next
			}
		}
		return fmt.Errorf("Password must have at least one %s character", name)
	}
	return nil
}
