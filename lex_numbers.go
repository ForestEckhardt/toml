package toml

import (
	"unicode"
)

func lexStartNumber(l *lexer) stateFn {
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
