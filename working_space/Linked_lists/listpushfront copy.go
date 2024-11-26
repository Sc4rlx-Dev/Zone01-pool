package piscine

type NodeLss struct {
	Data interface{}
	Next *NodeL
}

type Listss struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	new := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = new
		l.Tail = new
	} else {
		new.Next = l.Head
		l.Head = new
	}
}
