package parser

import (
	"testing"
	"zero/ast"
	"zero/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `let x = 6;
	let y = 9;
	let temp = 69420;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if program == nil {
		t.Fatal("p.ParseProgram() return nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"temp"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, idnt string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}
	if letStmt.Name.Value != idnt {
		t.Errorf("letStm.Name.Value not '%s'. got=%s", idnt, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != idnt {
		t.Errorf("letStm.Name.TokenLiteral() not '%s'. got=%s", idnt, letStmt.Name.TokenLiteral())
		return false
	}
	return true
}
