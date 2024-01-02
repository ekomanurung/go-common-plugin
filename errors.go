package common_plugin

type CustomError struct {
	Status int
	Err    error
}

func (e CustomError) Error() string {
	return e.Err.Error()
}
