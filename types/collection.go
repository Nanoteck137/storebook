package types

import "errors"

type CollectionType string

const (
	CollectionTypeUnknown CollectionType = "unknown"
	CollectionTypeSeries  CollectionType = "series"
	CollectionTypeAnime   CollectionType = "anime"
)

// func (t MediaType) IsMovie() bool {
// 	return t == MediaTypeMovie || t == MediaTypeAnimeMovie
// }

func IsValidCollectionType(t CollectionType) bool {
	switch t {
	case CollectionTypeUnknown,
		CollectionTypeSeries,
		CollectionTypeAnime:
		return true
	}

	return false
}

func ValidateCollectionType(val any) error {
	if s, ok := val.(string); ok {
		if s == "" {
			return nil
		}

		t := CollectionType(s)
		if !IsValidCollectionType(t) {
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

		t := CollectionType(s)
		if !IsValidCollectionType(t) {
			return errors.New("invalid type")
		}
	} else {
		return errors.New("expected string")
	}

	return nil
}
