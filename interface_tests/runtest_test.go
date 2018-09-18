package interface_tests

import (
	"testing"

	"github.com/qnl3/test"
	it "github.com/smartystreets/goconvey/convey"
)

func TestTruth(t *testing.T) {

	tests := test.Tests{
		{
			Given: "a true and false value",
			When:  "they are compared",
			Then: test.Then{
				"true should not be  equal to %v": {
					Got:    true,
					Assert: it.ShouldNotEqual,
					Wants:  false,
					Format: true,
				},
				"true should be  equal to %v": {
					Got:    true,
					Assert: it.ShouldEqual,
					Wants:  true,
					Format: true,
				},
			},
		},
		{
			Given: "a false and true value",
			When:  "they are compared",
			Then: test.Then{
				"false should not be  equal to %v": {
					Got:    false,
					Assert: it.ShouldNotEqual,
					Wants:  true,
					Format: true,
				},
				"false should be  equal to %v": {
					Got:    false,
					Assert: it.ShouldEqual,
					Wants:  false,
					Format: true,
				},
			},
		},
	}

	test.RunTests(tests, t)
}

func TestSingleValueAssertions(t *testing.T) {

	tests := test.Tests{
		{
			Given: "a nil value",
			When:  "i assert it.ShouldBeNil",
			Then: test.Then{
				"it should pass": {
					Got:    nil,
					Assert: it.ShouldBeNil,
				},
			},
		},
		{
			Given: "a non nil value",
			When:  "i assert it.ShouldNotBeNil",
			Then: test.Then{
				"it should pass": {
					Got:    "not nil",
					Assert: it.ShouldNotBeNil,
				},
			},
		},
	}

	test.RunTests(tests, t)
}
