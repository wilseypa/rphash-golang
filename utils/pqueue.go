package utils;

import "fmt";

type Node struct {
    row    int;
    col    int;
    myVal  int;
    sumVal int;
    parent *Node;
};

func NewNode(r, c, mv, sv int, n *Node) *Node {
    return &Node{r, c, mv, sv, n};
};

func (n *Node) Eq(o *Node) bool {
    return n.row == o.row && n.col == o.col;
};

func (n *Node) String() string {
    return fmt.Sprintf("{%d, %d, %d, %d}", n.row, n.col, n.myVal, n.sumVal);
};

func (n *Node) Row() int {
    return n.row;
};

func (n *Node) Col() int {
    return n.col;
};

func (n *Node) SetParent(p *Node) {
    n.parent = p;
};

func (n *Node) Parent() *Node {
    return n.parent;
};

func (n *Node) MyVal() int {
    return n.myVal;
};

func (n *Node) SumVal() int {
    return n.sumVal;
};

func (n *Node) SetSumVal(sv int) {
    n.sumVal = sv;
};

type PQueue []*Node;

func NewPQueue(compare func, compType interface{}) *PQueue {
    return &PQueue{compare:compare, compType:compType};
};

func (pq *PQueue) IsEmpty() bool {
    return len(*pq) == 0;
}

func (pq *PQueue) Poll() interface{} {

};

func (pq *PQueue) Add(i interface{}) {
    a := *pq;
    n := len(a);
    a = a[0 : n+1];
    r := i.(*Node);
    a[n] = r;
    *pq = a;
};

func (pq *PQueue) Remove() interface{} {

};

func (pq *PQueue) Poll() interface{} {
    a := *pq;
    *pq = a[0 : len(a)-1];
    r := a[len(a)-1];
    return r;
};

func (pq *PQueue) Length() int {
    return len(*pq);
};

func (pq *PQueue) Less(i, j int) bool {
    I := (*pq)[i];
    J := (*pq)[j];
    return (I.sumVal + I.myVal) < (J.sumVal + J.myVal);
};

func (pq *PQueue) Swap(i, j int) {
    (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i];
};

func (pq *PQueue) String() string {
    var build string = "{";
    for _, v := range *pq {
        build += v.String();
    }
    build += "}";
    return build;
};
