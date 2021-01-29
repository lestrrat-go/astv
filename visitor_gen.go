package astv

import "go/ast"

type Visitor struct {
	hArrayType      ArrayTypeHandler
	hAssignStmt     AssignStmtHandler
	hBadDecl        BadDeclHandler
	hBadExpr        BadExprHandler
	hBadStmt        BadStmtHandler
	hBasicLit       BasicLitHandler
	hBinaryExpr     BinaryExprHandler
	hBlockStmt      BlockStmtHandler
	hBranchStmt     BranchStmtHandler
	hCallExpr       CallExprHandler
	hCaseClause     CaseClauseHandler
	hChanType       ChanTypeHandler
	hCommClause     CommClauseHandler
	hComment        CommentHandler
	hCommentGroup   CommentGroupHandler
	hCompositeLit   CompositeLitHandler
	hDeclStmt       DeclStmtHandler
	hDeferStmt      DeferStmtHandler
	hEllipsis       EllipsisHandler
	hEmptyStmt      EmptyStmtHandler
	hExprStmt       ExprStmtHandler
	hField          FieldHandler
	hFieldList      FieldListHandler
	hForStmt        ForStmtHandler
	hFuncDecl       FuncDeclHandler
	hFuncLit        FuncLitHandler
	hFuncType       FuncTypeHandler
	hGenDecl        GenDeclHandler
	hGoStmt         GoStmtHandler
	hIdent          IdentHandler
	hIfStmt         IfStmtHandler
	hImportSpec     ImportSpecHandler
	hIncDecStmt     IncDecStmtHandler
	hIndexExpr      IndexExprHandler
	hInterfaceType  InterfaceTypeHandler
	hKeyValueExpr   KeyValueExprHandler
	hLabeledStmt    LabeledStmtHandler
	hMapType        MapTypeHandler
	hParenExpr      ParenExprHandler
	hRangeStmt      RangeStmtHandler
	hReturnStmt     ReturnStmtHandler
	hSelectStmt     SelectStmtHandler
	hSelectorExpr   SelectorExprHandler
	hSendStmt       SendStmtHandler
	hSliceExpr      SliceExprHandler
	hStarExpr       StarExprHandler
	hStructType     StructTypeHandler
	hSwitchStmt     SwitchStmtHandler
	hTypeAssertExpr TypeAssertExprHandler
	hTypeSpec       TypeSpecHandler
	hUnaryExpr      UnaryExprHandler
	hValueSpec      ValueSpecHandler
}

