package toml

import (
	"fmt"
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestTOML(t *testing.T) {
	suite := spec.New("toml", spec.Report(report.Terminal{}))
	suite("LexBooleans", testLexBooleans)
	suite("LexComments", testLexComments)
	suite("LexKeys", testLexKeys)
	suite("LexNumbers", testLexNumbers)
	suite("LexStrings", testLexStrings)
	suite.Run(t)
}

func mockParser(input string) ([]item, error) {
	l := lex(input)

	var items []item
	for {
		i := l.nextItem()
		switch {
		case i.typ == itemEOF:
			items = append(items, i)
			return items, nil
		case i.typ == itemError:
			return []item{}, fmt.Errorf(i.val)
		default:
			items = append(items, i)
		}
	}
}
