package main


import (
	"fmt"
	// "piscine"
)


type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	size := 0
 	cur := l.Head
	for cur != nil {
		size++
		cur = cur.Next
	}
	return size
}
func main() {

	link := &List{}

	link.Head = &NodeL{Data: "Hello"}
	link.Head.Next = &NodeL{Data: "man"}
	link.Head.Next.Next = &NodeL{Data: "how are you"}

	fmt.Println("List size:", ListSize(link))


}
// func main() {

// 	link := &piscine.List{}

// 	piscine.ListPushBack(link, "Hello")
// 	piscine.ListPushBack(link, "man")
// 	piscine.ListPushBack(link, "how are you")

// 	for link.Head != nil {
// 		fmt.Println(link.Head.Data)
// 		link.Head = link.Head.Next
// 	}
// }

// import "fmt"

// type node struct {
// 	data int
// 	next *node
// }
// type linkedlist struct {
// 	head *node
// }
//  func CreateLinkedlist() *linkedlist {
// 	return &linkedlist{}
// }
// func (ll *linkedlist) Append(data int) {
// 	newNode :=  &node{data: data, next: nil}
// 	if  ll.head == nil {
// 		ll.head = newNode
// 		return
// 	}
// 	cur := ll.head
// 	for cur.next != nil {
// 		cur = cur.next
// 	}
// 	cur.next = newNode
// }
// func (ll *linkedlist) Display() {
// 	cur := ll.head
// 	for cur != nil {
// 		fmt.Printf("%d -> ", cur.data)
// 		cur = cur.next
// 	}
// 	fmt.Println("nil")
// }
// func main() {
// 	// Create a new linked list
// 	ll := CreateLinkedlist()
// 	// Add elements to the linked list
// 	ll.Append(10)
// 	ll.Append(20)
// 	ll.Append(30)
// 	// Display the linked list
// 	fmt.Println("Linked List:")
// 	ll.Display()
// }