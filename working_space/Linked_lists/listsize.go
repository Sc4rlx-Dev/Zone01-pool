package piscine

type NodeL1 struct {
	Data interface{}
	Next *NodeL
}

type List1 struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	cur := l.Head
	size := 0
	for cur != nil {
		size++
		cur = cur.Next
	}
	return size
}
