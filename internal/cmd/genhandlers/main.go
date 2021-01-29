package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/lestrrat-go/codegen"
	"github.com/pkg/errors"
)

func main() {
	if err := _main(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var typs = []struct {
	Name string
}{
	{Name: "ArrayType"},
	{Name: "AssignStmt"},
	{Name: "BadDecl"},
	{Name: "BadExpr"},
	{Name: "BadStmt"},
	{Name: "BasicLit"},
	{Name: "BinaryExpr"},
	{Name: "BlockStmt"},
	{Name: "BranchStmt"},
	{Name: "CallExpr"},
	{Name: "CaseClause"},
	{Name: "ChanType"},
	{Name: "CommClause"},
	{Name: "Comment"},
	{Name: "CommentGroup"},
	{Name: "CompositeLit"},
	{Name: "DeclStmt"},
	{Name: "DeferStmt"},
	{Name: "Ellipsis"},
	{Name: "EmptyStmt"},
	{Name: "ExprStmt"},
	{Name: "Field"},
	{Name: "FieldList"},
	{Name: "ForStmt"},
	{Name: "FuncDecl"},
	{Name: "FuncLit"},
	{Name: "FuncType"},
	{Name: "GenDecl"},
	{Name: "GoStmt"},
	{Name: "Ident"},
	{Name: "IfStmt"},
	{Name: "ImportSpec"},
	{Name: "IncDecStmt"},
	{Name: "IndexExpr"},
	{Name: "InterfaceType"},
	{Name: "KeyValueExpr"},
	{Name: "LabeledStmt"},
	{Name: "MapType"},
	{Name: "ParenExpr"},
	{Name: "RangeStmt"},
	{Name: "ReturnStmt"},
	{Name: "SelectStmt"},
	{Name: "SelectorExpr"},
	{Name: "SendStmt"},
	{Name: "SliceExpr"},
	{Name: "StarExpr"},
	{Name: "StructType"},
	{Name: "SwitchStmt"},
	{Name: "TypeAssertExpr"},
	{Name: "TypeSpec"},
	{Name: "UnaryExpr"},
	{Name: "ValueSpec"},
}

func _main() error {
	if err := genHandlers(); err != nil {
		return errors.Wrap(err, `failed to generate handlers`)
	}

	if err := genVisitor(); err != nil {
		return errors.Wrap(err, `failed to generate visitor`)
	}
	return nil
}

func genVisitor() error {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "package astv")

	fmt.Fprintf(&buf, "\n\ntype Visitor struct {")
	for _, typ := range typs {
		fmt.Fprintf(&buf, "\nh%[1]s %[1]sHandler", typ.Name)
	}
	fmt.Fprintf(&buf, "\n}")

	fmt.Fprintf(&buf, "\n\nfunc (v *Visitor) Handler(h interface{}) error {")
	for _, typ := range typs {
		fmt.Fprintf(&buf, "\nif x, ok := h.(%sHandler); ok {", typ.Name)
		fmt.Fprintf(&buf, "\nv.h%s = x", typ.Name)
		fmt.Fprintf(&buf, "\n}")
	}
	fmt.Fprintf(&buf, "\nreturn nil")
	fmt.Fprintf(&buf, "\n}")

	fmt.Fprintf(&buf, "\n\nfunc (v *Visitor) Visit(n ast.Node) ast.Visitor {")
	fmt.Fprintf(&buf, "\nswitch n := n.(type) {")
	for _, typ := range typs {
		fmt.Fprintf(&buf, "\ncase *ast.%s:", typ.Name)
		fmt.Fprintf(&buf, "\nif h := v.h%s; h != nil {", typ.Name)
		fmt.Fprintf(&buf, "\nif ! h.%s(n) {", typ.Name)
		fmt.Fprintf(&buf, "\nreturn nil")
		fmt.Fprintf(&buf, "\n}")
		fmt.Fprintf(&buf, "\n}")
	}
	fmt.Fprintf(&buf, "\n}")
	fmt.Fprintf(&buf, "\nreturn v")
	fmt.Fprintf(&buf, "\n}")

	if err := codegen.WriteFile("visitor_gen.go", &buf, codegen.WithFormatCode(true)); err != nil {
		if cfe, ok := err.(codegen.CodeFormatError); ok {
			fmt.Fprint(os.Stderr, cfe.Source())
		}

		return errors.Wrap(err, `failed to write file`)
	}
	return nil
}

func genHandlers() error {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "package astv")

	for _, typ := range typs {
		fmt.Fprintf(&buf, "\n\ntype %sHandler interface {", typ.Name)
		fmt.Fprintf(&buf, "%[1]s(*ast.%[1]s) bool", typ.Name)
		fmt.Fprintf(&buf, "\n}")
	}

	if err := codegen.WriteFile("handlers_gen.go", &buf, codegen.WithFormatCode(true)); err != nil {
		if cfe, ok := err.(codegen.CodeFormatError); ok {
			fmt.Fprint(os.Stderr, cfe.Source())
		}

		return errors.Wrap(err, `failed to write file`)
	}
	return nil
}
