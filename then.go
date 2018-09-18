package test

// Then Used for bdd testing
type Then map[string]struct {
	Got      interface{}
	Assert   func(interface{}, ...interface{}) string
	Wants    interface{}
	Format   bool
	Setup    func() interface{}
	TearDown func() interface{}
}

// Must takes a function which returns a value and an error, returning only the value.
func Must(value interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return value
}

// MustGetError takes a function which return a value and an error, returning only the error.
func MustGetError(value interface{}, err error) error {
	return err
}
