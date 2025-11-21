package types

import "errors"

type JobStatus string

const (
	JobStatusQueued JobStatus = "queued"
	JobStatusRunning JobStatus = "running"
	JobStatusSuccess JobStatus = "success"
	JobStatusFailed  JobStatus = "failed"
)

func IsValidJobStatus(t JobStatus) bool {
	switch t {
	case JobStatusQueued,
		JobStatusRunning,
		JobStatusSuccess,
		JobStatusFailed:
		return true
	}

	return false
}

func ValidateJobStatus(val any) error {
	if s, ok := val.(string); ok {
		if s == "" {
			return nil
		}

		t := JobStatus(s)
		if !IsValidJobStatus(t) {
			return errors.New("invalid status")
		}
	} else if p, ok := val.(*string); ok {
		if p == nil {
			return nil
		}

		s := *p
		if s == "" {
			return nil
		}

		t := JobStatus(s)
		if !IsValidJobStatus(t) {
			return errors.New("invalid status")
		}
	} else {
		return errors.New("expected string")
	}

	return nil
}
