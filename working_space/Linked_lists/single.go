package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	New := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = New
		l.Tail = New
	} else {
		l.Tail.Next = New
		l.Tail = New
	}
}

