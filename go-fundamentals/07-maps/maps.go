package main

var (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordAlreadyExist = DictionaryErr("cannot add a word that already exist")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type (
	Dictionary    map[string]string
	DictionaryErr string
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordAlreadyExist
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
