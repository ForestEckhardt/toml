package toml

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func testLexStrings(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect
	)

	context("Strings", func() {
		context("basic string", func() {
			it("lexes key and basic string", func() {
				items, err := mockParser(`key = "value"`)
				Expect(err).NotTo(HaveOccurred())

				Expect(items).To(Equal([]item{
					{typ: itemKeyStart},
					{typ: itemText, val: "key"},
					{typ: itemString, val: `value`},
					{typ: itemEOF},
				}))
			})

			it("lexes key and basic string with escapes", func() {
				items, err := mockParser(`key = "value\n\t\u1111\U22222222"`)
				Expect(err).NotTo(HaveOccurred())

				Expect(items).To(Equal([]item{
					{typ: itemKeyStart},
					{typ: itemText, val: "key"},
					{typ: itemString, val: `value\n\t\u1111\U22222222`},
					{typ: itemEOF},
				}))
			})
		})

		context("multi-line basic strings", func() {
			it("lexes key and basic string", func() {
				items, err := mockParser(`key = """Here are fifteen quotation marks: ""\"""\"""\"""\"""\"."""`)
				Expect(err).NotTo(HaveOccurred())

				Expect(items).To(Equal([]item{
					{typ: itemKeyStart},
					{typ: itemText, val: "key"},
					{typ: itemMultiLineString, val: `Here are fifteen quotation marks: ""\"""\"""\"""\"""\".`},
					{typ: itemEOF},
				}))
			})

			it("lexes key and basic string", func() {
				items, err := mockParser(`key = """
The quick brown \


  fox jumps over \
    the lazy dog."""`)
				Expect(err).NotTo(HaveOccurred())

				Expect(items).To(Equal([]item{
					{typ: itemKeyStart},
					{typ: itemText, val: "key"},
					{typ: itemMultiLineString, val: `
The quick brown \


  fox jumps over \
    the lazy dog.`},
					{typ: itemEOF},
				}))
			})
		})

		context("raw strings", func() {
			it("lexes key and raw string", func() {
				items, err := mockParser(`key = 'Tom \"Dubs\" Preston-Werner'`)
				Expect(err).NotTo(HaveOccurred())

				Expect(items).To(Equal([]item{
					{typ: itemKeyStart},
					{typ: itemText, val: "key"},
					{typ: itemRawString, val: `Tom \"Dubs\" Preston-Werner`},
					{typ: itemEOF},
				}))
			})
		})

		context("multi-line raw strings", func() {
			it("lexes key and raw string", func() {
				items, err := mockParser(`key = '''
The 'first' newline is
trimmed in raw strings.
   All other whitespace
   is preserved.
'''`)
				Expect(err).NotTo(HaveOccurred())

				Expect(items).To(Equal([]item{
					{typ: itemKeyStart},
					{typ: itemText, val: "key"},
					{typ: itemMultiLineRawString, val: `
The 'first' newline is
trimmed in raw strings.
   All other whitespace
   is preserved.
`},
					{typ: itemEOF},
				}))
			})
		})
	})
}
