package maps

func Search(dictionary map[string]string, keyword string) string {
	return dictionary[keyword]
}

type Dictionary map[string]string

func (d Dictionary) Search(keyword string) (string, error) {
	return d[keyword], nil
}
