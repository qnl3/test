package test

//Tests used for bdd testing
type Tests []struct {
	Given    string
	When     string
	Then     Then
	Setup    func() interface{}
	TearDown func() interface{}
}
