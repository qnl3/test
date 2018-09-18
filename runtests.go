package test

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// RunTests run test
func RunTests(tests Tests, t *testing.T) {
	for _, tdd := range tests {

		// execute test setup
		if tdd.Setup != nil {
			tdd.Setup()
		}

		convey.Convey(fmt.Sprintf("Given %s", tdd.Given), t, func() {
			convey.Convey(fmt.Sprintf("When %s", tdd.When), func() {

				for thenTestKey, thenTest := range tdd.Then {
					// format then statement if Format is true
					var thenString string

					if thenTest.Format {
						var formatParameter interface{}
						if thenTest.Wants != nil {
							formatParameter = thenTest.Wants
						} else {
							formatParameter = thenTest.Got
						}
						thenString = fmt.Sprintf(
							"%s %s",
							"Then",
							fmt.Sprintf(thenTestKey, formatParameter),
						)
					} else {
						thenString = fmt.Sprintf("Then %s", thenTestKey)
					}
					// execute then setup
					if thenTest.Setup != nil {
						thenTest.Setup()
					}

					convey.Convey(thenString, func() {
						if thenTest.Wants != nil {
							convey.So(thenTest.Got, thenTest.Assert, thenTest.Wants)
						} else {
							convey.So(thenTest.Got, thenTest.Assert)
						}
					})

					//excute the teardown
					if thenTest.TearDown != nil {
						thenTest.TearDown()
					}

					/* change prefix to And after first test
					if thenPrefix != "And" {
						thenPrefix = "And"
					}
					*/
				}
			})
		})

		// execute test setup
		if tdd.TearDown != nil {
			tdd.TearDown()
		}
	}
}
