package main_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("TestSpec", func() {
	Describe("Test", func() {
		Context("opening", func() {
			It("should print Test", func() {
				fmt.Println("test")
			})
		})
	})
})
