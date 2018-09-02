package openapi

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
					Got:    true == false,
					Assert: it.ShouldNotEqual,
					Wants:  true,
					Format: true,
				},
				"true should be  equal to %v": {
					Got:    true == true,
					Assert: it.ShouldEqual,
					Wants:  true,
					Format: true,
				},
			},
		},
	}

	test.RunTests(tests, t)
}
