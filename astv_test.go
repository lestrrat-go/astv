package astv_test

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/lestrrat-go/astv"
	"github.com/stretchr/testify/assert"
)

type Handler struct {
	accumulator []ast.Node
}

func (h *Handler) BasicLit(n *ast.BasicLit) bool {
	h.accumulator = append(h.accumulator, n)
	return true
}
func (h *Handler) Ident(n *ast.Ident) bool {
	h.accumulator = append(h.accumulator, n)
	return true
}

func TestASTVisitor(t *testing.T) {
	src := `
package p
const c = 1.0
var X = f(3.14)*2 + c
`
	h := &Handler{}
	v := astv.NewVisitor()
	if !assert.NoError(t, v.Handler(h), `v.Handler should succeed`) {
		return
	}

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Inspect the AST and print all identifiers and literals.
	ast.Walk(v, f)

	if !assert.Len(t, h.accumulator, 8) {
		return
	}
}
