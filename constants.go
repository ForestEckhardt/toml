package toml

type itemType int

const (
	itemError itemType = iota
	itemEOF

	itemFloat
	itemInteger
	itemKeyStart
	itemMultiLineRawString
	itemMultiLineString
	itemPropertyStart
	itemRawString
	itemString
	itemText
	itemBoolean
	itemDateTime
)

const (
	eof = -1

	commentStart       = '#'
	escapeCharacter    = '\\'
	keySeperator       = '='
	rawStringDelimiter = '\''
	stringDelimiter    = '"'
	tableSeperator     = '.'
)

func (i item) String() string {
	switch i.typ {
	case itemError:
		return "Error"
	case itemEOF:
		return "EOF"
	case itemKeyStart:
		return "Key Start"
	case itemString:
		return "String"
	case itemMultiLineString:
		return "Multi-line String"
	case itemRawString:
		return "Raw String"
	case itemMultiLineRawString:
		return "Multi-line Raw String"
	case itemInteger:
		return "Integer"
	case itemFloat:
		return "Float"
	case itemText:
		return "Text"
	case itemBoolean:
		return "Boolean"
	case itemDateTime:
		return "Date-Time"
	case itemPropertyStart:
		return "Property Start"
	default:
		return "Add this type"
	}
}
