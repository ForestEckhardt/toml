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

	context("floats", func() {
		it("lexes key and float decimal", func() {
			items, err := mockParser(`key = -3.1_4_1_5`)
			Expect(err).NotTo(HaveOccurred())

			Expect(items).To(Equal([]item{
				{typ: itemKeyStart},
				{typ: itemText, val: "key"},
				{typ: itemFloat, val: `-3.1_4_1_5`},
				{typ: itemEOF},
			}))
		})

		it("lexes key and float exponent", func() {
			items, err := mockParser(`key = -2E-2_2`)
			Expect(err).NotTo(HaveOccurred())

			Expect(items).To(Equal([]item{
				{typ: itemKeyStart},
				{typ: itemText, val: "key"},
				{typ: itemFloat, val: `-2E-2_2`},
				{typ: itemEOF},
			}))
		})

		it("lexes key and mixed float", func() {
			items, err := mockParser(`key = -2_0.6_1_5E-2_2`)
			Expect(err).NotTo(HaveOccurred())

			Expect(items).To(Equal([]item{
				{typ: itemKeyStart},
				{typ: itemText, val: "key"},
				{typ: itemFloat, val: `-2_0.6_1_5E-2_2`},
				{typ: itemEOF},
			}))
		})

		it("lexes key and float inf", func() {
			items, err := mockParser(`key = inf`)
			Expect(err).NotTo(HaveOccurred())

			Expect(items).To(Equal([]item{
				{typ: itemKeyStart},
				{typ: itemText, val: "key"},
				{typ: itemFloat, val: `inf`},
				{typ: itemEOF},
			}))
		})

		it("lexes key and float nan", func() {
			items, err := mockParser(`key = nan`)
			Expect(err).NotTo(HaveOccurred())

			Expect(items).To(Equal([]item{
				{typ: itemKeyStart},
				{typ: itemText, val: "key"},
				{typ: itemFloat, val: `nan`},
				{typ: itemEOF},
			}))
		})

		it("lexes key and float +nan", func() {
			items, err := mockParser(`key = +nan`)
			Expect(err).NotTo(HaveOccurred())

			Expect(items).To(Equal([]item{
				{typ: itemKeyStart},
				{typ: itemText, val: "key"},
				{typ: itemFloat, val: `+nan`},
				{typ: itemEOF},
			}))
		})
	})

	context("lexes key and offset date-time", func() {
		it("lexes key and offset date-time", func() {
			items, err := mockParser(`key = 1979-05-27T00:32:00.999999+07:00`)
			Expect(err).NotTo(HaveOccurred())

			Expect(items).To(Equal([]item{
				{typ: itemKeyStart},
				{typ: itemText, val: "key"},
				{typ: itemDateTime, val: `1979-05-27T00:32:00.999999+07:00`},
				{typ: itemEOF},
			}))
		})

		it("lexes key and offset date-time where T is a space", func() {
			items, err := mockParser(`key = 1979-05-27 00:32:00.999999+07:00`)
			Expect(err).NotTo(HaveOccurred())

			Expect(items).To(Equal([]item{
				{typ: itemKeyStart},
				{typ: itemText, val: "key"},
				{typ: itemDateTime, val: `1979-05-27 00:32:00.999999+07:00`},
				{typ: itemEOF},
			}))
		})

		it("lexes key and local date-time", func() {
			items, err := mockParser(`key = 1979-05-27T00:32:00.999999`)
			Expect(err).NotTo(HaveOccurred())

			Expect(items).To(Equal([]item{
				{typ: itemKeyStart},
				{typ: itemText, val: "key"},
				{typ: itemDateTime, val: `1979-05-27T00:32:00.999999`},
				{typ: itemEOF},
			}))
		})

		it("lexes key and local date", func() {
			items, err := mockParser(`key = 1979-05-27`)
			Expect(err).NotTo(HaveOccurred())

			Expect(items).To(Equal([]item{
				{typ: itemKeyStart},
				{typ: itemText, val: "key"},
				{typ: itemDateTime, val: `1979-05-27`},
				{typ: itemEOF},
			}))
		})

		it("lexes key and local time", func() {
			items, err := mockParser(`key = 00:32:00.999999`)
			Expect(err).NotTo(HaveOccurred())

			Expect(items).To(Equal([]item{
				{typ: itemKeyStart},
				{typ: itemText, val: "key"},
				{typ: itemDateTime, val: `00:32:00.999999`},
				{typ: itemEOF},
			}))
		})
	})
}

//strconv.ParseInt("int", 0, 64)
