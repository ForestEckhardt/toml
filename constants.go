package toml

type itemType int

const (
	itemError itemType = iota
	itemEOF

	itemInteger
	itemKeyStart
	itemMultiLineRawString
	itemMultiLineString
	itemPropertyStart
	itemRawString
	itemString
	itemText
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
	case itemText:
		return "Text"
	case itemPropertyStart:
		return "Property Start"
	default:
		return "Add this type"
	}
}
