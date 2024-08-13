package main

import (
	"fmt"
	"unsafe"
)

const (
	ErrValueNotFound = DictionaryErr("no value found for the given key")
	ErrAlreadyExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]
	if !ok {
		return "", ErrValueNotFound
	}
	return def, nil
}

func (d Dictionary) Add(k, v string) error {
	_, err := d.Search(v)
	switch err {
	case ErrValueNotFound:
		d[k] = v
	case nil:
		return ErrAlreadyExists
	default:
		return err
	}
	return nil
}

func main() {
	var m map[int]int
	var p uintptr
	fmt.Println(unsafe.Sizeof(m), unsafe.Sizeof(p)) // 8 8 (linux/amd64/arm)
}
