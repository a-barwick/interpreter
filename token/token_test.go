package token

import "testing"

func TestGetIdent(t *testing.T) {
	if ident := GetIdent("fn"); ident != FUNCTION {
		t.Errorf("Got identifier %s; want %s", ident, FUNCTION)
	}

	if ident := GetIdent("let"); ident != LET {
		t.Errorf("Got identifier %s; want %s", ident, LET)
	}
}
