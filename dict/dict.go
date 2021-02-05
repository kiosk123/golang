package dict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errorNotFound   = errors.New("Not found")
	errorCantUpdate = errors.New("Cant update non-existing word")
	errWordExists   = errors.New("New word already exists")
	errorCantDelete = errors.New("Cant delete non-existing word")
)

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

// update word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errorNotFound:
		return errorCantUpdate
	case nil:
		d[word] = def
	}
	return nil
}

// delete word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case errorNotFound:
		return errorCantDelete
	case nil:
		delete(d, word)
	}
	return nil
}
