package util_test

import (
	"github.com/liam-lai/service-show-filter/app/util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Schema", func() {
	Context("Should Remove Bracket String", func() {
		It("bracket in between the string", func() {
			input := "The (Le Goût) Taste"
			expected := "The Taste"
			output := util.RemoveBracket(input)
			Expect(output).To(Equal(expected))
		})
		It("bracket at the front the string", func() {
			input := "(Le Goût) The Taste"
			expected := "The Taste"
			output := util.RemoveBracket(input)
			Expect(output).To(Equal(expected))
		})
		It("bracket at the end the string", func() {
			input := "The Taste (Le Goût)"
			expected := "The Taste"
			output := util.RemoveBracket(input)
			Expect(output).To(Equal(expected))
		})
	})

})
