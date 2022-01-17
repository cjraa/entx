package validator

import (
	"errors"
	"net/url"
	"strings"
)

func HttpURI(s string) error {
	uri, err := url.ParseRequestURI(strings.ToLower(s))
	ErrInvalidLink := errors.New("value should be valid link")
	if err != nil {
		return ErrInvalidLink
	}

	if uri.Scheme != "http" && uri.Scheme != "https" {
		return ErrInvalidLink
	}

	return nil
}
