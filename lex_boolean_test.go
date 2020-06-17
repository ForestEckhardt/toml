package toml

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func testLexBooleans(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect
	)
	context("booleans", func() {
		it("lexes key and true", func() {
			items, err := mockParser(`key = true`)
			Expect(err).NotTo(HaveOccurred())

			Expect(items).To(Equal([]item{
				{typ: itemKeyStart},
				{typ: itemText, val: "key"},
				{typ: itemBoolean, val: `true`},
				{typ: itemEOF},
			}))
		})

		it("lexes key and false", func() {
			items, err := mockParser(`key = false`)
			Expect(err).NotTo(HaveOccurred())

			Expect(items).To(Equal([]item{
				{typ: itemKeyStart},
				{typ: itemText, val: "key"},
				{typ: itemBoolean, val: `false`},
				{typ: itemEOF},
			}))
		})
	})
}
