package nodes

import (
  "librarian/tree/members/nodes"
  "testing"
)

func TestAnd(t *testing.T) {
  left, right := 1, 2
  and := &nodes.And{&nodes.Binary{left, right}}
  if left != and.Left {
    t.Errorf("Expect Left And leaf to equal %v, got %v.", left, and.Left)
  } else if right != and.Right {
    t.Errorf("Expect Right And leaf to equal %v, got %v.", right, and.Right)
  }
}