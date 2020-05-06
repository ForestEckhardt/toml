package toml

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func testLexNumbers(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect
	)
	context("integers", func() {
		it("lexes key and integer", func() {
			items, err := mockParser(`key = 42`)
			Expect(err).NotTo(HaveOccurred())

			Expect(items).To(Equal([]item{
				{typ: itemKeyStart},
				{typ: itemText, val: "key"},
				{typ: itemInteger, val: `42`},
				{typ: itemEOF},
			}))
		})

		it("lexes key and positive integer", func() {
			items, err := mockParser(`key = +42`)
			Expect(err).NotTo(HaveOccurred())

			Expect(items).To(Equal([]item{
				{typ: itemKeyStart},
				{typ: itemText, val: "key"},
				{typ: itemInteger, val: `+42`},
				{typ: itemEOF},
			}))
		})

		it("lexes key and negative integer", func() {
			items, err := mockParser(`key = -42`)
			Expect(err).NotTo(HaveOccurred())

			Expect(items).To(Equal([]item{
				{typ: itemKeyStart},
				{typ: itemText, val: "key"},
				{typ: itemInteger, val: `-42`},
				{typ: itemEOF},
			}))
		})

		it("lexes key and underscore seperated integer", func() {
			items, err := mockParser(`key = 1_2_3_4`)
			Expect(err).NotTo(HaveOccurred())

			Expect(items).To(Equal([]item{
				{typ: itemKeyStart},
				{typ: itemText, val: "key"},
				{typ: itemInteger, val: `1_2_3_4`},
				{typ: itemEOF},
			}))
		})

		it("lexes key and hexadecimal integer", func() {
			items, err := mockParser(`key = 0xdead_beef`)
			Expect(err).NotTo(HaveOccurred())

			Expect(items).To(Equal([]item{
				{typ: itemKeyStart},
				{typ: itemText, val: "key"},
				{typ: itemInteger, val: `0xdead_beef`},
				{typ: itemEOF},
			}))
		})

		it("lexes key and octal integer", func() {
			items, err := mockParser(`key = 0o7_5_5`)
			Expect(err).NotTo(HaveOccurred())

			Expect(items).To(Equal([]item{
				{typ: itemKeyStart},
				{typ: itemText, val: "key"},
				{typ: itemInteger, val: `0o7_5_5`},
				{typ: itemEOF},
			}))
		})

		it("lexes key and binary integer", func() {
			items, err := mockParser(`key = 0b1_0_0`)
			Expect(err).NotTo(HaveOccurred())

			Expect(items).To(Equal([]item{
				{typ: itemKeyStart},
				{typ: itemText, val: "key"},
				{typ: itemInteger, val: `0b1_0_0`},
				{typ: itemEOF},
			}))
		})
	})
}

//strconv.ParseInt("int", 0, 64)
