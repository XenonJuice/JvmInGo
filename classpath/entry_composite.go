package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	var compositeEntry []Entry
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (c CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range c {
		data, entry, err := entry.readClass(className)
		if err == nil {
			return data, entry, err
		}
	}
	return nil, nil, errors.New("class not found :" + className)
}

func (c CompositeEntry) toString() string {
	names := make([]string, len(c))
	for i, entry := range c {
		names[i] = entry.toString()
	}
	return strings.Join(names, pathListSeparator)
}
