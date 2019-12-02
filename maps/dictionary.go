package main

const (
	ErrUnknownWord = DictionaryErr("unable to find the word you were looking for")
	ErrWordExists  = DictionaryErr("found duplicate key")
	ErrUpdateWord  = DictionaryErr("unable to update word that doesn't exist in dictionary")
	ErrDeleteWord  = DictionaryErr("unable to delete word that wasn't found")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrUnknownWord
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrUnknownWord:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrUnknownWord:
		return ErrUpdateWord
	case nil:
		d[key] = value
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
