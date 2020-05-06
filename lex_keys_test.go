package toml

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func testLexKeys(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect
	)

	context("when the key is quoted", func() {
		context("using double quotes", func() {
			it("lexes key and basic string", func() {
				items, err := mockParser(`"key" = "value"`)
				Expect(err).NotTo(HaveOccurred())

				Expect(items).To(Equal([]item{
					{typ: itemKeyStart},
					{typ: itemText, val: "key"},
					{typ: itemString, val: `value`},
					{typ: itemEOF},
				}))
			})
		})

		context("using single quotes", func() {
			it("lexes key and basic string", func() {
				items, err := mockParser(`'key' = "value"`)
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

	context("dotted key for properties", func() {
		it("lexes key with dot into property", func() {
			items, err := mockParser(`key.subkey = "value"`)
			Expect(err).NotTo(HaveOccurred())

			Expect(items).To(Equal([]item{
				{typ: itemKeyStart},
				{typ: itemText, val: "key"},
				{typ: itemPropertyStart},
				{typ: itemKeyStart},
				{typ: itemText, val: "subkey"},
				{typ: itemString, val: `value`},
				{typ: itemEOF},
			}))
		})
	})

	context("complex mixed dotted key for properties", func() {
		it("lexes key with dot into property", func() {
			items, err := mockParser(`key. subkey."anotherkey" = "value"`)
			Expect(err).NotTo(HaveOccurred())

			Expect(items).To(Equal([]item{
				{typ: itemKeyStart},
				{typ: itemText, val: "key"},
				{typ: itemPropertyStart},
				{typ: itemKeyStart},
				{typ: itemText, val: "subkey"},
				{typ: itemPropertyStart},
				{typ: itemKeyStart},
				{typ: itemText, val: "anotherkey"},
				{typ: itemString, val: `value`},
				{typ: itemEOF},
			}))
		})
	})
}
