package parser

import (
	_ "github.com/nwehr/monkey/ast"
	"github.com/nwehr/monkey/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
	let a =  5;
	let b = 10;
	`

	p := New(lexer.New(input))

	program, err := p.ParseProgram()
	if err != nil {
		t.Fatalf("Could not parse program: %v", err)
	}

	if len(program.Statements) != 2 {
		t.Fatalf("Expected 2 statements, got %d", len(program.Statements))
	}

}
