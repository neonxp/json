package parser

func scanNumber(l *lexer) bool {
	l.AcceptWhile("0123456789")
	if l.AtStart() {
		// not found any digit
		return false
	}
	l.Accept(".")
	l.AcceptWhile("0123456789")
	return !l.AtStart()
}

func scanQuotedString(l *lexer, quote rune) bool {
	start := l.Pos
	if l.Next() != quote {
		l.Back()
		return false
	}
	for {
		ch := l.Next()
		switch ch {
		case eof:
			l.Pos = start // Return position to start
			return false  // Unclosed quote string?
		case '\\':
			l.Next() // Skip next char
		case quote:
			return true // Closing quote
		}
	}
}
