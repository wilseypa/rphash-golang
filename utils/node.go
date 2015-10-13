package utils

import "fmt"

type Node struct {
    row    int
    col    int
    myVal  int
    sumVal int
    parent *Node
}

func NewNode(r, c, mv, sv int, n *Node) *Node {
    return &Node{r, c, mv, sv, n}
}

func (n *Node) Eq(o *Node) bool {
    return n.row == o.row && n.col == o.col
}

func (n *Node) String() string {
    return fmt.Sprintf("{%d, %d, %d, %d}", n.row, n.col, n.myVal, n.sumVal)
}

func (n *Node) Row() int {
    return n.row
}

func (n *Node) Col() int {
    return n.col
}

func (n *Node) SetParent(p *Node) {
    n.parent = p
}

func (n *Node) Parent() *Node {
    return n.parent
}

func (n *Node) MyVal() int {
    return n.myVal
}

func (n *Node) SumVal() int {
    return n.sumVal
}

func (n *Node) SetSumVal(sv int) {
    n.sumVal = sv
}
