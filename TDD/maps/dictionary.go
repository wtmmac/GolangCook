package maps

import "errors"

func Search(dictionary map[string]string, keyword string) string {
	return dictionary[keyword]
}

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Add(key, value string) {
	d[key] = value
}

func (d Dictionary) Search(keyword string) (string, error) {
	result, ok := d[keyword]

	if !ok {
		return "", ErrNotFound
	}
	return result, nil
}