func (v *Visitor) Handler(h interface{}) error {
	if x, ok := h.(ArrayTypeHandler); ok {
		v.hArrayType = x
	}
	if x, ok := h.(AssignStmtHandler); ok {
		v.hAssignStmt = x
	}
	if x, ok := h.(BadDeclHandler); ok {
		v.hBadDecl = x
	}
	if x, ok := h.(BadExprHandler); ok {
		v.hBadExpr = x
	}
	if x, ok := h.(BadStmtHandler); ok {
		v.hBadStmt = x
	}
	if x, ok := h.(BasicLitHandler); ok {
		v.hBasicLit = x
	}
	if x, ok := h.(BinaryExprHandler); ok {
		v.hBinaryExpr = x
	}
	if x, ok := h.(BlockStmtHandler); ok {
		v.hBlockStmt = x
	}
	if x, ok := h.(BranchStmtHandler); ok {
		v.hBranchStmt = x
	}
	if x, ok := h.(CallExprHandler); ok {
		v.hCallExpr = x
	}
	if x, ok := h.(CaseClauseHandler); ok {
		v.hCaseClause = x
	}
	if x, ok := h.(ChanTypeHandler); ok {
		v.hChanType = x
	}
	if x, ok := h.(CommClauseHandler); ok {
		v.hCommClause = x
	}
	if x, ok := h.(CommentHandler); ok {
		v.hComment = x
	}
	if x, ok := h.(CommentGroupHandler); ok {
		v.hCommentGroup = x
	}
	if x, ok := h.(CompositeLitHandler); ok {
		v.hCompositeLit = x
	}
	if x, ok := h.(DeclStmtHandler); ok {
		v.hDeclStmt = x
	}
	if x, ok := h.(DeferStmtHandler); ok {
		v.hDeferStmt = x
	}
	if x, ok := h.(EllipsisHandler); ok {
		v.hEllipsis = x
	}
	if x, ok := h.(EmptyStmtHandler); ok {
		v.hEmptyStmt = x
	}
	if x, ok := h.(ExprStmtHandler); ok {
		v.hExprStmt = x
	}
	if x, ok := h.(FieldHandler); ok {
		v.hField = x
	}
	if x, ok := h.(FieldListHandler); ok {
		v.hFieldList = x
	}
	if x, ok := h.(ForStmtHandler); ok {
		v.hForStmt = x
	}
	if x, ok := h.(FuncDeclHandler); ok {
		v.hFuncDecl = x
	}
	if x, ok := h.(FuncLitHandler); ok {
		v.hFuncLit = x
	}
	if x, ok := h.(FuncTypeHandler); ok {
		v.hFuncType = x
	}
	if x, ok := h.(GenDeclHandler); ok {
		v.hGenDecl = x
	}
	if x, ok := h.(GoStmtHandler); ok {
		v.hGoStmt = x
	}
	if x, ok := h.(IdentHandler); ok {
		v.hIdent = x
	}
	if x, ok := h.(IfStmtHandler); ok {
		v.hIfStmt = x
	}
	if x, ok := h.(ImportSpecHandler); ok {
		v.hImportSpec = x
	}
	if x, ok := h.(IncDecStmtHandler); ok {
		v.hIncDecStmt = x
	}
	if x, ok := h.(IndexExprHandler); ok {
		v.hIndexExpr = x
	}
	if x, ok := h.(InterfaceTypeHandler); ok {
		v.hInterfaceType = x
	}
	if x, ok := h.(KeyValueExprHandler); ok {
		v.hKeyValueExpr = x
	}
	if x, ok := h.(LabeledStmtHandler); ok {
		v.hLabeledStmt = x
	}
	if x, ok := h.(MapTypeHandler); ok {
		v.hMapType = x
	}
	if x, ok := h.(ParenExprHandler); ok {
		v.hParenExpr = x
	}
	if x, ok := h.(RangeStmtHandler); ok {
		v.hRangeStmt = x
	}
	if x, ok := h.(ReturnStmtHandler); ok {
		v.hReturnStmt = x
	}
	if x, ok := h.(SelectStmtHandler); ok {
		v.hSelectStmt = x
	}
	if x, ok := h.(SelectorExprHandler); ok {
		v.hSelectorExpr = x
	}
	if x, ok := h.(SendStmtHandler); ok {
		v.hSendStmt = x
	}
	if x, ok := h.(SliceExprHandler); ok {
		v.hSliceExpr = x
	}
	if x, ok := h.(StarExprHandler); ok {
		v.hStarExpr = x
	}
	if x, ok := h.(StructTypeHandler); ok {
		v.hStructType = x
	}
	if x, ok := h.(SwitchStmtHandler); ok {
		v.hSwitchStmt = x
	}
	if x, ok := h.(TypeAssertExprHandler); ok {
		v.hTypeAssertExpr = x
	}
	if x, ok := h.(TypeSpecHandler); ok {
		v.hTypeSpec = x
	}
	if x, ok := h.(UnaryExprHandler); ok {
		v.hUnaryExpr = x
	}
	if x, ok := h.(ValueSpecHandler); ok {
		v.hValueSpec = x
	}
	return nil
}

