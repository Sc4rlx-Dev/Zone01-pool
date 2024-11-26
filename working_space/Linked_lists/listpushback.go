package piscine

type NodeLs struct {
	Data interface{}
	Next *NodeL
}

type Lists struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	new := &NodeL{Data: data}
	if l.Tail == nil {
		l.Head = new
		l.Tail = new
	} else {
		l.Tail.Next = new
		l.Tail = new
	}
}
