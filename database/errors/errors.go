package errs

import "errors"

var (
	ErrDocumentNotFound = errors.New("document not found")
	ErrFailedToFetch    = errors.New("failed to fetch document")
	ErrFailedToInsert   = errors.New("failed to insert document")
)
