package account

import (
	"fmt"
	"unicode"
)

func CheckUsername(username string) error {
	length := len(username)
	if length < 6 {
		return fmt.Errorf("Length should be at least 6")
	} else if length > 20 {
		return fmt.Errorf("Length should be not be greater than 20")
	}

	return nil
}

// copy from https://gist.github.com/fearblackcat/d0199d6a47d60b067a4d4be173b0ef97
func CheckPassword(s string) error {
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
		return fmt.Errorf("password must have at least one %s character", name)
	}
	return nil
}
