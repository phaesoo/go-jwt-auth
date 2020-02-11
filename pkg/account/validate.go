package account

import (
	"fmt"
	"regexp"
	"unicode"
)

// IsValidUsername : validate username
func IsValidUsername(s string) (bool, error) {
	// check length
	length := len(s)
	if length < 6 {
		return false, fmt.Errorf("username length should be at least 6")
	} else if length > 20 {
		return false, fmt.Errorf("Username length should be not be greater than 20")
	}

	// check first letter
	if unicode.IsOneOf([]*unicode.RangeTable{unicode.Number, unicode.Digit}, rune(s[0])) {
		return false, fmt.Errorf("Username should not have to start with digits")
	}

	// check by loop
	for _, classes := range map[string][]*unicode.RangeTable{
		"valid types": {unicode.Lower, unicode.Number, unicode.Digit},
	} {
		for _, r := range s {
			if !unicode.IsOneOf(classes, r) {
				return false, fmt.Errorf("Username should have to be composed of lower case and digits")
			}
		}
	}
	return true, nil
}

// IsValidPassword : validate password
func IsValidPassword(s string) (bool, error) {
	// check length
	if len(s) < 6 {
		return false, fmt.Errorf("Password length should be at least 6")
	} else if len(s) > 20 {
		return false, fmt.Errorf("Password length should be not be greater than 20")
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
		return false, fmt.Errorf("Password must have at least one %s character", name)
	}
	return true, nil
}

// IsValidEmail : validate email
func IsValidEmail(s string) (bool, error) {
	match, err := regexp.MatchString(`^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$`, s)
	if err != nil {
		return false, err
	}
	return match, nil
}
