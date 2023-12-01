package maps

import (
	"errors"
)

func Search(dictionary map[string]string, keyword string) string {
	return dictionary[keyword]
}

var (
	ErrWordExists = errors.New("cannot add word because it already exists")
	ErrNotFound   = errors.New("could not find the word you were looking for")
)

type Dictionary map[string]string

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Search(keyword string) (string, error) {
	result, ok := d[keyword]

	if !ok {
		return "", ErrNotFound
	}

	return result, nil
}
