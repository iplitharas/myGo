package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")
var ErrExists = errors.New("could not add the key because it's already exists")

func (d Dictionary) Search(key string) (string, error) {
	if value, ok := d[key]; ok {
		return value, nil
	}
	return "", ErrNotFound
}

func (d Dictionary) Add(key, value string) error {
	if _, ok := d[key]; !ok {
		d[key] = value
		return nil
	}
	return ErrExists
}
