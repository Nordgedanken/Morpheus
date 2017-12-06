package main_test

import (
	"fmt"

	"github.com/onsi/ginkgo"
)

var Describe = ginkgo.Describe
var Context = ginkgo.Context
var It = ginkgo.It
var _ = Describe("TestSpec", func() {
	Describe("Test", func() {
		Context("opening", func() {
			It("should print Test", func() {
				fmt.Println("test")
			})
		})
	})
})
