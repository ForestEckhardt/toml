package toml

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func testLexComments(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect
	)
	context("Comments", func() {
		context("when there are comments", func() {
			it("lexes key and basic string but ignores the comments", func() {
				items, err := mockParser(`#some comment
key = "value" #some other comment`)
				Expect(err).NotTo(HaveOccurred())

				Expect(items).To(Equal([]item{
					{typ: itemKeyStart},
					{typ: itemText, val: "key"},
					{typ: itemString, val: `value`},
					{typ: itemEOF},
				}))
			})
		})
	})
}
