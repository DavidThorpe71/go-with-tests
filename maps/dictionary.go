package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New(`word not found`)

func (d Dictionary) Search(word string) (string, error) {

	item, ok := d[word]

	if !ok {
		return "", ErrNotFound

	}
	return item, nil
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}
