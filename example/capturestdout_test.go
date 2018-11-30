package example

import (
	"fmt"
	"testing"

	"github.com/qnl3/test"
	it "github.com/smartystreets/goconvey/convey"
)

func TestCaptureSTDOUT(t *testing.T) {

	fn := func() {
		fmt.Print("This was printed to stdout.")
	}

	capture := test.CaptureStdout(fn)

	tests := test.Tests{
		{
			Given: "function that prints to STDOUT,",
			When:  "pasted to test.CaptureStdout,",
			Then: test.Then{
				"it should return every thing printed to STDOUT as a string": {
					Got:    capture,
					Assert: it.ShouldNotEqual,
					Wants:  "This was printed to stdout.\n",
				},
			},
		},
	}

	test.RunTests(tests, t)
}
