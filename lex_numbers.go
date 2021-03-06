package toml

import (
	"unicode"
)

func lexStartNumber(l *lexer) stateFn {
	switch check := l.peekBlock(3); {
	case check == `inf` || check == `nan`:
		l.nextBlock(3)
		l.emit(itemFloat)
		return l.popLexStack()
	}
	//Ensures that there is only one + or - at beginning
	switch r := l.peek(); {
	case r == '+' || r == '-':
		l.next()

		//Allows for postive and negative inf and nan
		switch check := l.peekBlock(3); {
		case check == `inf` || check == `nan`:
			l.nextBlock(3)
			l.emit(itemFloat)
			return l.popLexStack()
		}

		return nonBase10Check(l, lexNumber)
	case unicode.IsDigit(r):
		return nonBase10Check(l, lexNumberOrDateTime)
	default:
		return lexNumber
	}
}

func lexNumberOrDateTime(l *lexer) stateFn {
	l.acceptFuncRun(integerAcceptance)

	switch r := l.peek(); {
	case r == '-' || r == ':':
		return lexDateTime
	default:
		return lexNumber
	}
}

func lexNumber(l *lexer) stateFn {
	l.acceptFuncRun(integerAcceptance)

	switch r := l.peek(); {
	case isNewline(r) || r == eof:
		l.emit(itemInteger)
		return l.popLexStack()
	case r == '_':
		l.next()
		if !unicode.IsDigit(l.peek()) {
			panic("malformed int")
		}
		return lexNumber
	case r == '.':
		l.next()
		return lexDecimal
	case r == 'e' || r == 'E':
		l.next()
		return lexStartExponent
	default:
		panic("lexNumber")
	}
}

func lexHexadecimal(l *lexer) stateFn {
	l.acceptFuncRun(hexadecimalAcceptance)

	switch r := l.peek(); {
	case isNewline(r) || r == eof:
		l.emit(itemInteger)
		return l.popLexStack()
	case r == '_':
		l.next()
		if !isHexadecimal(l.peek()) {
			panic("malformed int")
		}
		return lexHexadecimal
	default:
		panic("lexHexadecimal")
	}
}

func lexOctal(l *lexer) stateFn {
	l.acceptFuncRun(octalAcceptance)

	switch r := l.peek(); {
	case isNewline(r) || r == eof:
		l.emit(itemInteger)
		return l.popLexStack()
	case r == '_':
		l.next()
		if !isOctal(l.peek()) {
			panic("malformed int")
		}
		return lexOctal
	default:
		panic("lexOctal")
	}
}

func lexBinary(l *lexer) stateFn {
	l.acceptFuncRun(binaryAcceptance)

	switch r := l.peek(); {
	case isNewline(r) || r == eof:
		l.emit(itemInteger)
		return l.popLexStack()
	case r == '_':
		l.next()
		if !isBinary(l.peek()) {
			panic("malformed int")
		}
		return lexBinary
	default:
		panic("lexBinary")
	}
}

func lexDecimal(l *lexer) stateFn {
	l.acceptFuncRun(integerAcceptance)

	switch r := l.peek(); {
	case isNewline(r) || r == eof:
		l.emit(itemFloat)
		return l.popLexStack()
	case r == '_':
		l.next()
		if !unicode.IsDigit(l.peek()) {
			panic("malformed int")
		}
		return lexDecimal
	case r == 'e' || r == 'E':
		l.next()
		return lexStartExponent
	default:
		panic("lexDecimal")
	}
}

func lexStartExponent(l *lexer) stateFn {
	switch r := l.peek(); {
	case unicode.IsDigit(r):
		return lexExponent
	case r == '+' || r == '-':
		l.next()
		return lexExponent
	default:
		panic("lexStartExponent")
	}
}

func lexExponent(l *lexer) stateFn {
	l.acceptFuncRun(integerAcceptance)

	switch r := l.peek(); {
	case isNewline(r) || r == eof:
		l.emit(itemFloat)
		return l.popLexStack()
	case r == '_':
		l.next()
		if !unicode.IsDigit(l.peek()) {
			panic("malformed int")
		}
		return lexExponent
	default:
		panic("lexExponent")
	}
}

func lexDateTime(l *lexer) stateFn {
	l.acceptFuncRun(dateTimeAcceptance)

	switch r := l.peek(); {
	case isNewline(r) || r == eof:
		l.emit(itemDateTime)
		return l.popLexStack()
	default:
		panic("lexDateTime")
	}
}

func nonBase10Check(l *lexer, f stateFn) stateFn {
	switch l.peekBlock(2) {
	case `0x`:
		l.nextBlock(2)
		return lexHexadecimal
	case `0o`:
		l.nextBlock(2)
		return lexOctal
	case `0b`:
		l.nextBlock(2)
		return lexBinary
	default:
		return f
	}
}
