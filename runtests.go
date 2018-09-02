package test

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// RunTests run test
func RunTests(tests Tests, t *testing.T) {
	for _, bdd := range tests {

		// execute test setup
		if bdd.Setup != nil {
			bdd.Setup()
		}

		convey.Convey(fmt.Sprintf("Given %s", bdd.Given), t, func() {
			convey.Convey(fmt.Sprintf("When %s", bdd.When), func() {
				for tk, tt := range bdd.Then {
					// format then statement if Format is true
					var thenString string
					if tt.Format {
						thenString = fmt.Sprintf(tk, tt.Wants)
					} else {
						thenString = tk
					}
					// execute then setup
					if tt.Setup != nil {
						tt.Setup()
					}

					convey.Convey(thenString, func() {
						convey.So(tt.Got, tt.Assert, tt.Wants)
					})

					//excute the teardown
					if tt.TearDown != nil {
						tt.TearDown()
					}
				}
			})
		})

		// execute test setup
		if bdd.TearDown != nil {
			bdd.TearDown()
		}
	}
}
