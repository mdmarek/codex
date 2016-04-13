package sql

import (
	"fmt"
)

import (
	"github.com/chuckpreslar/codex/lib/aux/consts"
	"github.com/chuckpreslar/codex/lib/interfaces"
	"github.com/chuckpreslar/codex/lib/nodes"
)

// And ...
type And struct{}

// VisitAndNode ...
func (a And) VisitAndNode(node *nodes.And, visitor interfaces.Visitor) string {
	return fmt.Sprintf("(%s %s %s)", node.Left.Accept(visitor), consts.And, node.Right.Accept(visitor))
}