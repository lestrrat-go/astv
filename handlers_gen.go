package astv

import "go/ast"

type ArrayTypeHandler interface {
	ArrayType(*ast.ArrayType) bool
}

type AssignStmtHandler interface {
	AssignStmt(*ast.AssignStmt) bool
}

type BadDeclHandler interface {
	BadDecl(*ast.BadDecl) bool
}

type BadExprHandler interface {
	BadExpr(*ast.BadExpr) bool
}

type BadStmtHandler interface {
	BadStmt(*ast.BadStmt) bool
}

type BasicLitHandler interface {
	BasicLit(*ast.BasicLit) bool
}

type BinaryExprHandler interface {
	BinaryExpr(*ast.BinaryExpr) bool
}

type BlockStmtHandler interface {
	BlockStmt(*ast.BlockStmt) bool
}

type BranchStmtHandler interface {
	BranchStmt(*ast.BranchStmt) bool
}

type CallExprHandler interface {
	CallExpr(*ast.CallExpr) bool
}

type CaseClauseHandler interface {
	CaseClause(*ast.CaseClause) bool
}

type ChanTypeHandler interface {
	ChanType(*ast.ChanType) bool
}

type CommClauseHandler interface {
	CommClause(*ast.CommClause) bool
}

type CommentHandler interface {
	Comment(*ast.Comment) bool
}

type CommentGroupHandler interface {
	CommentGroup(*ast.CommentGroup) bool
}

type CompositeLitHandler interface {
	CompositeLit(*ast.CompositeLit) bool
}

type DeclStmtHandler interface {
	DeclStmt(*ast.DeclStmt) bool
}

type DeferStmtHandler interface {
	DeferStmt(*ast.DeferStmt) bool
}

type EllipsisHandler interface {
	Ellipsis(*ast.Ellipsis) bool
}

type EmptyStmtHandler interface {
	EmptyStmt(*ast.EmptyStmt) bool
}

type ExprStmtHandler interface {
	ExprStmt(*ast.ExprStmt) bool
}

type FieldHandler interface {
	Field(*ast.Field) bool
}

type FieldListHandler interface {
	FieldList(*ast.FieldList) bool
}

type ForStmtHandler interface {
	ForStmt(*ast.ForStmt) bool
}

type FuncDeclHandler interface {
	FuncDecl(*ast.FuncDecl) bool
}

type FuncLitHandler interface {
	FuncLit(*ast.FuncLit) bool
}

type FuncTypeHandler interface {
	FuncType(*ast.FuncType) bool
}

type GenDeclHandler interface {
	GenDecl(*ast.GenDecl) bool
}

type GoStmtHandler interface {
	GoStmt(*ast.GoStmt) bool
}

type IdentHandler interface {
	Ident(*ast.Ident) bool
}

type IfStmtHandler interface {
	IfStmt(*ast.IfStmt) bool
}

type ImportSpecHandler interface {
	ImportSpec(*ast.ImportSpec) bool
}

type IncDecStmtHandler interface {
	IncDecStmt(*ast.IncDecStmt) bool
}

type IndexExprHandler interface {
	IndexExpr(*ast.IndexExpr) bool
}

type InterfaceTypeHandler interface {
	InterfaceType(*ast.InterfaceType) bool
}

type KeyValueExprHandler interface {
	KeyValueExpr(*ast.KeyValueExpr) bool
}

type LabeledStmtHandler interface {
	LabeledStmt(*ast.LabeledStmt) bool
}

type MapTypeHandler interface {
	MapType(*ast.MapType) bool
}

type ParenExprHandler interface {
	ParenExpr(*ast.ParenExpr) bool
}

type RangeStmtHandler interface {
	RangeStmt(*ast.RangeStmt) bool
}

type ReturnStmtHandler interface {
	ReturnStmt(*ast.ReturnStmt) bool
}

type SelectStmtHandler interface {
	SelectStmt(*ast.SelectStmt) bool
}

type SelectorExprHandler interface {
	SelectorExpr(*ast.SelectorExpr) bool
}

type SendStmtHandler interface {
	SendStmt(*ast.SendStmt) bool
}

type SliceExprHandler interface {
	SliceExpr(*ast.SliceExpr) bool
}

type StarExprHandler interface {
	StarExpr(*ast.StarExpr) bool
}

type StructTypeHandler interface {
	StructType(*ast.StructType) bool
}

type SwitchStmtHandler interface {
	SwitchStmt(*ast.SwitchStmt) bool
}

type TypeAssertExprHandler interface {
	TypeAssertExpr(*ast.TypeAssertExpr) bool
}

type TypeSpecHandler interface {
	TypeSpec(*ast.TypeSpec) bool
}

type UnaryExprHandler interface {
	UnaryExpr(*ast.UnaryExpr) bool
}

type ValueSpecHandler interface {
	ValueSpec(*ast.ValueSpec) bool
}

type DefaultHandler interface {
	Handle(ast.Node) bool
}
