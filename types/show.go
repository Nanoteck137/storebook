package types

import "errors"

type ShowType string

const (
	ShowTypeUnknown  ShowType = "unknown"
	ShowTypeTVSeries ShowType = "tv-series"
	ShowTypeAnime    ShowType = "anime"
)

func IsValidShowType(t ShowType) bool {
	switch t {
	case ShowTypeUnknown,
		ShowTypeTVSeries,
		ShowTypeAnime:
		return true
	}

	return false
}

func ValidateShowType(val any) error {
	if s, ok := val.(string); ok {
		if s == "" {
			return nil
		}

		t := ShowType(s)
		if !IsValidShowType(t) {
			return errors.New("invalid type")
		}
	} else if p, ok := val.(*string); ok {
		if p == nil {
			return nil
		}

		s := *p
		if s == "" {
			return nil
		}

		t := ShowType(s)
		if !IsValidShowType(t) {
			return errors.New("invalid type")
		}
	} else {
		return errors.New("expected string")
	}

	return nil
}
