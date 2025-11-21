package types

import "errors"

type NotificationType string

const (
	NotificationTypeUnknown     NotificationType = "unknown"
	NotificationTypeGeneric     NotificationType = "generic"
	NotificationTypePartRelease NotificationType = "part-release"
)

func IsValidNotificationType(t NotificationType) bool {
	switch t {
	case NotificationTypeUnknown,
		NotificationTypeGeneric,
		NotificationTypePartRelease:
		return true
	}

	return false
}

func ValidateNotificationType(val any) error {
	if s, ok := val.(string); ok {
		if s == "" {
			return nil
		}

		t := NotificationType(s)
		if !IsValidNotificationType(t) {
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

		t := NotificationType(s)
		if !IsValidNotificationType(t) {
			return errors.New("invalid type")
		}
	} else {
		return errors.New("expected string")
	}

	return nil
}
