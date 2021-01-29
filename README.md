# astv - An Alternate Go AST Visitor

Handling Go AST usually involves a huge type switch:

```
func Visit(n ast.Node) ast.Visitor {
  switch n := n.(type) {
  case *ast.BinaryExpr:
    ...
  case *ast.GenDecl:
    ...
  }
}
```

The API in this module compartmentalizes this into Handlers.
The logic will stil be pretty much the same as a huge type switch, but you can organize them into small(er) method calls for better maintainability.

```
func (h *myHandler) BinaryExpr(n *ast.BinaryExpr) bool {
  ...
  return true // return true to keep recursing
}

func (h *myHandler) GenDecl(n *ast.GenDecl) bool {
  ...
  return false // return false to stop recursing
}

v := astv.NewVisitor()
v.Handler(h)

ast.Walk(v, node)
```

# TODO

Way to remove handlers.
Way to get/set individual handlers.
