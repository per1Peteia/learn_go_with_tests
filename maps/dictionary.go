package main

import ()

const (
	ErrNotFound        = DictError("value not found")
	ErrKeyExists       = DictError("key already exists")
	ErrKeyDoesNotExist = DictError("key does not exist")
)

type DictError string

func (e DictError) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	val, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return val, nil
}

func (d Dictionary) Add(key, value string) error {
	if _, ok := d[key]; ok {
		return ErrKeyExists
	}
	d[key] = value
	return nil
}

func (d Dictionary) Update(key, value string) error {
	if _, ok := d[key]; !ok {
		return ErrKeyDoesNotExist
	}
	d[key] = value
	return nil
}
