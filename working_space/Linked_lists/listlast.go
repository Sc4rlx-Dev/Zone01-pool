package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListLast(l *List) interface{} {
	cur := l.Head
	for cur != nil {
		if cur.Next == nil {
			return cur.Data
		}
		cur = cur.Next
	}
	return nil
}
