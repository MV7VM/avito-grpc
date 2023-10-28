package common

import "errors"

var ErrNotFound = errors.New("such short link does not exist")
var ErrNotUrl = errors.New("isn't a url")
