package types

import (
	"errors"
	"os"
	"path"
)

type WorkDir string

func (d WorkDir) String() string {
	return string(d)
}

func (d WorkDir) DatabaseFile() string {
	return path.Join(d.String(), "data.db")
}

func (d WorkDir) CollectionsDir() string {
	return path.Join(d.String(), "collections")
}

func (d WorkDir) CollectionDirById(id string) CollectionDir {
	return CollectionDir(path.Join(d.CollectionsDir(), id))
}

type CollectionDir string

func (d CollectionDir) String() string {
	return string(d)
}

func (d CollectionDir) Images() string {
	return path.Join(d.String(), "images")
}

func (d CollectionDir) Create() error {
	dirs := []string{
		d.String(),
		d.Images(),
	}

	for _, dir := range dirs {
		err := os.Mkdir(dir, 0755)
		if err != nil && !os.IsExist(err) {
			return err
		}
	}

	return nil
}

type AdminStatus string

const (
	AdminStatusNotFixed AdminStatus = "not-fixed"
	AdminStatusFixed    AdminStatus = "fixed"
)

func IsValidAdminStatus(l AdminStatus) bool {
	switch l {
	case AdminStatusNotFixed,
		AdminStatusFixed:
		return true
	}

	return false
}

func ValidateAdminStatus(val any) error {
	if s, ok := val.(string); ok {
		if s == "" {
			return nil
		}

		t := AdminStatus(s)
		if !IsValidAdminStatus(t) {
			return errors.New("invalid admin status")
		}
	} else if p, ok := val.(*string); ok {
		if p == nil {
			return nil
		}

		s := *p
		if s == "" {
			return nil
		}

		t := AdminStatus(s)
		if !IsValidAdminStatus(t) {
			return errors.New("invalid admin status")
		}
	} else {
		return errors.New("expected string")
	}

	return nil
}
