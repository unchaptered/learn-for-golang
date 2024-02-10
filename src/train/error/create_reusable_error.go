package main

import "errors"

var ErrNotFound = errors.New("Employee not found!")

func create_reusable_error(id int) (*string, error) {
	if id != 1001 {
		return nil, ErrNotFound
	}

	result := "hello"
	return &result, nil
}
