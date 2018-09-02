package openapi

import (
	"testing"

	"github.com/qnl3/test"
	it "github.com/smartystreets/goconvey/convey"
)

func TestTwoValueExample(t *testing.T) {

	tests := test.Tests{
		{
			Given: "a function that return two values",
			When:  "called with test.Must",
			Then: test.Then{
				"it should return only the first value": {
					Got: test.Must(
						func() (bool, error) {
							return true, nil
						}(),
					),
					Assert: it.ShouldNotEqual,
					Wants:  false,
					Format: false,
				},
			},
		},
		{
			Given: "a function that return a value and an error value",
			When:  "called with test.MustError",
			Then: test.Then{
				"it should return only the error value": {
					Got: test.MustGetError(
						func() (bool, error) {
							return true, nil
						}(),
					),
					Assert: it.ShouldHaveSameTypeAs,
					Wants:  nil,
					Format: false,
				},
			},
		},
	}

	test.RunTests(tests, t)
}
