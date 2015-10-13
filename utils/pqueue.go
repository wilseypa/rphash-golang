package utils

type PQueue []*Node

func NewPQueue() *PQueue {
    return &PQueue{}
}

/* TODO: Need Poll. */

func (pq *PQueue) IsEmpty() bool {
    return len(*pq) == 0
}

func (pq *PQueue) Push(i interface{}) {
    a := *pq
    n := len(a)
    a = a[0 : n+1]
    r := i.(*Node)
    a[n] = r
    *pq = a
}

func (pq *PQueue) Pop() interface{} {
    a := *pq
    *pq = a[0 : len(a)-1]
    r := a[len(a)-1]
    return r
}

func (pq *PQueue) Len() int {
    return len(*pq)
}

// post: true iff is i less than j
func (pq *PQueue) Less(i, j int) bool {
    I := (*pq)[i]
    J := (*pq)[j]
    return (I.sumVal + I.myVal) < (J.sumVal + J.myVal)
}

func (pq *PQueue) Swap(i, j int) {
    (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PQueue) String() string {
    var build string = "{"
    for _, v := range *pq {
        build += v.String()
    }
    build += "}"
    return build
}
