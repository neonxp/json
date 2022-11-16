package parser

func initJson(l *lexer) stateFunc {
	ignoreWhiteSpace(l)
	switch {
	case l.Accept("{"):
		l.Emit(lObjectStart)
		return stateInObject
	case l.Accept("["):
		l.Emit(lArrayStart)
	case l.Peek() == eof:
		return nil
	}
	return l.Errorf("Unknown token: %s", string(l.Peek()))
}

func stateInObject(l *lexer) stateFunc {
	// we in object, so we expect field keys and values
	ignoreWhiteSpace(l)
	if l.Accept("}") {
		l.Emit(lObjectEnd)
		// If meet close object return to previous state (including initial)
		return l.PopState()
	}
	ignoreWhiteSpace(l)
	l.Accept(",")
	ignoreWhiteSpace(l)
	if !scanQuotedString(l, '"') {
		return l.Errorf("Unknown token: %s", string(l.Peek()))
	}
	l.Emit(lObjectKey)
	ignoreWhiteSpace(l)
	if !l.Accept(":") {
		return l.Errorf("Expected ':'")
	}
	ignoreWhiteSpace(l)
	l.Emit(lObjectValue)
	switch {
	case scanQuotedString(l, '"'):
		l.Emit(lString)
		ignoreWhiteSpace(l)
		l.Accept(",")
		l.Ignore()
		ignoreWhiteSpace(l)
		return stateInObject
	case scanNumber(l):
		l.Emit(lNumber)
		ignoreWhiteSpace(l)
		l.Accept(",")
		l.Ignore()
		ignoreWhiteSpace(l)
		return stateInObject
	case l.AcceptAnyOf([]string{"true", "false"}, true):
		l.Emit(lBoolean)
		ignoreWhiteSpace(l)
		l.Accept(",")
		l.Ignore()
		ignoreWhiteSpace(l)
		return stateInObject
	case l.AcceptString("null", true):
		l.Emit(lNull)
		ignoreWhiteSpace(l)
		l.Accept(",")
		l.Ignore()
		ignoreWhiteSpace(l)
		return stateInObject
	case l.Accept("{"):
		l.Emit(lObjectStart)
		l.PushState(stateInObject)
		return stateInObject
	case l.Accept("["):
		l.Emit(lArrayStart)
		l.PushState(stateInObject)
		return stateInArray
	}
	return l.Errorf("Unknown token: %s", string(l.Peek()))
}

func stateInArray(l *lexer) stateFunc {
	ignoreWhiteSpace(l)
	l.Accept(",")
	ignoreWhiteSpace(l)
	switch {
	case scanQuotedString(l, '"'):
		l.Emit(lString)
	case scanNumber(l):
		l.Emit(lNumber)
	case l.AcceptAnyOf([]string{"true", "false"}, true):
		l.Emit(lBoolean)
	case l.AcceptString("null", true):
		l.Emit(lNull)
	case l.Accept("{"):
		l.Emit(lObjectStart)
		l.PushState(stateInArray)
		return stateInObject
	case l.Accept("["):
		l.Emit(lArrayStart)
		l.PushState(stateInArray)
		return stateInArray
	case l.Accept("]"):
		l.Emit(lArrayEnd)
		return l.PopState()
	}
	return stateInArray
}

func ignoreWhiteSpace(l *lexer) {
	l.AcceptWhile(" \n\t") // ignore whitespaces
	l.Ignore()
}
