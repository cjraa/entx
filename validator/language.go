package validator

import "golang.org/x/text/language"

func ValidLanguage(s string) error {
	_, err := language.Parse(s)
	return err
}
