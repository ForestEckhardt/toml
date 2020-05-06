package toml

func lexBasicString(l *lexer) stateFn {
	l.acceptFuncRun(basicStringAcceptance)

	switch r := l.peek(); {
	case isNewline(r):
		panic("unexpected newline in basic string")
	case r == escapeCharacter:
		l.next()
		l.pushLexStack(lexBasicString)
		return lexEscapeCharacter
	case r == stringDelimiter:
		l.emit(itemString)
		l.skip()
		return l.popLexStack()
	default:
		panic("lexBasicString")
	}
}

func lexMultiLineBasicString(l *lexer) stateFn {
	l.acceptFuncRun(basicStringAcceptance)

	switch r := l.peek(); {
	case isNewline(r):
		l.next()
		return lexMultiLineBasicString
	case r == escapeCharacter:
		l.next()
		l.pushLexStack(lexMultiLineBasicString)
		return lexMultiLineEscapeCharacter
	case r == stringDelimiter:
		if l.peekBlock(3) == `"""` {
			l.emit(itemMultiLineString)
			l.skipBlock(3)
			return l.popLexStack()
		}
		l.next()
		return lexMultiLineBasicString
	default:
		panic("lexMultiLineBasicString")
	}
}

func lexMultiLineEscapeCharacter(l *lexer) stateFn {
	if isNewline(l.peek()) {
		l.next()
		return l.popLexStack()
	}
	return lexEscapeCharacter
}

func lexEscapeCharacter(l *lexer) stateFn {
	switch r := l.next(); {
	case allowedEscapeCharacter(r):
		return l.popLexStack()
	case r == 'u':
		return lexShortUnicode
	case r == 'U':
		return lexLongUnicode
	default:
		panic("lexEscapeCharacter")
	}
}

func lexShortUnicode(l *lexer) stateFn {
	for i := 0; i < 4; i++ {
		if !isHexadecimal(l.next()) {
			panic("lexShortUnicode")
		}
	}
	return l.popLexStack()
}

func lexLongUnicode(l *lexer) stateFn {
	for i := 0; i < 8; i++ {
		if !isHexadecimal(l.next()) {
			panic("lexLongUnicode")
		}
	}
	return l.popLexStack()
}

func lexRawString(l *lexer) stateFn {
	l.acceptFuncRun(rawStringAcceptance)

	switch r := l.peek(); {
	case isNewline(r):
		panic("unexpected newline in basic string")
	case r == rawStringDelimiter:
		l.emit(itemRawString)
		l.skip()
		return l.popLexStack()
	default:
		panic("lexRawString")
	}
}

func lexMultiLineRawString(l *lexer) stateFn {
	l.acceptFuncRun(rawStringAcceptance)

	switch r := l.peek(); {
	case isNewline(r):
		l.next()
		return lexMultiLineRawString
	case r == rawStringDelimiter:
		if l.peekBlock(3) == `'''` {
			l.emit(itemMultiLineRawString)
			l.skipBlock(3)
			return l.popLexStack()
		}
		l.next()
		return lexMultiLineRawString
	default:
		panic("lexMultiLineRawString")
	}
}
