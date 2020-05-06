package toml

import (
	"unicode"
)

func lexStartNumber(l *lexer) stateFn {
	//Ensures that there is only one + or - at beginning
	switch r := l.peek(); {
	case r == '+' || r == '-':
		l.next()
	}

	if l.peekBlock(3) == `inf` {
		l.nextBlock(3)
		l.emit(itemFloat)
		return l.popLexStack()
	}

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
	case unicode.IsDigit(r) || r == '+' || r == '-':
		if r == '+' || r == '-' {
			l.next()
		}
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
