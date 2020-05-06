package toml

func lexKeyStart(l *lexer) stateFn {
	switch r := l.peek(); {
	case isWhitespace(r):
		l.skip()
		return lexKeyStart
	case r == stringDelimiter:
		l.skip()
		l.emit(itemKeyStart)
		return lexStringKey
	case r == rawStringDelimiter:
		l.skip()
		l.emit(itemKeyStart)
		return lexRawStringKey
	default:
		l.emit(itemKeyStart)
		return lexBareKey
	}
}

func lexStringKey(l *lexer) stateFn {
	l.acceptFuncRun(basicStringAcceptance)

	switch r := l.peek(); {
	case r == escapeCharacter:
		l.next()
		return lexEscapeCharacter
	case r == stringDelimiter:
		l.emit(itemText)
		l.skip()
		return lexKeyEnd
	default:
		panic("lexStringKey")
	}
}

func lexRawStringKey(l *lexer) stateFn {
	l.acceptFuncRun(rawStringAcceptance)

	switch r := l.peek(); {
	case isNewline(r):
		panic("unexpected new line in key")
		// return l.errorf("unexpected new line in basic string")
	case r == rawStringDelimiter:
		l.emit(itemText)
		l.skip()
		return lexKeyEnd
	default:
		panic("lexRawStringKey")
	}
}

func lexBareKey(l *lexer) stateFn {
	l.acceptFuncRun(keyAcceptance)

	switch r := l.peek(); {
	case isWhitespace(r):
		l.emit(itemText)
		return lexKeyEnd
	case r == keySeperator:
		l.emit(itemText)
		return lexKeyEnd
	case r == tableSeperator:
		l.emit(itemText)
		return lexKeyEnd
	default:
		panic("lexBareKey")
		// return l.errorf("invalid character %q in key", r)
	}

}

func lexKeyEnd(l *lexer) stateFn {
	switch r := l.peek(); {
	case r == keySeperator:
		l.skip()
		return lexValue
	case isWhitespace(r):
		l.skip()
		return lexKeyEnd
	case r == tableSeperator:
		l.skip()
		return lexPropertyStart
	default:
		panic("lexKeyEnd")
		// return l.errorf("expected a key seperator %q, but found %q", keySeperator, r)
	}
}

func lexPropertyStart(l *lexer) stateFn {
	l.emit(itemPropertyStart)
	return lexKeyStart
}
