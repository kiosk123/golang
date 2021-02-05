package dict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errorNotFound = errors.New("Not found")
var errWordExists = errors.New("New word already exists")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errorNotFound
}

// add new word
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errorNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}
