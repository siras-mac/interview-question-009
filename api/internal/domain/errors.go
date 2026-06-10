package domain

import "errors"

var (
	// ErrNotFound is returned when a requested entity does not exist.
	ErrNotFound = errors.New("not found")
	// ErrEmptyMessage is returned when a comment message is blank.
	ErrEmptyMessage = errors.New("message must not be empty")
	// ErrEmptyAuthor is returned when a comment author name is blank.
	ErrEmptyAuthor = errors.New("author name must not be empty")
)
