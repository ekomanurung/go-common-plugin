package exception

import "errors"

var NotFoundError = errors.New("NOt Found Errors")

type Exception struct {
	Code   int
	Errors error
}