func (v *Visitor) Visit(n ast.Node) ast.Visitor {
	switch n := n.(type) {
	case *ast.ArrayType:
		if h := v.hArrayType; h != nil {
			if !h.ArrayType(n) {
				return nil
			}
		}
	case *ast.AssignStmt:
		if h := v.hAssignStmt; h != nil {
			if !h.AssignStmt(n) {
				return nil
			}
		}
	case *ast.BadDecl:
		if h := v.hBadDecl; h != nil {
			if !h.BadDecl(n) {
				return nil
			}
		}
	case *ast.BadExpr:
		if h := v.hBadExpr; h != nil {
			if !h.BadExpr(n) {
				return nil
			}
		}
	case *ast.BadStmt:
		if h := v.hBadStmt; h != nil {
			if !h.BadStmt(n) {
				return nil
			}
		}
	case *ast.BasicLit:
		if h := v.hBasicLit; h != nil {
			if !h.BasicLit(n) {
				return nil
			}
		}
	case *ast.BinaryExpr:
		if h := v.hBinaryExpr; h != nil {
			if !h.BinaryExpr(n) {
				return nil
			}
		}
	case *ast.BlockStmt:
		if h := v.hBlockStmt; h != nil {
			if !h.BlockStmt(n) {
				return nil
			}
		}
	case *ast.BranchStmt:
		if h := v.hBranchStmt; h != nil {
			if !h.BranchStmt(n) {
				return nil
			}
		}
	case *ast.CallExpr:
		if h := v.hCallExpr; h != nil {
			if !h.CallExpr(n) {
				return nil
			}
		}
	case *ast.CaseClause:
		if h := v.hCaseClause; h != nil {
			if !h.CaseClause(n) {
				return nil
			}
		}
	case *ast.ChanType:
		if h := v.hChanType; h != nil {
			if !h.ChanType(n) {
				return nil
			}
		}
	case *ast.CommClause:
		if h := v.hCommClause; h != nil {
			if !h.CommClause(n) {
				return nil
			}
		}
	case *ast.Comment:
		if h := v.hComment; h != nil {
			if !h.Comment(n) {
				return nil
			}
		}
	case *ast.CommentGroup:
		if h := v.hCommentGroup; h != nil {
			if !h.CommentGroup(n) {
				return nil
			}
		}
	case *ast.CompositeLit:
		if h := v.hCompositeLit; h != nil {
			if !h.CompositeLit(n) {
				return nil
			}
		}
	case *ast.DeclStmt:
		if h := v.hDeclStmt; h != nil {
			if !h.DeclStmt(n) {
				return nil
			}
		}
	case *ast.DeferStmt:
		if h := v.hDeferStmt; h != nil {
			if !h.DeferStmt(n) {
				return nil
			}
		}
	case *ast.Ellipsis:
		if h := v.hEllipsis; h != nil {
			if !h.Ellipsis(n) {
				return nil
			}
		}
	case *ast.EmptyStmt:
		if h := v.hEmptyStmt; h != nil {
			if !h.EmptyStmt(n) {
				return nil
			}
		}
	case *ast.ExprStmt:
		if h := v.hExprStmt; h != nil {
			if !h.ExprStmt(n) {
				return nil
			}
		}
	case *ast.Field:
		if h := v.hField; h != nil {
			if !h.Field(n) {
				return nil
			}
		}
	case *ast.FieldList:
		if h := v.hFieldList; h != nil {
			if !h.FieldList(n) {
				return nil
			}
		}
	case *ast.ForStmt:
		if h := v.hForStmt; h != nil {
			if !h.ForStmt(n) {
				return nil
			}
		}
	case *ast.FuncDecl:
		if h := v.hFuncDecl; h != nil {
			if !h.FuncDecl(n) {
				return nil
			}
		}
	case *ast.FuncLit:
		if h := v.hFuncLit; h != nil {
			if !h.FuncLit(n) {
				return nil
			}
		}
	case *ast.FuncType:
		if h := v.hFuncType; h != nil {
			if !h.FuncType(n) {
				return nil
			}
		}
	case *ast.GenDecl:
		if h := v.hGenDecl; h != nil {
			if !h.GenDecl(n) {
				return nil
			}
		}
	case *ast.GoStmt:
		if h := v.hGoStmt; h != nil {
			if !h.GoStmt(n) {
				return nil
			}
		}
	case *ast.Ident:
		if h := v.hIdent; h != nil {
			if !h.Ident(n) {
				return nil
			}
		}
	case *ast.IfStmt:
		if h := v.hIfStmt; h != nil {
			if !h.IfStmt(n) {
				return nil
			}
		}
	case *ast.ImportSpec:
		if h := v.hImportSpec; h != nil {
			if !h.ImportSpec(n) {
				return nil
			}
		}
	case *ast.IncDecStmt:
		if h := v.hIncDecStmt; h != nil {
			if !h.IncDecStmt(n) {
				return nil
			}
		}
	case *ast.IndexExpr:
		if h := v.hIndexExpr; h != nil {
			if !h.IndexExpr(n) {
				return nil
			}
		}
	case *ast.InterfaceType:
		if h := v.hInterfaceType; h != nil {
			if !h.InterfaceType(n) {
				return nil
			}
		}
	case *ast.KeyValueExpr:
		if h := v.hKeyValueExpr; h != nil {
			if !h.KeyValueExpr(n) {
				return nil
			}
		}
	case *ast.LabeledStmt:
		if h := v.hLabeledStmt; h != nil {
			if !h.LabeledStmt(n) {
				return nil
			}
		}
	case *ast.MapType:
		if h := v.hMapType; h != nil {
			if !h.MapType(n) {
				return nil
			}
		}
	case *ast.ParenExpr:
		if h := v.hParenExpr; h != nil {
			if !h.ParenExpr(n) {
				return nil
			}
		}
	case *ast.RangeStmt:
		if h := v.hRangeStmt; h != nil {
			if !h.RangeStmt(n) {
				return nil
			}
		}
	case *ast.ReturnStmt:
		if h := v.hReturnStmt; h != nil {
			if !h.ReturnStmt(n) {
				return nil
			}
		}
	case *ast.SelectStmt:
		if h := v.hSelectStmt; h != nil {
			if !h.SelectStmt(n) {
				return nil
			}
		}
	case *ast.SelectorExpr:
		if h := v.hSelectorExpr; h != nil {
			if !h.SelectorExpr(n) {
				return nil
			}
		}
	case *ast.SendStmt:
		if h := v.hSendStmt; h != nil {
			if !h.SendStmt(n) {
				return nil
			}
		}
	case *ast.SliceExpr:
		if h := v.hSliceExpr; h != nil {
			if !h.SliceExpr(n) {
				return nil
			}
		}
	case *ast.StarExpr:
		if h := v.hStarExpr; h != nil {
			if !h.StarExpr(n) {
				return nil
			}
		}
	case *ast.StructType:
		if h := v.hStructType; h != nil {
			if !h.StructType(n) {
				return nil
			}
		}
	case *ast.SwitchStmt:
		if h := v.hSwitchStmt; h != nil {
			if !h.SwitchStmt(n) {
				return nil
			}
		}
	case *ast.TypeAssertExpr:
		if h := v.hTypeAssertExpr; h != nil {
			if !h.TypeAssertExpr(n) {
				return nil
			}
		}
	case *ast.TypeSpec:
		if h := v.hTypeSpec; h != nil {
			if !h.TypeSpec(n) {
				return nil
			}
		}
	case *ast.UnaryExpr:
		if h := v.hUnaryExpr; h != nil {
			if !h.UnaryExpr(n) {
				return nil
			}
		}
	case *ast.ValueSpec:
		if h := v.hValueSpec; h != nil {
			if !h.ValueSpec(n) {
				return nil
			}
		}
	}
	return v
}
