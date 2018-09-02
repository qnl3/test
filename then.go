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
