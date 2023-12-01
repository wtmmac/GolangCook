package maps

func Search(dictionary map[string]string, keyword string) string {
	return dictionary[keyword]
}

const (
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

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
